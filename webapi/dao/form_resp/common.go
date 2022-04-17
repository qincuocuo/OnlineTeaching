package form_resp

// StatusResp Common
type StatusResp struct {
	Status string `json:"status"` //基本状态(string)
}

// UserInfoResp 获取用户信息
type UserInfoResp struct {
	UserId        string `json:"user_id"`         //学生学号/教师工号
	Role          int    `json:"role"`            //用户角色
	UserName      string `json:"user_name"`       //用户名
	Grade         int    `json:"grade"`           //年级
	Class         int    `json:"class"`           //班级
	LoginTime     string `json:"login_time"`      //注册时间
	LastLoginTime string `json:"last_login_time"` //最近登陆时间
}

type UserPasswordResp struct {
	Password string `json:"password"` //用户密码
}
