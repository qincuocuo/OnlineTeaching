package controller

import (
	"webapi/dao/form_req"
	"webapi/internal/wrapper"
	"webapi/service"
	"webapi/support"
)

type CourseController struct {
}

// GetClassList
// @Summary 获取班级列表
// @Description get class list
// @Tags course
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth query form_req.GetClassListReq true "request data"
// @Success 200 {object} form_resp.GetClassListResp "response data"
// @Router /v1/course/class [get]
// @Security ApiKeyAuth
func (c CourseController) GetClassList(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.GetClassListHandler, true, &form_req.GetClassListReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
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
	wrapper.ApiWrapper(ctx, service.CreateCourseHandler, true, &form_req.CreateCourseReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}

// CourseList
// @Summary 课程列表
// @Description course list
// @Tags course
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth query form_req.CourseListReq true "request data"
// @Success 200 {object} form_resp.CourseListResp "response data"
// @Router /v1/course/ [get]
// @Security ApiKeyAuth
func (c CourseController) CourseList(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.CourseListHandler, true, &form_req.CourseListReq{PageSize: 10}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
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
	wrapper.ApiWrapper(ctx, service.UpdateCourseHandler, true, &form_req.UpdateCourseReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
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
	wrapper.ApiWrapper(ctx, service.DeleteCourseHandler, true, &form_req.DeleteCourseReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_FORM})
}

// EnterCourse
// @Summary 学生加入课程
// @Description enter course
// @Tags course
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.EnterCourseReq true "request data"
// @Success 200 {object} form_resp.StatusResp "response data"
// @Router /v1/course/enter [post]
// @Security ApiKeyAuth
func (c CourseController) EnterCourse(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, service.EnterCourseHandler, true, &form_req.EnterCourseReq{}, &wrapper.ApiConfig{ReqType: support.CHECKTYPE_JSON})
}
