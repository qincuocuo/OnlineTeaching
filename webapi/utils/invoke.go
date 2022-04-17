package utils

import "runtime"

type invoke struct{}

var Invoke invoke

func (invoke) GetFuncName(level int) (funcName string) {
	pc, _, _, _ := runtime.Caller(level)
	return runtime.FuncForPC(pc).Name()
}
