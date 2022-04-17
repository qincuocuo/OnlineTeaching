package cache

import (
	"git.moresec.cn/moresec/go-common/mlog"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var captchaDigitOpt = base64Captcha.ConfigDigit{
	Height:     80,
	Width:      240,
	MaxSkew:    0.7,
	DotCount:   80,
	CaptchaLen: 4,
}

var captchaCharacterOpt = base64Captcha.ConfigCharacter{
	Height:             80,
	Width:              240,
	Mode:               3,
	ComplexOfNoiseText: 0,
	ComplexOfNoiseDot:  0,
	IsUseSimpleFont:    false,
	IsShowHollowLine:   false,
	IsShowNoiseDot:     false,
	IsShowNoiseText:    false,
	IsShowSlimeLine:    false,
	IsShowSineLine:     false,
	CaptchaLen:         4,
}

// GenDigitCaptcha 生成验证码模块.
func GenDigitCaptcha() (id, pngData string) {
	captchaId, characterCap := base64Captcha.GenerateCaptcha("", captchaCharacterOpt)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(characterCap)
	return captchaId, base64Png
}

// VerifyCaptcha 校验验证码是否正确
func VerifyCaptcha(id, value string) bool {
	return base64Captcha.VerifyCaptcha(id, value)
}

// 设置验证码存储方式
type customizeRdsStore struct {
}

var defaultCaptchaStore = new(customizeRdsStore)

// Set 验证码过期时间 - 一分钟
func (s *customizeRdsStore) Set(id string, value string) {
	_, err := RedisCli.SetEx(RdxCaptchaKey(id), value, 60)
	if err != nil {
		mlog.Error("set captcha error", zap.Error(err), zap.String("id", id))
	}
}

func (s *customizeRdsStore) Get(id string, clear bool) (value string) {
	val, err := RedisCli.Get(RdxCaptchaKey(id))
	if err != nil {
		mlog.Error("get captcha error", zap.Error(err), zap.String("id", id))
		return ""
	}
	if clear {
		_ = RedisCli.Del(RdxCaptchaKey(id))
	}
	return val
}
