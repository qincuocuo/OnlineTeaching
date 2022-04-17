package form_req

type CreateLearningContentReq struct {
	UserId   string `json:"user_id"`   // 教师工号
	CourseId int    `json:"course_id"` // 课程id
	Title    string `json:"title"`     // 学习标题
}
