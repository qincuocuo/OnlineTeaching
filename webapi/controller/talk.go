package controller

import (
	"webapi/dao/form_req"
	"webapi/internal/wrapper"
	"webapi/service"
	"webapi/support"
)

type Talk struct{}

// CreateTalk
// @Summary 创建讨论话题
// @Description create talk
// @Tags talk
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.CreateTalkReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /v1/talk/ [post]
// @Security ApiKeyAuth
func (t *Talk) CreateTalk(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.CreateTalkHandler, true, form_req.CreateTalkReq{}, wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// TalkInfo
// @Summary 获取讨论话题详情
// @Description get talk info
// @Tags talk
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.TalkInfoReq true "request data"
// @Success 200 {object} form_resp.TalkInfoResp "response data"
// @Router /v1/talk/ [get]
// @Security ApiKeyAuth
func (t *Talk) TalkInfo(ctx *wrapper.Context) {}

// Talk
// @Summary 参与讨论
// @Description join talk
// @Tags talk
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.TalkReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /v1/talk/do/ [post]
// @Security ApiKeyAuth
func (t *Talk) Talk(ctx *wrapper.Context) {}
