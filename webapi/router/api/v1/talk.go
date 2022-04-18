package v1

import (
	"github.com/kataras/iris/v12/core/router"
	"webapi/controller"
	"webapi/internal/wrapper"
)

func RegisterTalkRouter(party router.Party) {
	party.Handle("POST", "/", wrapper.Handler(controller.Talk{}.CreateTalk))
	party.Handle("GET", "/", wrapper.Handler(controller.Talk{}.TalkInfo))
	party.Handle("POST", "/do/", wrapper.Handler(controller.Talk{}.Talk))
}
