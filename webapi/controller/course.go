package controller

import (
	"webapi/internal/wrapper"
)

type CourseController struct {
}

// CreateCourse
// @Summary 新建课程
// @Description create course
// @Tags course
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.CreateCourseReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /v1/course/ [post]
// @Security ApiKeyAuth
func (c CourseController) CreateCourse(ctx *wrapper.Context) {
}

// CourseList
// @Summary 新建课程
// @Description course list
// @Tags course
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.CourseListReq true "request data"
// @Success 200 {object} form_resp.CourseListResp "response data"
// @Router /v1/course/ [post]
// @Security ApiKeyAuth
func (c CourseController) CourseList(ctx *wrapper.Context) {
}

// UpdateCourse
// @Summary 编辑课程信息
// @Description update course info
// @Tags course
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.UpdateCourseReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /v1/course/update/ [post]
// @Security ApiKeyAuth
func (c CourseController) UpdateCourse(ctx *wrapper.Context) {
}

// DeleteCourse
// @Summary 删除课程
// @Description delete course
// @Tags course
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.DeleteCourseReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /v1/course/ [delete]
// @Security ApiKeyAuth
func (c CourseController) DeleteCourse(ctx *wrapper.Context) {
}

