package controller

import (
	"webapi/dao/form_req"
	"webapi/internal/wrapper"
	"webapi/service"
	"webapi/support"
)

type AuthController struct {
}

// VerifyCode
// @Summary 基础接口 - 获取验证码
// @Description get verifycode
// @Tags common
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} form_resp.AuthVerifyCodeResp "response data"
// @Router /auth/verifycode/ [get]
// @Security ApiKeyAuth
func (a AuthController) VerifyCode(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.VerifyCodeHandler, false, nil, nil)
}

// Login
// @Summary 登录
// @Description user login
// @Tags common
// @Accept json
// @Produce json
// @Param auth body form_req.AuthLoginReq true "request data"
// @Success 200 {object} form_resp.AuthLoginResp "response data"
// @Router /auth/login/ [post]
// @Security ApiKeyAuth
func (a AuthController) Login(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.LoginHandler, true, &form_req.AuthLoginReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// Logout
// @Summary 基础接口 - 用户登出
// @Description user logout
// @Tags common
// @Accept json
// @Produce json
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /auth/logout/ [post]
// @Security ApiKeyAuth
// @Param authorization header string true "authorization"
func (a AuthController) Logout(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.LogoutHandler, false, nil, nil)
}

// ChangePassword
// @Summary 基础接口 - 修改账户密码
// @Description change user password
// @Tags common
// @Accept json
// @Produce json
// @Param auth body form_req.ChangePasswordReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /auth/change_password/ [post]
// @Security ApiKeyAuth
// @Param authorization header string true "authorization"
func (a AuthController) ChangePassword(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.ChangePasswordHandler, true, form_req.ChangePasswordReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}
