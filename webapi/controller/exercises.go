package controller

import "webapi/internal/wrapper"

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
func (e *Exercises) CreateExercises(ctx *wrapper.Context) {}

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
func (e *Exercises) Exercises(ctx *wrapper.Context) {}