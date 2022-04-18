package form_req

type NoticeReq struct {
	UserId            string          `json:"user_id"`             // 学生学号
	CourseId          int             `json:"course_id"`           // 课程号
	LearningContentId int             `json:"learning_content_id"` // 学习内容id
}
