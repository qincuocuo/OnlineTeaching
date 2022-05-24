package form_req

type CreateCourseReq struct {
	CourseName string `json:"course_name"` // 课程名称
	Grade      int    `json:"grade"`       // 年级
	Class      int    `json:"class"`       // 班级
}

type CourseListReq struct {
	Grade               int    `json:"grade"`                 // 年级
	Class               int    `json:"class"`                 // 班级
	CreateTm            string `json:"create_tm"`             // 创建时间
	Search              string `json:"search"`                // 根据课程名搜索
	OrderingGrade       string `json:"ordering_grade"`        // 根据年级排序 "grade" "-grade
	OrderingTotalMember string `json:"ordering_total_member"` // 根据班级总人数排序 "total_member" "-total_member"
	OrderingCreateTm    string `json:"ordering_create_tm"`    // 根据创建时间排序 "create_tm" "-create_tm"
	Page                int    `json:"page"`                  // 页数,用于分页
	PageSize            int    `json:"page_size"`             // 每页数量，用于分页
}

type UpdateCourseReq struct {
	CourseId   int    `json:"course_id"`   // 课程id
	CourseName string `json:"course_name"` // 课程名称
	Grade      int    `json:"grade"`       // 年级
	Class      int    `json:"class"`       // 班级
}

type DeleteCourseReq struct {
	CourseId int    `json:"course_id"` // 课程id
}

type EnterCourseReq struct {
	CourseId int    `json:"course_id"` //课程id
}
