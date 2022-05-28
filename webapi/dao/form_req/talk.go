package form_req

type CreateTalkReq struct {
	ContentId int    `json:"content_id"` // 学习内容id
	Talk      string `json:"talk"`       // 讨论主题
}

type TalkInfoReq struct {
	CourseId          int `json:"course_id"`           // 课程号
	LearningContentId int `json:"learning_content_id"` // 学习内容id
	TalkId            int `json:"talk_id"`             // 讨论id
}

type TalkReq struct {
	CourseId          int `json:"course_id"`           // 课程号
	LearningContentId int `json:"learning_content_id"` // 学习内容id
	TalkId            int `json:"talk_id"`             // 讨论id
}
