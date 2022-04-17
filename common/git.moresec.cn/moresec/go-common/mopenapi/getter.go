package mopenapi

type signCheckGetter struct {
	// default HMAC-SHA256
	signMethod string
	// default 1.0
	signVersion string
}

func (s *signCheckGetter) Key() string {
	return s.signMethod + ":" + s.signVersion
}

type SignCheckerOption func(getter *signCheckGetter)

func WithSignMethod(signMethod string) SignCheckerOption {
	return func(getter *signCheckGetter) {
		getter.signMethod = signMethod
	}
}

func WithSignVersion(signVersion string) SignCheckerOption {
	return func(getter *signCheckGetter) {
		getter.signVersion = signVersion
	}
}
