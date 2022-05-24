package models

type LearningContent struct {
	ContentId     int      `json:"content_id"`     //学习内容id
	CourseId      int      `json:"course_id"`      //课程id
	Title         string   `json:"title"`          //标题
	FinishedNum   int      `json:"finished_num"`   //已学人数
	UnfinishedNum int      `json:"unfinished_num"` //未学人数
	Finished      []string `json:"finished"`       //已学学生id
	Unfinished    []string `json:"unfinished"`     //未学学生id
}

func (l *LearningContent) CollectName() string {
	return "tb_learning_content"
}
