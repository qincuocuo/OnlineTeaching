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
	endTm := time.Now().Add(time.Duration(req.RegisterTm) * time.Second)
	var userList = make([]string, 0)
	registerDoc := models.Register{
		ManagerId:  ctx.UserToken.UserId,
		ContentId:  req.ContentId,
		Finished:   userList,
		Unfinished: courseDoc.StudentId,
		CreateTime: time.Now(),
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
		resp.Count = len(registerDoc.Unfinished)
		for _, item := range registerDoc.Unfinished {
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
	}
	support.SendApiResponse(ctx, resp, "success")
	return
}

func RegisterHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.RegisterReq)
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
	if registerDoc.EndTime.Before(time.Now()) {
		support.SendApiErrorResponse(ctx, support.RegisterExpired, 0)
		return nil
	}
	if registerDoc.ManagerId != ctx.UserToken.UserId {
		support.SendApiErrorResponse(ctx, support.UserNoPermission, 0)
		return nil
	}
	finish := append(registerDoc.Finished, ctx.UserToken.UserId)
	unfinish := make([]string, 0)
	for _, s := range registerDoc.Unfinished {
		if s == ctx.UserToken.UserId {
			continue
		}
		unfinish = append(unfinish, s)
	}
	update := bson.M{"finished": finish, "unfinished": unfinish}
	err = mongo.Register.Update(traceCtx, query, update)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.JoinInRegisterFailed, 0)
		return nil
	}
	resp := form_resp.StatusResp{Status: "ok"}
	support.SendApiResponse(ctx, resp, "success")
	return
}
