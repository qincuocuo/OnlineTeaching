package form_resp

// AuthVerifyCodeResp 获取验证码响应结构体
type AuthVerifyCodeResp struct {
	CaptId string `json:"captid"` //验证码唯一ID
	Image  string `json:"image"`  //验证码图片数据
}

// AuthLoginResp 用户登录响应结构体
type AuthLoginResp struct {
	Uid           int    `json:"uid"`           //用户ID
	Role          int    `json:"role"`          //用户权限
	Enable        bool   `json:"enable"`        //是否允许登录 (true)-允许登录、(false)-拒绝登录
	Authorization string `json:"authorization"` //用户Token信息
}
