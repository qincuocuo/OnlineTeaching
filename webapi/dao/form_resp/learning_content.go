package form_resp

type LearningContentListResp struct {
	Count  int                   `json:"count"`
	Result []LearningContentItem `json:"result"`
}

type LearningContentItem struct {
	Content   string `json:"content"`   // 学习内容
	Learned   int    `json:"learned"`   // 已学习
	Unlearned int    `json:"unlearned"` // 未学习
	FilePath  string `json:"file_path"` // 文件路径
}

type LearningResultResp struct {
	Count       int           `json:"count"`        // 人数
	StudentInfo []StudentItem `json:"student_info"` // 学生信息
}

type StudentItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
