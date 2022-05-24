package form_req

type CreateLearningContentReq struct {
	UserId   string `json:"user_id"`   // 教师工号
	CourseId int    `json:"course_id"` // 课程id
}

type LearningContentListReq struct {
	CourseId          int    `json:"course_id"`          // 课程id
	Search            string `json:"search"`             // 根据学习内容搜索
	OrderingLearned   string `json:"ordering_learned"`   // 根据已学习人数排序 learned/-learned
	OrderingUnLearned string `json:"ordering_unlearned"` // 根据未学习人数排序 unlearned/-unlearned
	Page              int    `json:"page"`               // 分页
	PageSize          int    `json:"page_size"`          // 分页大小
}

type LearningResultReq struct {
	ContentId int    `json:"content_id"` // 学习内容id
	Status    string `json:"status"`     // 学习状态 learned-已学习/unlearned-未学习
}

type LearningReq struct {
	UserId    string `json:"user_id"`    // 学生学号
	ContentId int    `json:"content_id"` // 学习内容id
	FileName  string `json:"file_name"`  // 文件名
}
