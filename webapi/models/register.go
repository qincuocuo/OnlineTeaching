package models

import "time"

// Register 签到表
type Register struct {
	ContentId  int       `bson:"content_id"`  //学习内容id
	ManagerId  string    `bson:"manager_id"`  //教师id
	CreateTime time.Time `bson:"create_time"` //签到创建时间
	EndTime    time.Time `bson:"end_time"`    //签到结束时间
	Finished   []string  `bson:"finished"`    //完成签到学生id
}

func (r *Register) CollectName() string {
	return "tb_register"
}
