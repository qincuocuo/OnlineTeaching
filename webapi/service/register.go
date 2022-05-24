package service

import (
	"github.com/globalsign/mgo/bson"
	"webapi/dao/form_req"
	"webapi/dao/form_resp"
	"webapi/dao/mongo"
	"webapi/internal/utils"
	"webapi/internal/wrapper"
	"webapi/models"
	"webapi/support"
)

func CreateRegisterHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.CreateRegisterReq)
	query := bson.M{}
	if req.ContentId > 0 {
		query["content_id"] = req.ContentId
	}
	var contentDoc models.LearningContent
	contentDoc, err = mongo.Content.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetLearningContentListFailed, 0)
		return nil
	}
	var courseDoc models.Course
	courseDoc, err = mongo.Course.FindOne(traceCtx, bson.M{"course_id": contentDoc.CourseId})
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetCourseInfoFailed, 0)
		return nil
	}
	if courseDoc.ManagerId != ctx.UserToken.UserId {
		support.SendApiErrorResponse(ctx, support.UserNoPermission, 0)
		return nil
	}
	maxId := mongo.Register.GetMaxId(traceCtx)
	var userList = make([]string, 0)
	registerDoc := models.Register{
		RegisterId: maxId,
		ManagerId: ctx.UserToken.UserId,
		ContentId: req.ContentId,
		Finished: userList,
		Unfinished: userList,
	}
	err = mongo.Register.Create(traceCtx, registerDoc)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.CreateRegisterFailed, 0)
		return nil
	}
	resp := form_resp.StatusResp{Status: "ok"}
	support.SendApiResponse(ctx, resp, "success")
	return
}

func RegisterHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.CreateRegisterReq)
	query := bson.M{}
	if req.ContentId > 0 {
		query["content_id"] = req.ContentId
	}
	var contentDoc models.LearningContent
	contentDoc, err = mongo.Content.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetLearningContentListFailed, 0)
		return nil
	}
	var courseDoc models.Course
	courseDoc, err = mongo.Course.FindOne(traceCtx, bson.M{"course_id": contentDoc.CourseId})
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetCourseInfoFailed, 0)
		return nil
	}
	if !utils.IsContainInSlice(ctx.UserToken.UserId, courseDoc.StudentId) {
		support.SendApiErrorResponse(ctx, support.UserNoPermission, 0)
		return nil
	}

	resp := form_resp.StatusResp{Status: "ok"}
	support.SendApiResponse(ctx, resp, "success")
	return
}