package form_req

type GetClassListReq struct {
	Grade int `from:"grade" json:"grade" validate:"required"` //年级
}

type CreateCourseReq struct {
	CourseName string `json:"course_name" validate:"required"` // 课程名称
	Grade      int    `json:"grade" validate:"required"`       // 年级
	Class      int    `json:"class" validate:"required"`       // 班级
}

type CourseListReq struct {
	Grade               int    `form:"grade" json:"grade"`                                 // 年级
	Class               int    `form:"class" json:"class"`                                 // 班级
	CreateTm            string `form:"create_tm" json:"create_tm"`                         // 创建时间
	Search              string `form:"search" json:"search"`                               // 根据课程名搜索
	OrderingGrade       string `form:"ordering_grade" json:"ordering_grade"`               // 根据年级排序 "grade" "-grade
	OrderingTotalMember string `form:"ordering_total_member" json:"ordering_total_member"` // 根据班级总人数排序 "total_member" "-total_member"
	OrderingCreateTm    string `form:"ordering_create_tm" json:"ordering_create_tm"`       // 根据创建时间排序 "create_tm" "-create_tm"
	Page                int    `form:"page" json:"page"`                                   // 页数,用于分页
	PageSize            int    `form:"page_size" json:"page_size"`                         // 每页数量，用于分页
}

type UpdateCourseReq struct {
	CourseId   int    `json:"course_id" validate:"required"` // 课程id
	CourseName string `json:"course_name"`                   // 课程名称
	Grade      int    `json:"grade"`                         // 年级
	Class      int    `json:"class"`                         // 班级
}

type DeleteCourseReq struct {
	CourseId int `json:"course_id" validate:"required"` // 课程id
}

type EnterCourseReq struct {
	CourseId int `json:"course_id" validate:"required"` //课程id
}

type CourseInfoReq struct {
	Grade int `form:"grade" json:"grade" validate:"required"` // 年级
	Class int `form:"class" json:"class" validate:"required"` // 班级
}
