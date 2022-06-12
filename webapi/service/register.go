package service

import (
	"github.com/globalsign/mgo/bson"
	"time"
	"webapi/dao/form_req"
	"webapi/dao/form_resp"
	"webapi/dao/mongo"
	"webapi/internal/wrapper"
	"webapi/models"
	"webapi/support"
	"webapi/utils"
)

func CreateRegisterHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.CreateRegisterReq)
	query := bson.M{}
	if req.ContentId > 0 {
		query["content_id"] = req.ContentId
	}
	if mongo.Register.IsExists(traceCtx, req.ContentId) {
		support.SendApiErrorResponse(ctx, support.RegisterExists, 0)
		return nil
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
	startTm := time.Now().Add(time.Duration(8) * time.Hour)
	endTm := startTm.Add(time.Duration(req.RegisterTm) * time.Minute)
	var userList = make([]string, 0)
	registerDoc := models.Register{
		ManagerId:  ctx.UserToken.UserId,
		ContentId:  req.ContentId,
		Finished:   userList,
		CreateTime: startTm,
		EndTime:    endTm,
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

func RegisterResultHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.RegisterResultReq)
	resp := form_resp.RegisterResultResp{}
	query := bson.M{}
	if req.ContentId > 0 {
		query["content_id"] = req.ContentId
	}
	var registerDoc models.Register
	registerDoc, err = mongo.Register.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.RegisterNotFount, 0)
		return nil
	}

	contentDoc, err := mongo.Content.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.LearningContentNotFound, 0)
		return err
	}
	courseDoc, err := mongo.Course.FindOne(traceCtx, bson.M{"course_id": contentDoc.CourseId})
	if err != nil {
		support.SendApiErrorResponse(ctx, support.CourseNotExists, 0)
		return err
	}

	students, err := mongo.User.GetByGradeAndClass(traceCtx, courseDoc.Grade, courseDoc.Class)
	if err != nil {
		support.SendApiErrorResponse(ctx, "get user failed", 0)
		return err
	}

	if req.RegisterResult == "finished" {
		resp.Count = len(registerDoc.Finished)
		for _, item := range registerDoc.Finished {
			var userDoc models.User
			userDoc, err = mongo.User.FindByUserId(traceCtx, item)
			if err != nil {
				continue
			}
			msg := form_resp.StudentItem{
				Id:   item,
				Name: userDoc.UserName,
			}
			resp.StudentInfo = append(resp.StudentInfo, msg)
		}
	} else if req.RegisterResult == "unfinished" {
		resp.Count = len(students) - len(registerDoc.Finished)
		for _, user := range students {
			if utils.IsContainInSlice(user.UserId, registerDoc.Finished) {
				continue
			}

			msg := form_resp.StudentItem{
				Id:   user.UserId,
				Name: user.UserName,
			}
			resp.StudentInfo = append(resp.StudentInfo, msg)
		}
	}
	support.SendApiResponse(ctx, resp, "success")
	return
}

func UserRegisterInfoHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {

	return
}

func RegisterHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.RegisterReq)
	query := bson.M{}
	if req.ContentId > 0 {
		query["content_id"] = req.ContentId
	}

	if ctx.UserToken.Role != 2 {
		support.SendApiErrorResponse(ctx, support.UserNoPermission, 0)
		return nil
	}

	var contentDoc models.LearningContent
	contentDoc, err = mongo.Content.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.LearningContentNotFound, 0)
		return nil
	}
	var courseDoc models.Course
	courseDoc, err = mongo.Course.FindOne(traceCtx, bson.M{"course_id": contentDoc.CourseId})
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetCourseInfoFailed, 0)
		return nil
	}

	user, err := mongo.User.FindByUserId(traceCtx, ctx.UserToken.UserId)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.UserNotExist, 0)
		return
	}

	if courseDoc.Grade != user.Grade && courseDoc.Class != user.Class {
		support.SendApiErrorResponse(ctx, support.UserNoPermission, 0)
		return nil
	}

	var registerDoc models.Register
	registerDoc, err = mongo.Register.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.RegisterNotFount, 0)
		return nil
	}

	if utils.IsContainInSlice(ctx.UserToken.UserId, registerDoc.Finished) {
		support.SendApiErrorResponse(ctx, support.RegisterFinished, 0)
		return nil
	}

	if registerDoc.EndTime.Before(time.Now()) {
		support.SendApiErrorResponse(ctx, support.RegisterExpired, 0)
		return nil
	}

	if utils.IsContainInSlice(ctx.UserToken.UserId, registerDoc.Finished) {
		finish := append(registerDoc.Finished, ctx.UserToken.UserId)
		update := bson.M{"finished": finish}

		err = mongo.Register.Update(traceCtx, query, update)
		if err != nil {
			support.SendApiErrorResponse(ctx, support.JoinInRegisterFailed, 0)
			return nil
		}
	}
	resp := form_resp.StatusResp{Status: "ok"}
	support.SendApiResponse(ctx, resp, "success")
	return
}
