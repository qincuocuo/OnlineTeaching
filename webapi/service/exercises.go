package service

import (
	"github.com/globalsign/mgo/bson"
	"math/rand"
	"time"
	"webapi/dao/form_req"
	"webapi/dao/form_resp"
	"webapi/dao/mongo"
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

	exist := mongo.Exercises.IsExist(traceCtx, query)
	if exist {
		questions := make([]models.QuestionsItem, 0)
		for _, item := range req.Exercises {
			msg := models.QuestionsItem{
				Id:       rand.Int(),
				Type:     item.Type,
				Question: item.Question,
				Answer:   item.Answer,
				Options:  item.Options,
			}
			questions = append(questions, msg)
		}

		err = mongo.Exercises.Update(traceCtx, query, bson.M{
			"create_tm": time.Now(),
			"questions": questions,
		})
		if err != nil {
			support.SendApiErrorResponse(ctx, support.CreateCourseFailed, 0)
			return err
		}
	} else {
		var exercisesDoc models.Exercises

		exercisesDoc.ContentId = req.ContentId
		exercisesDoc.CreateTm = time.Now()

		for _, item := range req.Exercises {
			msg := models.QuestionsItem{
				Id:       rand.Int(),
				Type:     item.Type,
				Question: item.Question,
				Answer:   item.Answer,
				Options:  item.Options,
			}
			exercisesDoc.Questions = append(exercisesDoc.Questions, msg)
		}

		err = mongo.Exercises.Create(traceCtx, exercisesDoc)
		if err != nil {
			support.SendApiErrorResponse(ctx, support.CreateExercisesFailed, 0)
			return nil
		}
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

	if courseDoc.ManagerId != ctx.UserToken.UserId && ctx.UserToken.Role == 1 {
		support.SendApiErrorResponse(ctx, support.UserNoPermission, 0)
		return nil
	}

	user, err := mongo.User.FindByUserId(traceCtx, ctx.UserToken.UserId)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.UserNotExist, 0)
		return err
	}
	if ctx.UserToken.Role == 2 && courseDoc.Grade != user.Grade && courseDoc.Class != user.Class {
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
			Id:       item.Id,
			Type:     item.Type,
			Question: item.Question,
			Options:  item.Options,
		}
		if ctx.UserToken.Role == 1 {
			msg.Answer = item.Answer
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
		var flag bool
		for _, question := range exercisesDoc.Questions {
			if answer.Id == question.Id {
				if answer.Answer == question.Answer {
					flag = true
					break
				}
			}
		}
		msg := form_resp.ExerciseResult{
			Id:      answer.Id,
			Correct: flag,
			Answer:  answer.Answer,
		}
		resp.Results = append(resp.Results, msg)
	}

	support.SendApiResponse(ctx, resp, "success")
	return
}
