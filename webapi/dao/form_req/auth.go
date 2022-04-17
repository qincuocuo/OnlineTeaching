package form_req

// AuthLoginReq 用户登录
type AuthLoginReq struct {
	CaptId   string `json:"captid" validate:"required"`   //验证码唯一ID
	UserId   string `json:"user_id" validate:"required"`  //用户Id
	Password string `json:"password" validate:"required"` //用户密码
	Vcode    string `json:"vcode" validate:"required"`    //验证码
}
