package controller

import (
	"webapi/dao/form_req"
	"webapi/internal/wrapper"
	"webapi/service"
	"webapi/support"
)

type Register struct {
}

// CreateRegister
// @Summary 创建签到任务
// @Description create register
// @Tags register
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.CreateRegisterReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /v1/register/ [post]
// @Security ApiKeyAuth
func (r Register) CreateRegister(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.CreateRegisterHandler, true, form_req.CreateRegisterReq{}, wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// RegisterResult
// @Summary 签到结果
// @Description register result
// @Tags register
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.RegisterResultReq true "request data"
// @Success 200 {object} form_resp.RegisterResultResp "response data"
// @Router /v1/register/result/ [get]
// @Security ApiKeyAuth
func (r Register) RegisterResult(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.RegisterResultHandler, true, form_req.RegisterResultReq{}, wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// Register
// @Summary 学生参与签到
// @Description register
// @Tags register
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.RegisterReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /v1/register/do [post]
// @Security ApiKeyAuth
func (r Register) Register(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.RegisterHandler, true, form_req.RegisterReq{}, wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}
