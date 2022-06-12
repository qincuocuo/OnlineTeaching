package service

import (
	"git.moresec.cn/moresec/go-common/mbase"
	"git.moresec.cn/moresec/go-common/mlog"
	"github.com/globalsign/mgo/bson"
	"go.uber.org/zap"
	"strings"
	"time"
	"webapi/dao/form_req"
	"webapi/dao/form_resp"
	"webapi/dao/mongo"
	"webapi/internal/wrapper"
	"webapi/models"
	"webapi/support"
	"webapi/utils"
	utils2 "webapi/utils"
)

func GetClassListHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtc := ctx.Request().Context()
	req := reqBody.(*form_req.GetClassListReq)
	resp := form_resp.GetClassListResp{}
	query := bson.M{"grade": req.Grade}
	var courseDoc []models.Course
	courseDoc, err = mongo.Course.FindAll(traceCtc, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.GetCourseFailed, 0)
	}
	for _, course := range courseDoc {
		if !utils2.IsContainIntInSlice(course.Class, resp) {
			resp = append(resp, course.Class)
		}
	}
	support.SendApiResponse(ctx, resp, "success")
	return
}

func CreateCourseHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.CreateCourseReq)
	resp := form_resp.StatusResp{}
	if ctx.UserToken.Role == 2 {
		support.SendApiErrorResponse(ctx, support.UserNoPermission, 0)
		return nil
	}
	query := bson.M{"manager_id": ctx.UserToken.UserId}
	if req.Grade > 0 {
		query["grade"] = req.Grade
	}
	if req.Class > 0 {
		query["class"] = req.Class
	}
	if len(req.CourseName) > 0 {
		query["name"] = req.CourseName
	}
	var count int
	count, err = mongo.Course.FindCount(traceCtx, query)
	if count != 0 {
		support.SendApiErrorResponse(ctx, support.CourseIsExists, 0)
		return nil
	}
	courseInfo := models.Course{
		CourseId:    mongo.Course.GetMaxId(traceCtx),
		ManagerId:   ctx.UserToken.UserId,
		Name:        req.CourseName,
		Grade:       req.Grade,
		Class:       req.Class,
		TotalMember: 0,
		CreateTm:    time.Now(),
	}
	if err = mongo.Course.Create(traceCtx, courseInfo); err != nil {
		mlog.WithContext(traceCtx).Error("create course failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.CreateCourseFailed, 0)
		return nil
	}
	support.SendApiResponse(ctx, resp, "success")
	return
}

func CourseListHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.CourseListReq)
	resp := form_resp.CourseListResp{}
	query := bson.M{}
	if ctx.UserToken.Role == 1 {
		query["manager_id"] = ctx.UserToken.UserId
		if req.Grade > 0 {
			query["grade"] = req.Grade
		}
		if req.Class > 0 {
			query["class"] = req.Class
		}
	} else if ctx.UserToken.Role == 2 {
		//query["student_id"] = bson.M{"$elemMatch": bson.M{"$in": []string{ctx.UserToken.UserId}}}
		userDoc, err := mongo.User.FindByUserId(traceCtx, ctx.UserToken.UserId)
		if err != nil {
			support.SendApiErrorResponse(ctx, support.UserNotExist, 0)
			return err
		}

		query["grade"] = userDoc.Grade
		query["class"] = userDoc.Class
	}

	if len(req.CreateTm) > 0 {
		content := strings.Split(req.CreateTm, "--")
		if len(content) > 1 {
			beginTime, e1 := mbase.StrToTime(content[0])
			endTime, e2 := mbase.StrToTime(content[1])
			if e1 == nil && e2 == nil {
				query["create_tm"] = bson.M{"$gte": beginTime, "$lte": endTime}
			}
		}
	}
	if len(req.Search) > 0 {
		query["name"] = bson.M{"$regex": req.Search}
	}
	var sorter string
	sorter = "-create_tm"
	if len(req.OrderingGrade) > 0 {
		sorter = req.OrderingGrade
	} else if len(req.OrderingTotalMember) > 0 {
		sorter = req.OrderingTotalMember
	} else if len(req.OrderingCreateTm) > 0 {
		sorter = req.OrderingCreateTm
	}
	var courseDocs []models.Course
	courseDocs, err = mongo.Course.FindSortByLimitAndSkip(traceCtx, query, req.Page, req.PageSize, sorter)
	if err != nil {
		support.SendApiErrorResponse(ctx, "get course list fail", 0)
		return nil
	}
	for _, item := range courseDocs {
		msg := form_resp.CourseListItem{
			CourseId:    item.CourseId,
			CourseName:  item.Name,
			Grade:       item.Grade,
			Class:       item.Class,
			TotalMember: item.TotalMember,
			CreateTm:    utils.Time2String(item.CreateTm),
		}
		if ctx.UserToken.Role == 1 {
			subQuery := bson.M{
				"grade": item.Grade,
				"class": item.Class,
			}

			total, err := mongo.User.UserCountByGradeAndClass(traceCtx, subQuery)
			if err != nil {
				support.SendApiErrorResponse(ctx, "get user count failed", 0)
				return err
			}
			msg.TotalMember = total
		}

		resp.Result = append(resp.Result, msg)
	}
	resp.Count = len(resp.Result)
	support.SendApiResponse(ctx, resp, "success")
	return
}

func UpdateCourseHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.UpdateCourseReq)
	query := bson.M{"course_id": req.CourseId}
	var courseDoc models.Course
	courseDoc, err = mongo.Course.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.CourseNotExists, 0)
		return nil
	}
	if courseDoc.ManagerId != ctx.UserToken.UserId {
		support.SendApiErrorResponse(ctx, support.UserNoPermission, 0)
		return nil
	}
	upset := bson.M{"name": req.CourseName, "grade": req.Grade, "class": req.Class}
	err = mongo.Course.Update(traceCtx, query, upset)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.UpdateCourseFailed, 0)
		return nil
	}
	resp := form_resp.StatusResp{Status: "ok"}
	support.SendApiResponse(ctx, resp, "success")
	return
}

func DeleteCourseHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.DeleteCourseReq)
	query := bson.M{"course_id": req.CourseId}
	var courseDoc models.Course
	courseDoc, err = mongo.Course.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.CourseNotExists, 0)
		return nil
	}
	if courseDoc.ManagerId != ctx.UserToken.UserId {
		support.SendApiErrorResponse(ctx, support.UserNoPermission, 0)
		return nil
	}
	err = mongo.Course.Delete(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.DeleteCourseFailed, 0)
		return nil
	}
	resp := form_resp.StatusResp{Status: "ok"}
	support.SendApiResponse(ctx, resp, "success")
	return
}

func EnterCourseHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.EnterCourseReq)
	query := bson.M{"course_id": req.CourseId}
	var courseDoc models.Course
	courseDoc, err = mongo.Course.FindOne(traceCtx, query)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.CourseNotExists, 0)
		return nil
	}
	totalNum := courseDoc.TotalMember + 1
	stuIds := append(courseDoc.StudentId, ctx.UserToken.UserId)
	upset := bson.M{"student_id": stuIds, "total_member": totalNum}
	err = mongo.Course.Update(traceCtx, query, upset)
	if err != nil {
		support.SendApiErrorResponse(ctx, support.EnterCourseFailed, 0)
		return nil
	}
	resp := form_resp.StatusResp{Status: "ok"}
	support.SendApiResponse(ctx, resp, "success")
	return
}

func CourseInfoHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.CourseInfoReq)
	resp := form_resp.CourseInfoResp{}
	query := bson.M{}
	if req.Grade > 0 {
		query["grade"] = req.Grade
	}
	if req.Class > 0 {
		query["class"] = req.Class
	}
	if len(req.Search) > 0 {
		query["name"] = bson.M{"$regex": req.Search}
	}
	var courseDoc []models.Course
	courseDoc, err = mongo.Course.FindAll(traceCtx, query)
	if len(courseDoc) == 0 {
		support.SendApiErrorResponse(ctx, support.CourseNotFound, 0)
		return nil
	}
	for _, item := range courseDoc {
		msg := form_resp.CourseInfoItem{
			CourseId:   item.CourseId,
			CourseName: item.Name,
		}
		resp.Results = append(resp.Results, msg)
	}
	resp.Count = len(courseDoc)
	support.SendApiResponse(ctx, resp, "success")
	return
}
