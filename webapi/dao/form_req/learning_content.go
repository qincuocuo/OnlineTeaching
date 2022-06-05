package form_req

type CreateLearningContentReq struct {
	CourseId int    `form:"course_id" json:"course_id" validate:"required"` // 课程id
	Title    string `form:"title" json:"title"`                             // 课程标题
}

type LearningContentListReq struct {
	CourseId          int    `form:"course_id" json:"course_id" validate:"required"` // 课程id
	Search            string `form:"search" json:"search"`                           // 根据学习内容搜索
	OrderingLearned   string `form:"ordering_learned" json:"ordering_learned"`       // 根据已学习人数排序 learned/-learned
	OrderingUnLearned string `form:"ordering_unlearned" json:"ordering_un_learned"`  // 根据未学习人数排序 unlearned/-unlearned
	Page              int    `form:"page" json:"page"`                               // 分页
	PageSize          int    `form:"page_size" json:"page_size"`                     // 分页大小
}

type LearningResultReq struct {
	ContentId int    `form:"content_id" json:"content_id" validate:"required"` // 学习内容id
	Status    string `form:"status" json:"status" validate:"required"`         // 学习状态 learned-已学习/unlearned-未学习
}

type LearningReq struct {
	ContentId int `json:"content_id" form:"content_id" validate:"required"` // 学习内容id
}

type StartChatReq struct {
	ContentId int `json:"content_id" validate:"required"` // 学习内容id
}
