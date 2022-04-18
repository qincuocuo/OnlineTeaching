package form_resp

type RegisterResultResp struct {
	Count       int           `json:"count"`        // 人数
	StudentInfo []StudentItem `json:"student_info"` // 学生信息
}
