package form_resp

type NoticeResp struct {
	Count  int          `json:"count"`
	Result []NoticeItem `json:"result"`
}

type NoticeItem struct {
	Type  int `json:"type"`  //通知类型 0-签到/1-课堂讨论/2-课后练习
	Count int `json:"count"` //通知数量
}
