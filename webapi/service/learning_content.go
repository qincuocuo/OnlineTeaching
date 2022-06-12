package service

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
	"webapi/dao/form_req"
	"webapi/dao/form_resp"
	"webapi/dao/mongo"
	"webapi/internal/chat"
	"webapi/internal/wrapper"
	"webapi/models"
	"webapi/support"
	"webapi/utils"
)

func LearningContentListHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.LearningContentListReq)
	resp := form_resp.LearningContentListResp{}
	query := bson.M{}
	if req.CourseId > 0 {
		query["course_id"] = req.CourseId
	}
	course, err := mongo.Course.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.CourseNotExists, 0)
		return nil
	}
	if len(req.Search) > 0 {
		query["title"] = bson.M{"$regex": req.Search}
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
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetLearningContentListFailed, 0)
		return nil
	}
	total, err := mongo.User.CountByGradeAndClass(traceCtx, course.Grade, course.Class)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.UserCountFailed, 0)
		return err
	}

	for _, item := range contentDoc {
		msg := form_resp.LearningContentItem{
			ContentId: item.ContentId,
			Content:   item.Title,
			Learned:   item.FinishedNum,
			Unlearned: total - item.FinishedNum,
		}

		register, _ := mongo.Register.FindOne(traceCtx, bson.M{"content_id": item.ContentId})
		if time.Now().Before(register.EndTime) {
			msg.Register = true
		}
		if utils.IsContainInSlice(ctx.UserToken.UserId, item.Finished) {
			msg.Registered = true
		}

		resp.Result = append(resp.Result, msg)
	}
	support.SendApiResponse(ctx, resp, "success")
	return
}

func CreateLearningContentHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.CreateLearningContentReq)
	resp := form_resp.StatusResp{Status: "ok"}
	file, fh, err := ctx.FormFile("file")
	if err != nil {
		support.SendApiErrorResponse(ctx, support.UploadLearningContentFailed, 0)
		return nil
	}
	defer file.Close()
	dir := fmt.Sprintf("%s/%d", "/workspace/data", req.CourseId)
	//filePath := fmt.Sprintf("%s/%d/%s", ".", req.ContentId, fh.Filename)

	if !utils.IsExistDir(dir) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			support.SendApiErrorResponse(ctx, support.UploadLearningContentFailed, 0)
			return nil
		}
	}
	filePath := fmt.Sprintf("%s/%s", dir, fh.Filename)

	dest, err := os.Create(filePath)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.UploadLearningContentFailed, 0)
		return nil
	}
	defer dest.Close()

	_, err = io.Copy(dest, file)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.UploadLearningContentFailed, 0)
		return nil
	}
	query := bson.M{"course_id": req.CourseId}
	_, err = mongo.Course.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.CourseNotExists, 0)
		return
	}

	learningContent := models.LearningContent{
		ContentId:   mongo.Content.GetMaxId(traceCtx),
		CourseId:    req.CourseId,
		Title:       req.Title,
		FileName:    fh.Filename,
		FinishedNum: 0,
		Finished:    nil,
	}
	err = mongo.Content.Create(traceCtx, learningContent)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.CreateLearningContentFailed, 0)
		return nil
	}
	support.SendApiResponse(ctx, resp, "success")
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

	users, err := mongo.User.GetByGradeAndClass(traceCtx, courseDoc.Grade, courseDoc.Class)
	if err != nil {
		support.SendApiErrorResponse(ctx, "获取学生失败", 0)
		return err
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
				Id:   item,
				Name: userDoc.UserName,
			}
			resp.StudentInfo = append(resp.StudentInfo, msg)
		}
	} else if req.Status == "unlearned" {
		resp.Count = len(users) - contentDoc.FinishedNum
		for _, user := range users {
			if utils.IsContainInSlice(user.UserId, contentDoc.Finished) {
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

func LearningHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.LearningReq)
	query := bson.M{"content_id": req.ContentId}
	var contentDoc models.LearningContent
	contentDoc, err = mongo.Content.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetLearningContentListFailed, 0)
		return nil
	}

	courseDoc, err := mongo.Course.FindOne(traceCtx, bson.M{"course_id": contentDoc.CourseId})
	if err != nil {
		support.SendApiErrorResponse(ctx, support.CourseNotFound, 0)
		return err
	}

	user, err := mongo.User.FindByUserId(traceCtx, ctx.UserToken.UserId)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.UserNotExist, 0)
		return nil
	}

	if ctx.UserToken.Role == 2 {
		if courseDoc.Grade != user.Grade && courseDoc.Class != user.Class {
			support.SendApiErrorResponse(ctx, support.UserNotInClass, 0)
			return nil
		}

		finished := contentDoc.Finished
		finishNum := contentDoc.FinishedNum

		if !utils.IsContainInSlice(user.UserId, finished) {

			finished = append(contentDoc.Finished, ctx.UserToken.UserId)
			finishNum = contentDoc.FinishedNum + 1

			upset := bson.M{"finished_num": finishNum, "finished": finished}
			err = mongo.Content.Update(traceCtx, query, upset)
			if err != nil {
				support.SendApiErrorResponse(ctx, support.UpdateContentFailed, 0)
				return nil
			}
		}
	}

	support.SendApiResponse(ctx, form_resp.StatusResp{Status: "ok"}, "success")
	return
}

func LearningContentHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.LearningContentReq)
	query := bson.M{"content_id": req.ContentId}
	var contentDoc models.LearningContent
	contentDoc, err = mongo.Content.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetLearningContentListFailed, 0)
		return nil
	}

	filePath := fmt.Sprintf("%s/%d/%s", "/workspace/data", contentDoc.CourseId, contentDoc.FileName)
	file, err := os.Open(filePath)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.OpenFileFailed, 0)
		return nil
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetLearningContentListFailed, 0)
		return nil
	}

	if strings.HasSuffix(filePath, ".pdf") {
		ctx.ResponseWriter().Header().Set("Content-type", "application/pdf")
		ctx.Write(data)
		return nil
	} else if strings.HasSuffix(filePath, ".mp4") {
		ctx.ResponseWriter().Header().Set("Content-Type", "video/mp4")
		ctx.Write(data)
		return nil
	}

	ctx.ResponseWriter().Header().Set("Content-Disposition", "attachment;filename="+url.QueryEscape(filepath.Base(filePath)))
	_, err = ctx.Write(data)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetLearningContentListFailed, 0)
		return nil
	}

	return
}

func StartChatHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}).Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
	if err != nil {
		support.SendApiErrorResponse(ctx, "创建websocket失败", 0)
		return err
	}

	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.StartChatReq)

	var contentDoc models.LearningContent

	query := bson.M{"content_id": req.ContentId}
	contentDoc, err = mongo.Content.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetLearningContentListFailed, 0)
		return nil
	}

	//go chat.Process(traceCtx, contentDoc.ContentId, req.UserId, conn)
	go chat.Process(context.TODO(), contentDoc.ContentId, req.UserId, conn)

	return
}
