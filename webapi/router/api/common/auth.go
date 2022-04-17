package common

import (
	"webapi/controller"
	"webapi/internal/wrapper"

	"github.com/kataras/iris/v12/core/router"
)

func RegisterAuthRouter(party router.Party) {
	party.Handle("GET", "/auth/verifycode/", wrapper.Handler(controller.AuthController{}.VerifyCode))
	party.Handle("POST", "/auth/login/", wrapper.Handler(controller.AuthController{}.Login))
	party.Handle("POST", "/auth/logout/", wrapper.Handler(controller.AuthController{}.Logout))
}
