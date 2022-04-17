package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// User 用户表
type User struct {
	ID              bson.ObjectId `bson:"_id,omitempty"`
	Role            int           `bson:"role"`               //用户角色 1-老师、2-学生
	UserId          string        `bson:"user_id"`            //学生学号/老师工号
	UserName        string        `bson:"username"`           //用户名
	Password        string        `bson:"password"`           //用户密码
	Grade           int           `bson:"grade"`              //年级
	Class           int           `bson:"class"`              //班级
	LastPwdChangeTm time.Time     `bson:"last_pwd_change_tm"` //最近一次修改密码时间
	LastLoginTm     time.Time     `bson:"last_login_tm"`      //最近登录时间
	InsertTm        time.Time     `bson:"insert_tm"`          //入库时间
	UpdateTm        time.Time     `bson:"update_tm"`          //更新时间
}

func (u *User) CollectName() string {
	return "tb_user" // 用户表
}
