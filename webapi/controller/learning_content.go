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

func (l *LearningController) LearningListContent(ctx *wrapper.Context) {
}
