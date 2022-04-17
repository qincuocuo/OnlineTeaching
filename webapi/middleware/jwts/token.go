package jwts

import (
	"common/models"
	"fmt"
	"time"
	"webapi/common"
	"webapi/config"
	"webapi/internal/db"
	"webapi/utils"

	mbase "git.moresec.cn/moresec/go-common/mbase/basic"
	"gopkg.in/mgo.v2/bson"

	"github.com/dgrijalva/jwt-go"

	"github.com/kataras/iris/v12/context"
)

// GenerateToken 生成 JWT token
func GenerateToken(user *common.UserToken, secretKey string, keep bool) (string, error) {
	var expireTime time.Time
	if keep {
		expireTime = time.Now().AddDate(1000, 0, 0)
	} else {
		fmt.Println(config.IrisConf.JWT.JwtTimeOut)
		expireTime = time.Now().Add(time.Duration(config.IrisConf.JWT.JwtTimeOut) * time.Second)
	}
	claims := Claims{
		user.UserId,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "iris-casbins-jwt",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(secretKey))
	return token, err
}

// GetUserToken 获取用户Token
func GetUserToken(ctx context.Context) *common.UserToken {
	if ctx.Values().Get(DefaultContextKey) != nil {
		return ctx.Values().Get(DefaultContextKey).(*common.UserToken)
	}
	return nil
}

func GetTokenRemainingTime(token string) int32 {
	parsedToken, err := jwt.Parse(token, jwts.Config.ValidationKeyGetter)
	if err != nil {
		jwts.logf("Error parsing token1: %v", err)
		return 0
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		remain, err := remainingTime(claims["exp"].(float64))
		if err != nil {
			jwts.logf("Error get remain time, claims: %v", claims)
			return 0
		}
		return int32(remain)
	}
	return 0
}

func InitSysToken() (err error) {
	err = utils.Retry(5, 30, func() error {
		var sysDoc models.SystemInfo
		if _, err := db.MongoCli.FindOne(sysDoc.CollectName(), bson.M{}, &sysDoc); err != nil {
			return err
		}
		SysToken = mbase.String.MD5String(sysDoc.Id.Hex(), 32)
		JwtSecKey += SysToken
		return nil
	})
	return
}
