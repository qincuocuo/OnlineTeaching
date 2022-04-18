package controller

import "webapi/internal/wrapper"

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
func (r *Register) CreateRegister(ctx *wrapper.Context) {}

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
func (r *Register) RegisterResult(ctx *wrapper.Context) {}

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
func (r *Register) Register(ctx *wrapper.Context) {}