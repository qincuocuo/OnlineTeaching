package merror

import "git.moresec.cn/moresec/go-common/merrors/mcode"

// Option is option for creating error.
type Option struct {
	Error error      // Wrapped error if any.
	Stack bool       // Whether recording stack information into error.
	Text  string     // Error text, which is created by New* functions.
	Code  mcode.Code // Error code if necessary.
}

// NewOption creates and returns an error with Option.
// It is the senior usage for creating error, which is often used internally in framework.
func NewOption(option Option) error {
	err := &Error{
		error: option.Error,
		text:  option.Text,
		code:  option.Code,
	}
	if option.Stack {
		err.stack = callers()
	}
	return err
}
