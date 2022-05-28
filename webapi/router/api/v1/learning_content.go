package v1

import (
	"github.com/kataras/iris/v12/core/router"
	"webapi/controller"
	"webapi/internal/wrapper"
)

func RegisterLearningContentRouter(party router.Party) {
	party.Handle("POST", "/", wrapper.Handler(controller.LearningController{}.CreateLearningContent))  //添加学习内容
	party.Handle("GET", "/", wrapper.Handler(controller.LearningController{}.LearningContentList))     //获取学习内容列表
	party.Handle("GET", "/result/", wrapper.Handler(controller.LearningController{}.LearningResult))   //获取学习结果
	party.Handle("GET", "/learning/", wrapper.Handler(controller.LearningController{}.Learning))       //进入学习
	party.Handle("GET", "/learning/chat/", wrapper.Handler(controller.LearningController{}.StartChat)) //进入学习
}
