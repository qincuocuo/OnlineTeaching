package service

import (
	"github.com/globalsign/mgo/bson"
	"webapi/dao/form_req"
	"webapi/dao/form_resp"
	"webapi/dao/mongo"
	"webapi/internal/wrapper"
	"webapi/models"
	"webapi/support"
)

func LearningContentListHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.LearningContentListReq)
	resp := form_resp.LearningContentListResp{}
	query := bson.M{}
	if req.CourseId > 0 {
		query["course_id"] = req.CourseId
	}
	_, err = mongo.Course.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.CourseNotExists, 0)
		return nil
	}
	if len(req.Search) > 0 {
		query["title"] = req.Search
	}
	sorter := "-learned"
	if len(req.OrderingLearned) > 0 {
		sorter = req.OrderingLearned
	} else if len(req.OrderingUnLearned) > 0 {
		sorter = req.OrderingUnLearned
	}
	resp.Count, err = mongo.Content.FindCount(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetLearningContentListFailed, 0)
		return nil
	}
	var contentDoc []models.LearningContent
	contentDoc, err = mongo.Content.FindSortByLimitAndSkip(traceCtx, query, req.Page, req.PageSize, sorter)
	if err != nil{
		support.SendApiErrorResponse(ctx, support.GetLearningContentListFailed, 0)
		return nil
	}
	for _, item := range contentDoc {
		msg := form_resp.LearningContentItem{
			ContentId: item.ContentId,
			Content: item.Title,
			Learned: item.FinishedNum,
			Unlearned: item.UnfinishedNum,
		}
		resp.Result = append(resp.Result, msg)
	}
	support.SendApiResponse(ctx, resp, "success")
	return
}

func CreateLearningContentHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	return
}

func LearningResultHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.LearningResultReq)
	resp := form_resp.LearningResultResp{}
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
	if req.Status == "learned" {
		resp.Count = contentDoc.FinishedNum
		for _, item := range contentDoc.Finished {
			var userDoc models.User
			userDoc, err = mongo.User.FindByUserId(traceCtx, item)
			if err != nil {
				continue
			}
			msg := form_resp.StudentItem{
				Id: item,
				Name: userDoc.UserName,
			}
			resp.StudentInfo = append(resp.StudentInfo, msg)
		}
	} else if req.Status == "unlearned" {
		resp.Count = contentDoc.UnfinishedNum
		for _, item := range contentDoc.Unfinished {
			var userDoc models.User
			userDoc, err = mongo.User.FindByUserId(traceCtx, item)
			if err != nil {
				continue
			}
			msg := form_resp.StudentItem{
				Id: item,
				Name: userDoc.UserName,
			}
			resp.StudentInfo = append(resp.StudentInfo, msg)
		}
	}
	support.SendApiResponse(ctx, resp, "success")
	return
}
