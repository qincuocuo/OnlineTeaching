package form_resp

type GetClassListResp []int

type CourseListResp struct {
	Count  int              `json:"count"`
	Result []CourseListItem `json:"result"`
}

type CourseListItem struct {
	CourseId    int    `json:"course_id"`    // 课程ID
	CourseName  string `json:"course_name"`  // 课程名称
	Grade       int    `json:"grade"`        // 年级
	Class       int    `json:"class"`        // 班级
	TotalMember int    `json:"total_member"` // 总人数
	CreateTm    string `json:"create_tm"`    // 创建时间
}

type CourseInfoResp struct {
	Results []CourseInfoItem `json:"results"`
	Count int `json:"count"`
}

type CourseInfoItem struct {
	CourseId    int    `json:"course_id"`    // 课程ID
	CourseName  string `json:"course_name"`  // 课程名称
}