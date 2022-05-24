package jwts

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"webapi/common"
	"webapi/config"
	"webapi/dao/mongo"
	"webapi/dao/redis"
	"webapi/support"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"

	"git.moresec.cn/moresec/go-common/mlog"
	"go.uber.org/zap"

	"github.com/kataras/iris/v12/context"
)

type (
	errorHandler   func(context.Context, string)
	TokenExtractor func(context.Context) (string, error) //token提取方法，内置的
	Jwts           struct {
		Config Config
	}
)

var (
	jwts      *Jwts
	lock      sync.Mutex
	JwtSecKey = common.JWTFixedSecKey
	SysToken  string
)

// Serve jwt 服务
func Serve(ctx context.Context) bool {
	ConfigJWT()
	if err := jwts.CheckJWT(ctx); err != nil {
		mlog.Debug("Check jwt error", zap.Error(err), zap.String("url", ctx.Request().RequestURI))
		return false
	}
	return true
}

// FromAuthHeader 获取表单JwtToken头部信息(bearer)
func FromAuthHeader(ctx context.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", nil
	}

	authHeaderParts := strings.Split(authHeader, " ") //bearer格式解析
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", fmt.Errorf("Authorization header format must be Bearer {token} ")
	}
	return authHeaderParts[1], nil
}

// ConfigJWT 配置Jwt中间件
func ConfigJWT() {
	if jwts != nil {
		return
	}

	lock.Lock()
	defer lock.Unlock()

	cfg := Config{
		ContextKey: DefaultContextKey,
		ValidationKeyGetter: func(token *jwt.Token) (i interface{}, e error) {
			return []byte(JwtSecKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(ctx context.Context, errMsg string) {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.StopExecution()
		},
		Extractor:           FromAuthHeader,
		Expiration:          true,
		Debug:               true,
		EnableAuthOnOptions: false,
	}
	jwts = &Jwts{Config: cfg}
}

// CheckJWT 检验Jwt的Token是否正确
func (j *Jwts) CheckJWT(ctx context.Context) error {
	if !j.Config.EnableAuthOnOptions {
		if ctx.Method() == iris.MethodOptions {
			return nil
		}
	}
	// 使用自定义token提取函数(FromAuthHeader)
	token, err := j.Config.Extractor(ctx)
	if err != nil {
		j.logf("Error extracting token: %v", err)
		return fmt.Errorf("Error extracting token: %v ", err)
	}
	// Jwt token 为空时
	if token == "" {
		j.logf("  Error: No credentials found (CredentialsOptional=false)")
		return fmt.Errorf(support.TokenParseFailedAndEmpty)
	}
	// 校验Jwt的token是否在黑名单中
	if redis.InJwtBlacklist(token) {
		return fmt.Errorf(support.TokenParseFailedAndInvalid)
	}
	if !redis.InJwtWhitelist(token) {
		return fmt.Errorf(support.TokenParseFailedAndInvalid)
	} else {
		if err := redis.FlushJwtWhitelist(token); err != nil {
			return fmt.Errorf(support.TokenFlushFailed)
		}
	}
	// 进行Jwt Token 的解析
	parseToken, err := jwt.Parse(token, j.Config.ValidationKeyGetter)
	if err != nil {
		j.logf("Error parsing token: %v")
		return fmt.Errorf("Error parsing token: %v ", err)
	}
	// 数据算法校验(alg字段)
	if j.Config.SigningMethod != nil && j.Config.SigningMethod.Alg() != parseToken.Header["alg"] {
		message := fmt.Sprintf("Expected %s signing method but token specified %s",
			j.Config.SigningMethod.Alg(), parseToken.Header["alg"])
		j.logf("Error validating token algorithm: %s", message)
	}
	// 数据字段校验
	if !parseToken.Valid { //参数值为空时
		j.logf(support.TokenParseFailedAndInvalid)
		return fmt.Errorf(support.TokenParseFailedAndInvalid)
	}
	if claims, ok := parseToken.Claims.(jwt.MapClaims); ok {
		user := &common.UserToken{
			UserId: claims["userId"].(string),
			Role:   claims["role"].(int),
		}
		ctx.Values().Set(j.Config.ContextKey, user)
		if config.IrisConf.JWT.JwtRenewSwitch {
			renewJwt(ctx, claims, user)
		}
	}
	return nil
}

// jwt 续期函数
func renewJwt(ctx context.Context, claims jwt.MapClaims, user *common.UserToken) {
	remain, err := remainingTime(claims["exp"].(float64)) //jwt token 续租
	if err != nil {
		mlog.Debug("renewJwt: generate token error", zap.Error(err), zap.Any("user", user))
		return
	}
	if remain < config.IrisConf.JWT.JwtTimeOut/10 {
		token, err := GenerateToken(user, JwtSecKey, false) //生成JWT Token
		if err != nil {
			mlog.Debug("renewJwt: generate token error", zap.Error(err), zap.Any("user", user))
			return
		}
		systemCfg := mongo.SystemConfig.Get(ctx.Request().Context())
		if err = redis.SetJwtWhitelist(token, int32(systemCfg.NoOpExitTm)); err != nil {
			mlog.Debug("renewJwt: generate token error", zap.Error(err), zap.Any("user", user))
			return
		}
		support.SetAuthCookie(ctx, "bearer "+token)
	}
}

// jwt 续期时间计算函数
func remainingTime(exp float64) (int64, error) {
	if int64(exp) > time.Now().Unix() {
		return int64(exp) - time.Now().Unix(), nil
	}
	return 0, fmt.Errorf("get claims expire time error, expireTime: %d", exp)
}

// jwt异常日志
func (j *Jwts) logf(format string, args ...interface{}) {
	if j.Config.Debug {
		mlog.Debugf(format, args...)
	}
}

type Claims struct {
	UserId string `json:"userId"`
	Role   int `json:"role"`
	jwt.StandardClaims
}
