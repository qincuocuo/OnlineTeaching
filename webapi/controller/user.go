package controller

import (
	"webapi/dao/form_req"
	"webapi/internal/wrapper"
	"webapi/service"
	"webapi/support"
)

type UserController struct {
}

// CreateUser
// @Summary 创建用户
// @Description create user
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param auth body form_req.CreateUserReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /auth/register/ [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (u UserController) CreateUser(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.CreateUserHandler, true, &form_req.CreateUserReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// UserInfo
// @Summary 获取用户信息
// @Description get user info
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 {object} form_resp.UserInfoResp "response data"
// @Router /v1/user/ [get]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (u UserController) UserInfo(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.UserInfoHandler, false, nil, nil)
}

// UserPassword
// @Summary 忘记密码
// @Description get user password
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param auth body form_req.UserPasswordReq true "request data"
// @Success 200 {object} form_resp.UserPasswordResp "response data"
// @Router /v1/user/password/ [post]
// @Security ApiKeyAuth
// @Param Authorization header string true "authentication"
func (u UserController) UserPassword(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.UserPasswordHandler, true, &form_req.UserPasswordReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// ChangePassword
// @Summary 修改账户密码
// @Description change user password
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param auth body form_req.ChangePasswordReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /v1/user/change_password/ [post]
// @Security ApiKeyAuth
// @Param authorization header string true "authorization"
func (a UserController) ChangePassword(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.ChangePasswordHandler, true, &form_req.ChangePasswordReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}
