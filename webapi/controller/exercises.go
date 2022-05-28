package controller

import (
	"webapi/dao/form_req"
	"webapi/internal/wrapper"
	"webapi/service"
	"webapi/support"
)

type Exercises struct{}

// CreateExercises
// @Summary 创建课后练习
// @Description create exercises
// @Tags exercises
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.CreateExercisesReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /v1/exercises/ [post]
// @Security ApiKeyAuth
func (e Exercises) CreateExercises(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.CreateExercisesHandler, true, &form_req.CreateExercisesReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// GetExercises
// @Summary 获取课后练习
// @Description get exercises
// @Tags exercises
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth query form_req.GetExercisesReq true "request data"
// @Success 200 {object} form_resp.GetExercisesResp "response data"
// @Router /v1/exercises/ [get]
// @Security ApiKeyAuth
func (e Exercises) GetExercises(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.GetExercisesHandler, true, &form_req.GetExercisesReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// Exercises
// @Summary 完成课后练习
// @Description exercises
// @Tags exercises
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.ExercisesReq true "request data"
// @Success 200 {object} form_resp.ExercisesResp "response data"
// @Router /v1/exercises/do/ [post]
// @Security ApiKeyAuth
func (e Exercises) Exercises(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.ExercisesHandler, true, &form_req.ExercisesReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}