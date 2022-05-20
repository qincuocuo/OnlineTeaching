package models

// Register 签到表
type Register struct {
	RegisterId int      `json:"register_id"` //签到任务id
	ManagerId  string   `json:"manager_id"`  //教师id
	CourseId   int      `json:"course_id"`   //课程id
	ContentId  int      `json:"content_id"`  //学习内容id
	Finished   []string `json:"finished"`    //完成签到学生id
	Unfinished []string `json:"unfinished"`  //未完成签到学生id
}

func (r *Register) CollectName() string {
	return "tb_register"
}
