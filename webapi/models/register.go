package models

import "time"

// Register 签到表
type Register struct {
	ContentId  int       `json:"content_id"`  //学习内容id
	ManagerId  string    `json:"manager_id"`  //教师id
	CreateTime time.Time `json:"create_time"` //签到创建时间
	EndTime    time.Time `json:"end_time"`    //签到结束时间
	Finished   []string  `json:"finished"`    //完成签到学生id
	Unfinished []string  `json:"unfinished"`  //未完成签到学生id
}

func (r *Register) CollectName() string {
	return "tb_register"
}
