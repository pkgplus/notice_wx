package v1

import (
	"github.com/kataras/iris"
	"github.com/xuebing1110/notice_wx/handlers"
)

func InitRouter(app *iris.Application) {

	api := app.Party("/api/v1/wxnotice")

	// user
	// api.Post("/users")
	api.Post("/login", handlers.UserLogin)
	// api.Post("/users", handlers.CreateUser)
	// api.Put("/session/:uid")

	// enery
	api.Post("/session/:sid/enery", handlers.SessionCheck, handlers.AddEnery)
	api.Get("/session/:sid/enery/count", handlers.SessionCheck, handlers.EneryCount)
	api.Get("/session/:sid/enery/one", handlers.SessionCheck, handlers.PopEnergy)

	// plugin
	api.Post("/session/:sid/plugins", handlers.SessionCheck, handlers.UserLogin)
}
