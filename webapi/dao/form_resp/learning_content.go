package form_resp

type LearningContentListResp struct {
	Count  int                   `json:"count"`
	Result []LearningContentItem `json:"result"`
}

type LearningContentItem struct {
	ContentId int    `json:"content_id"` // 学习内容id
	Content   string `json:"content"`    // 学习内容
	Learned   int    `json:"learned"`    // 已学习
	Unlearned int    `json:"unlearned"`  // 未学习
	Register  bool   `json:"register"`
}

type LearningResultResp struct {
	Count       int           `json:"count"`        // 人数
	StudentInfo []StudentItem `json:"student_info"` // 学生信息
}

type StudentItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
