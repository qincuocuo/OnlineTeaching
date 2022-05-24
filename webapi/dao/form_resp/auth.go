package form_resp

// AuthVerifyCodeResp 获取验证码响应结构体
type AuthVerifyCodeResp struct {
	CaptId string `json:"captid"` //验证码唯一ID
	Image  string `json:"image"`  //验证码图片数据
}

// AuthLoginResp 用户登录响应结构体
type AuthLoginResp struct {
	UserId           string    `json:"user_id"`           //用户ID
	Role          int    `json:"role"`          //用户权限
	Authorization string `json:"authorization"` //用户Token信息
}
