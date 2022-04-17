package service

import (
	"common/models"
	"time"
	"webapi/dao/form_req"
	"webapi/dao/form_resp"
	"webapi/dao/mongo"
	"webapi/internal/password"
	"webapi/internal/wrapper"
	"webapi/support"
	"webapi/utils"

	"github.com/globalsign/mgo/bson"

	"git.moresec.cn/moresec/go-common/mlog"
	"go.uber.org/zap"
)

// CreateUserHandler 创建用户
func CreateUserHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	traceCtx := ctx.Request().Context()
	req := reqBody.(*form_req.CreateUserReq)
	resp := form_resp.StatusResp{Status: support.StatusOK}
	if !utils.String.Compare(req.Password, req.Confirm) {
		support.SendApiErrorResponse(ctx, support.PasswordNotConfirm, 0)
		return nil
	}
	existQuery := bson.M{"username": req.Username, "role": req.Role}
	if mongo.User.IsExist(traceCtx, existQuery) {
		support.SendApiErrorResponse(ctx, support.UserIsExist, 0)
		return nil
	}
	passwordStrengthLevel := utils.Logic.GetPasswordStrength(req.Password)
	if passwordStrengthLevel == 0 {
		mlog.WithContext(traceCtx).Error("password too weak", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.PasswordStrengthFailed, 0)
		return nil
	}

	newUid := mongo.User.GetMaxUid(traceCtx)
	// 创建账户
	userDoc := models.User{
		ID:              bson.NewObjectId(),
		UID:             newUid,
		Role:            req.Role,
		UserName:        req.Username,
		Password:        password.MakePassword(req.Password),
		LastPwdChangeTm: time.Now(),
		LastLoginTm:     time.Now(),
		InsertTm:        time.Now(),
		UpdateTm:        time.Now(),
	}

	if err = mongo.User.Create(traceCtx, userDoc); err != nil {
		mlog.WithContext(traceCtx).Error("create user failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, support.CreateUserFailed, 0)
		return nil
	}
	//TOOD 创建用户HCL规范
	//var respDoc models.User
	//if respDoc, err = mongo.User.Get(traceCtx, newUid); err != nil {
	//	mlog.WithContext(traceCtx).Error("create user failed", zap.Error(err))
	//	support.SendApiErrorResponse(ctx, support.CreateUserFailed, 0)
	//	return nil
	//}
	//resp = form_json.UserResp{
	//	ID:               respDoc.ID.Hex(),
	//	UID:              respDoc.UID,
	//	Enable:           respDoc.Enable,
	//	UserType:         respDoc.UserType,
	//	UserName:         respDoc.UserName,
	//	Password:         respDoc.Password,
	//	PasswordStrength: respDoc.PasswordStrength,
	//	Mail:             respDoc.Mail,
	//	Mobile:           respDoc.Mobile,
	//	LastLoginIp:      respDoc.LastLoginIp,
	//	LastPwdChangeTm:  respDoc.LastPwdChangeTm,
	//	LastLoginTm:      respDoc.LastLoginTm,
	//	InsertTm:         respDoc.InsertTm,
	//	UpdateTm:         respDoc.UpdateTm,
	//	OpenID:           respDoc.OpenID,
	//}
	support.SendApiResponse(ctx, resp, "")
	return
}

// UserInfoHandler 获取用户信息
func UserInfoHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	return nil
}

func UserPasswordHandler(ctx *wrapper.Context, reqBody interface{}) (err error) {
	return nil
}
