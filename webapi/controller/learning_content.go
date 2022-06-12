package controller

import (
	"webapi/dao/form_req"
	"webapi/internal/wrapper"
	"webapi/service"
	"webapi/support"
)

type LearningController struct {
}

// CreateLearningContent
// @Summary 新增学习内容
// @Description delete course
// @Tags learning
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.CreateLearningContentReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /v1/learning_content/ [post]
// @Security ApiKeyAuth
func (l LearningController) CreateLearningContent(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.CreateLearningContentHandler, true, &form_req.CreateLearningContentReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// LearningContentList
// @Summary 学习内容列表
// @Description learning list
// @Tags learning
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.LearningContentListReq true "request data"
// @Success 200 {object} form_resp.LearningContentListResp "response data"
// @Router /v1/learning_content/ [get]
// @Security ApiKeyAuth
func (l LearningController) LearningContentList(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.LearningContentListHandler, true, &form_req.LearningContentListReq{PageSize: 10}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// LearningResult
// @Summary 查看学习情况
// @Description get learning result
// @Tags learning
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.LearningResultReq true "request data"
// @Success 200 {object} form_resp.LearningResultResp "response data"
// @Router /v1/learning_content/result/ [get]
// @Security ApiKeyAuth
func (l LearningController) LearningResult(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.LearningResultHandler, true, &form_req.LearningResultReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// Learning
// @Summary 进入学习
// @Description  learning
// @Tags learning
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth query form_req.LearningReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /v1/learning_content/learning/ [get]
// @Security ApiKeyAuth
func (l LearningController) Learning(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.LearningHandler, true, &form_req.LearningReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

func (l LearningController) LearningContent(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.LearningContentHandler, true, &form_req.LearningContentReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

func (l LearningController) StartChat(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.StartChatHandler, true, &form_req.StartChatReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}
