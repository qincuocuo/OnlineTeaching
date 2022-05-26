package controller

import "webapi/internal/wrapper"

type Notice struct{}

// Notice
// @Summary 通知
// @Description notice
// @Tags notice
// @Accept x-www-form-urlencoded
// @Produce json
// @Param auth body form_req.NoticeReq true "request data"
// @Success 200 {object} form_resp.NoticeResp "response data"
// @Router /v1/notice/ [get]
// @Security ApiKeyAuth
func (n Notice) Notice(ctx *wrapper.Context) {}
