package models

import "time"

type Exercises struct {
	ContentId int             `bson:"content_id"` //学习内容id
	Questions []QuestionsItem `bson:"questions"`  //习题
	CreateTm  time.Time       `bson:"create_tm"`  //创建时间
}

type QuestionsItem struct {
	Id       int      `bson:"id"`       //习题id
	Type     int      `bson:"type"`     //0-选择 1-判断
	Question string   `bson:"question"` //问题
	Answer   string   `bson:"answer"`   //答案
	Options  []string `bson:"options"`  //选项
}

func (e Exercises) CollectName() string {
	return "tb_exercises"
}
