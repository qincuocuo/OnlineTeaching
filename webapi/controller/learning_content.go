package controller

import "webapi/internal/wrapper"

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
func (l *LearningController) CreateLearningContent(ctx *wrapper.Context) {
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
func (l *LearningController) LearningContentList(ctx *wrapper.Context) {
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
func (l *LearningController) LearningResult(ctx *wrapper.Context) {
}

// Learning
// @Summary 进入学习
// @Description  learning
// @Tags learning
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.LearningReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /v1/learning_content/learning/ [post]
// @Security ApiKeyAuth
func (l *LearningController) Learning(ctx *wrapper.Context) {
}
