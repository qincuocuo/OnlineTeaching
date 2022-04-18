package form_req

import "time"

type CreateRegisterReq struct {
	UserId            string    `json:"user_id"`             // 教师工号
	CourseId          int       `json:"course_id"`           // 课程号
	LearningContentId int       `json:"learning_content_id"` // 学习内容id
	CreateTm          time.Time `json:"create_tm"`           // 创建时间
	RegisterTm        time.Time `json:"register_tm"`         // 签到时间限制，默认为2分钟
}

type RegisterResultReq struct {
	UserId            string `json:"user_id"`             // 教师工号
	CourseId          int    `json:"course_id"`           // 课程号
	LearningContentId int    `json:"learning_content_id"` // 学习内容id
	RegisterResult    int    `json:"register_result"`     // 签到结果，unfinished-未签到/finished-已签到
}

type RegisterReq struct {
	UserId            string `json:"user_id"`             // 学生学号
	CourseId          int    `json:"course_id"`           // 课程号
	LearningContentId int    `json:"learning_content_id"` // 学习内容id
}
