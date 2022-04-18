package v1

import (
	"github.com/kataras/iris/v12/core/router"
	"webapi/controller"
	"webapi/internal/wrapper"
)

func RegisterRouter(party router.Party) {
	party.Handle("POST", "/", wrapper.Handler(controller.Register{}.CreateRegister))
	party.Handle("GET", "/result/", wrapper.Handler(controller.Register{}.RegisterResult))
	party.Handle("POST", "/do/", wrapper.Handler(controller.Register{}.Register))
}
