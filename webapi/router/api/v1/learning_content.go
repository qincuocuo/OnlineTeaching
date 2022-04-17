package v1

import (
	"github.com/kataras/iris/v12/core/router"
	"webapi/controller"
	"webapi/internal/wrapper"
)

func RegisterLearningContentRouter(party router.Party) {
	party.Handle("POST", "/", wrapper.Handler(controller.LearningController{}.CreateLearningContent)) //添加学习内容
}


