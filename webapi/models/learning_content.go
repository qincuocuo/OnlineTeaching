package models

type LearningContent struct {
	ContentId   int      `bson:"content_id"`   //学习内容id
	CourseId    int      `bson:"course_id"`    //课程id
	Title       string   `bson:"title"`        //标题
	FinishedNum int      `bson:"finished_num"` //已学人数
	Finished    []string `bson:"finished"`     //已学学生id
}

func (l *LearningContent) CollectName() string {
	return "tb_learning_content"
}
