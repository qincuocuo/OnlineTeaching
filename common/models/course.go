package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Course struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	ManagerId string `bson:"manager_id"` // 管理员工号
	CourseId int `bson:"course_id"` // 课程id
	Name string `bson:"name"` // 课程名称
	Grade int `bson:"grade"` // 课程所属年级
	Class int `bson:"class"` // 班级
	TotalMember int `bson:"total_member"` // 学生人数
	CreateTm time.Time `bson:"create_tm"` // 创建时间
}

func (c *Course) CollectName() string {
	return "tb_course" // 课程表
}