package form_resp

type TalkInfoResp struct {
	Topic  string     `json:"topic"`  // 讨论话题
	Result []UserItem `json:"result"` // 讨论结果
}

type UserItem struct {
	UserId   string `json:"user_id"`   // 用户id
	UserName string `json:"user_name"` // 用户名
	TalkInfo string `json:"talk_info"` // 讨论信息
}
