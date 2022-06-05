package service

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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
	_, err = mongo.Course.FindOne(traceCtx, query)
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
	for _, item := range contentDoc {
		msg := form_resp.LearningContentItem{
			ContentId: item.ContentId,
			Content:   item.Title,
			Learned:   item.FinishedNum,
			Unlearned: item.UnfinishedNum,
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
	var courseDoc models.Course
	courseDoc, err = mongo.Course.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.CourseNotExists, 0)
		return
	}

	learningContent := models.LearningContent{
		ContentId:     mongo.Content.GetMaxId(traceCtx),
		CourseId:      req.CourseId,
		Title:         fh.Filename,
		FinishedNum:   0,
		UnfinishedNum: courseDoc.TotalMember,
		Finished:      nil,
		Unfinished:    courseDoc.StudentId,
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
		resp.Count = contentDoc.UnfinishedNum
		for _, item := range contentDoc.Unfinished {
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

	filePath := fmt.Sprintf("%s/%d/%s", "/workspace/data", contentDoc.CourseId, contentDoc.Title)
	file, err := os.Open(filePath)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.OpenFileFailed, 0)
		return nil
	}
	defer file.Close()

	_, err = mongo.User.FindByUserId(traceCtx, ctx.UserToken.UserId)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.UserNotExist, 0)
		return nil
	}

	finished := contentDoc.Finished
	unfinished := contentDoc.Unfinished
	finishNum := contentDoc.FinishedNum
	unfinishedNum := contentDoc.UnfinishedNum
	if ctx.UserToken.Role == 2 {
		finished = append(contentDoc.Finished, ctx.UserToken.UserId)
		unfinished = make([]string, 0)
		for _, s := range contentDoc.Unfinished {
			if s == ctx.UserToken.UserId {
				continue
			}
			unfinished = append(unfinished, s)
		}
		finishNum = contentDoc.FinishedNum + 1
		unfinishedNum = contentDoc.UnfinishedNum - 1
	}

	upset := bson.M{"finished_num": finishNum, "unfinished_num": unfinishedNum, "finished": finished, "unfinished": unfinished}
	err = mongo.Content.Update(traceCtx, query, upset)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.UpdateContentFailed, 0)
		return nil
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetLearningContentListFailed, 0)
		return nil
	}

	if strings.HasSuffix(filePath, ".pdf") {
		ctx.Write(data)
		return nil
	} else if strings.HasSuffix(filePath, ".mp4") {
		ctx.ResponseWriter().Header().Set("Content-Type", "video/mp4")
		ctx.Write(data)
		return nil
	}

	ctx.ResponseWriter().Header().Set("Content-Disposition", "attachment;filename="+url.QueryEscape(filepath.Base(filePath)))
	_, err = ctx.Write(data)

	_, err = ctx.Write(data)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetLearningContentListFailed, 0)
		return nil
	}

	return
}

func StartChatHandler(ctx *wrapper.Context, conn *websocket.Conn, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.StartChatReq)

	var contentDoc models.LearningContent

	query := bson.M{"content_id": req.ContentId}
	contentDoc, err = mongo.Content.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetLearningContentListFailed, 0)
		return nil
	}

	chat.Process(traceCtx, contentDoc.ContentId, ctx.UserToken.UserId, conn)

	return
}
