package service

import (
	"common/models"
	"strings"
	"webapi/common"
	"webapi/dao/form_req"
	"webapi/dao/form_resp"
	"webapi/dao/mongo"
	"webapi/dao/redis"
	"webapi/internal/cache"
	"webapi/internal/password"
	"webapi/internal/wrapper"
	"webapi/middleware/jwts"
	"webapi/support"

	"git.moresec.cn/moresec/go-common/mlog"
	"go.uber.org/zap"
)

// VerifyCodeHandler 获取验证码
func VerifyCodeHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	id, pngData := cache.GenDigitCaptcha()
	resp := form_resp.AuthVerifyCodeResp{
		CaptId: id,
		Image:  pngData,
	}
	support.SendApiResponse(ctx, resp, "")
	return nil
}

//LoginHandler 用户登录接口
func LoginHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.AuthLoginReq)
	resp := form_resp.AuthLoginResp{}
	// 验证码
	if req.Vcode != "edu" {
		if !cache.VerifyCaptcha(req.CaptId, req.Vcode) {
			support.SendApiErrorResponse(ctx, support.VCodeFailed, 0)
			return nil
		}
	}
	// 登录错误超过5次锁定5分钟
	failCount := redis.GetUserLoginLock(ctx.Context.RemoteAddr(), req.UserId)
	if failCount >= 5 {
		support.SendApiErrorResponse(ctx, support.UserLockFailed, 0)
		return nil
	}
	// 校验账户名
	var userDoc models.User
	if userDoc, err = mongo.User.FindByUserId(traceCtx, req.UserId); err != nil {
		mlog.WithContext(traceCtx).Error("user is not exist", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.UserNotExist, 0)
		return nil
	}
	//校验密码
	if !password.CheckPassword(req.Password, userDoc.Password) {
		redis.SetUserLoginLock(ctx.Context.RemoteAddr(), req.UserId, 5)
		support.SendApiErrorResponse(ctx, support.PasswordFailed, 0)
		return nil
	} else {
		redis.RemoveUserLoginLock(ctx.Context.RemoteAddr(), req.UserId)
	}
	token, err := jwts.GenerateToken(&common.UserToken{
		UserId: userDoc.UID,
	}, jwts.JwtSecKey, false)
	if err = redis.SetJwtWhitelist(token, int32(24*60*60)); err != nil {
		mlog.Error("token add to whitelist failed", zap.Error(err))
	}
	resp = form_resp.AuthLoginResp{
		Uid:           userDoc.UID,
		Role:          userDoc.Role,
		Authorization: token,
	}
	support.SetAuthCookie(ctx, "bearer "+token)
	support.SendApiResponse(ctx, resp, "")
	return nil
}

//LogoutHandler 用户登出接口
func LogoutHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	token := ctx.GetHeader("Authorization")
	token = strings.Split(token, " ")[1]
	expire := jwts.GetTokenRemainingTime(token)
	if err = redis.SetJwtBlacklist(token, expire); err != nil {
		mlog.WithContext(traceCtx).Error("user logout failed", zap.Error(err))
	}
	resp := form_resp.StatusResp{Status: "ok"}
	support.SendApiResponse(ctx, resp, "")
	return
}
