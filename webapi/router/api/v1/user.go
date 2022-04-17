package v1

import (
	"webapi/controller"
	"webapi/internal/wrapper"

	"github.com/kataras/iris/v12/core/router"
)

func RegisterUserRouter(party router.Party) {
	party.Handle("POST", "/", wrapper.Handler(controller.UserController{}.CreateUser))
	party.Handle("GET", "/", wrapper.Handler(controller.UserController{}.UserInfo))
	party.Handle("GET", "/password/", wrapper.Handler(controller.UserController{}.UserPassword))
	party.Handle("POST", "/change_password/", wrapper.Handler(controller.AuthController{}.ChangePassword))
}
