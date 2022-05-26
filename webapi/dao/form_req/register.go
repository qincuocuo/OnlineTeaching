package form_req

import "time"

type CreateRegisterReq struct {
	ContentId  int       `json:"content_id"`  // 学习内容id
	CreateTm   time.Time `json:"create_tm"`   // 创建时间
	RegisterTm int       `json:"register_tm"` // 签到时间限制，默认为2分钟
}

type RegisterResultReq struct {
	ContentId      int    `json:"content_id"`      // 学习内容id
	RegisterResult string `json:"register_result"` // 签到结果，unfinished-未签到/finished-已签到
}

type RegisterReq struct {
	ContentId int `json:"content_id"` // 学习内容id
}
