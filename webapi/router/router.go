package router

import (
	"webapi/router/api/common"
	v1 "webapi/router/api/v1"

	"github.com/kataras/iris/v12"
)

func InitRouters(app *iris.Application) {
	commonRouter := app.Party("/")
	{
		// 基础认证与登录接口
		common.RegisterAuthRouter(commonRouter)
	}
	appRouter := app.Party("/v1/")
	{
		// 用户接口
		appUserRouter := appRouter.Party("/user")
		{
			v1.RegisterUserRouter(appUserRouter)
		}

		// 课程接口
		appClassRouter := appRouter.Party("/course")
		{
			v1.RegisterCourseRouter(appClassRouter)
		}

		// 学习内容接口
		appLearningContentRouter := appRouter.Party("/learning_content")
		{
			v1.RegisterLearningContentRouter(appLearningContentRouter)
		}

		// 签到互动接口
		appActiveRouter := appRouter.Party("/active")
		{
			v1.RegisterActiveRouter(appActiveRouter)
		}

		// 讨论聊天接口
		appTalkRouter := appRouter.Party("/talk")
		{
			v1.RegisterTalkRouter(appTalkRouter)
		}

		// 习题接口
		appExercisesRouter := appRouter.Party("/exercises")
		{
			v1.RegisterExercisesRouter(appExercisesRouter)
		}
	}
}
