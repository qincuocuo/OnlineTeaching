package form_req

type CreateTalkReq struct {
	UserId            string `json:"user_id"`             // 教师工号
	CourseId          int    `json:"course_id"`           // 课程号
	LearningContentId int    `json:"learning_content_id"` // 学习内容id
	Talk              string `json:"talk"`                // 讨论主题
}

type TalkInfoReq struct {
	UserId            string `json:"user_id"`             // 教师工号
	CourseId          int    `json:"course_id"`           // 课程号
	LearningContentId int    `json:"learning_content_id"` // 学习内容id
	TalkId            int    `json:"talk_id"`             // 讨论id
}

type TalkReq struct {
	UserId            string `json:"user_id"`             // 教师工号/学生学号
	CourseId          int    `json:"course_id"`           // 课程号
	LearningContentId int    `json:"learning_content_id"` // 学习内容id
	TalkId            int    `json:"talk_id"`             // 讨论id
}
