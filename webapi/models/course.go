package models

import (
	"time"
)

type Course struct {
	CourseId    int       `bson:"course_id"`    // 课程id
	ManagerId   string    `bson:"manager_id"`   // 管理员工号
	Name        string    `bson:"name"`         // 课程名称
	Grade       int       `bson:"grade"`        // 课程所属年级
	Class       int       `bson:"class"`        // 班级
	TotalMember int       `bson:"total_member"` // 学生人数
	StudentId   []string  `bson:"student_id"`   // 学生id
	CreateTm    time.Time `bson:"create_tm"`    // 创建时间
}

func (c *Course) CollectName() string {
	return "tb_course" // 课程表
}
