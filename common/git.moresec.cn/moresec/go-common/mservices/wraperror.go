package mservices

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrEmpty    = errors.New("empty")
	ErrInvalid  = errors.New("invalid")
)

type WrapError struct {
	info string
	Err  error
}

func (w *WrapError) Error() string {
	info := ""
	if w.info != "" {
		info = w.info + ":"
	}
	if w.Err == nil {
		return info + "nil error"
	}
	return info + w.Err.Error()
}

func (w *WrapError) UnWrap() error {
	ue, ok := w.Err.(interface {
		UnWrap() error
	})
	if ok {
		return ue.UnWrap()
	}
	return w.Err
}
