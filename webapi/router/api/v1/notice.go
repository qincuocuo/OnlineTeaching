package v1

import (
	"github.com/kataras/iris/v12/core/router"
	"webapi/controller"
	"webapi/internal/wrapper"
)

func RegisterNoticeRouter(party router.Party) {
	party.Handle("GET", "/", wrapper.Handler(controller.Notice{}.Notice))
}
