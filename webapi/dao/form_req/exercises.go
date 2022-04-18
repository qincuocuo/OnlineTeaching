package form_req

type CreateExercisesReq struct {
	UserId            string          `json:"user_id"`             // 教师工号
	CourseId          int             `json:"course_id"`           // 课程号
	LearningContentId int             `json:"learning_content_id"` // 学习内容id
	Exercises         []ExercisesItem `json:"exercises"`           // 练习题
}

type ExercisesItem struct {
	Question string `json:"question"` // 题目
	Answer   string `json:"answer"`   // 答案
}

type ExercisesReq struct {
	Answer []string `json:"answer"` // 答案
}