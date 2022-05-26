package models

import "time"

type Talk struct {
	ContentId int       `json:"content_id"` //学习内容id
	Topic     string    `json:"topic"`      //讨论话题
	CreateTm  time.Time `json:"create_tm"`  //创建时间
}

func (t Talk) CollectName() string {
	return "tb_talk"
}

type TalkRecord struct {
	ContentId int       `json:"content_id"` //学习内容id
	UserId    int       `json:"user_id"`    //用户id
	Text      string    `json:"text"`       //用户讨论内容
	CreateTm  time.Time `json:"create_tm"`  //参与讨论时间
}

func (t TalkRecord) CollectName() string {
	return "tb_talk_record"
}
