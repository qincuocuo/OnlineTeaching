package form_req

type CreateRegisterReq struct {
	ContentId  int `json:"content_id" validate:"required"`  // 学习内容id
	RegisterTm int `json:"register_tm" validate:"required"` // 签到时间限制，默认为2分钟
}

type RegisterResultReq struct {
	ContentId      int    `form:"content_id" json:"content_id" validate:"required"`      // 学习内容id
	RegisterResult string `form:"register_result" json:"register_result" validate:"required"` // 签到结果，unfinished-未签到/finished-已签到
}

type RegisterReq struct {
	ContentId int `json:"content_id" validate:"required"` // 学习内容id
}
