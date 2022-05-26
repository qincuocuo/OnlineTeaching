package form_req

type CreateUserReq struct {
	Role     int    `json:"role"  validate:"required"`     //用户角色 1-老师、2-学生
	Username string `json:"username"  validate:"required"` //用户名
	UserId   string `json:"user_id" validate:"required"`   //学生学号/教师工号
	Password string `json:"password"  validate:"required"` //用户密码
	Confirm  string `json:"confirm"  validate:"required"`  //确认用户密码
	Grade    int    `json:"grade"`                         //年级 1-6
	Class    int    `json:"class"`                         //班级
}

type UserPasswordReq struct {
	UserId   string `json:"user_id"`   //学生学号/教师工号
	Role     int    `json:"role"`      //用户角色
	UserName string `json:"user_name"` //用户名
	Password string `json:"password"`  //新密码
	Confirm  string `json:"confirm"`   //确认用户密码
}

type ChangePasswordReq struct {
	UserId      string `json:"user_id"`                          //学生学号/教师工号
	Password    string `json:"password" validate:"required"`     //旧账户密码
	NewPassword string `json:"new_password" validate:"required"` //新账户密码
}
