package service

import (
	"github.com/globalsign/mgo/bson"
	"math/rand"
	"time"
	"webapi/dao/form_req"
	"webapi/dao/form_resp"
	"webapi/dao/mongo"
	"webapi/internal/utils"
	"webapi/internal/wrapper"
	"webapi/models"
	"webapi/support"
)

func CreateExercisesHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.CreateExercisesReq)
	query := bson.M{}
	if ctx.UserToken.Role != 1 {
		support.SendApiErrorResponse(ctx, support.UserNoPermission, 0)
		return nil
	}
	if req.ContentId > 0 {
		query["content_id"] = req.ContentId
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
	if courseDoc.ManagerId != ctx.UserToken.UserId {
		support.SendApiErrorResponse(ctx, support.UserNoPermission, 0)
		return nil
	}
	var exercisesDoc models.Exercises
	exercisesDoc.ContentId = req.ContentId
	exercisesDoc.CreateTm = time.Now()
	for _, item := range req.Exercises {
		msg := models.QuestionsItem{
			Id: rand.Int(),
			Type: item.Type,
			Question: item.Question,
			Answer: item.Answer,
			Options: item.Options,
		}
		exercisesDoc.Questions = append(exercisesDoc.Questions, msg)
	}
	err = mongo.Exercises.Create(traceCtx, exercisesDoc)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.CreateExercisesFailed, 0)
		return nil
	}
	resp := form_resp.StatusResp{Status: "ok"}
	support.SendApiResponse(ctx, resp, "success")
	return
}

func GetExercisesHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.GetExercisesReq)
	resp := form_resp.GetExercisesResp{}
	query := bson.M{}
	if req.ContentId > 0 {
		query["content_id"] = req.ContentId
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
	if courseDoc.ManagerId != ctx.UserToken.UserId && !utils.IsContainInSlice(ctx.UserToken.UserId, courseDoc.StudentId) {
		support.SendApiErrorResponse(ctx, support.UserNoPermission, 0)
		return nil
	}
	var exercisesDoc models.Exercises
	exercisesDoc, err = mongo.Exercises.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.ExercisesNotExists, 0)
		return nil
	}
	resp.Count = len(exercisesDoc.Questions)
	for _, item := range exercisesDoc.Questions {
		msg := form_resp.ExercisesItem{
			Id: item.Id,
			Type: item.Type,
			Question: item.Question,
			Options: item.Options,
		}
		resp.Exercises = append(resp.Exercises, msg)
	}
	support.SendApiResponse(ctx, resp, "success")
	return
}

func ExercisesHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.ExercisesReq)
	resp := form_resp.ExercisesResp{}
	query := bson.M{}
	if req.ContentId > 0 {
		query["content_id"] = req.ContentId
	}
	var exercisesDoc models.Exercises
	exercisesDoc, err = mongo.Exercises.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.ExercisesNotExists, 0)
		return nil
	}
	for _, answer := range req.Answers {
		for _, question := range exercisesDoc.Questions {
			flag := false
			if answer.Id == question.Id {
				if answer.Answer == question.Answer {
					flag = true
				}
			}
			msg := form_resp.ExerciseResult{
				Id: answer.Id,
				Correct: flag,
				Answer: question.Answer,
			}
			resp.Results = append(resp.Results, msg)
		}
	}
	support.SendApiResponse(ctx, resp, "success")
	return
}