package v1

import (
	"github.com/kataras/iris/v12/core/router"
	"webapi/controller"
	"webapi/internal/wrapper"
)

func RegisterExercisesRouter(party router.Party) {
	party.Handle("POST", "/", wrapper.Handler(controller.Exercises{}.CreateExercises))
	party.Handle("POST", "/do/", wrapper.Handler(controller.Exercises{}.Exercises))
}
