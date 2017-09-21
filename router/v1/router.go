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
	// api.Put("/user/:sid")

	// enery
	api.Post("/session/:sid/energy", handlers.SessionCheck, handlers.AddEnery)
	api.Get("/session/:sid/energy/count", handlers.SessionCheck, handlers.EneryCount)
	api.Get("/session/:sid/energy/one", handlers.SessionCheck, handlers.PopEnergy)

	// plugin
	api.Get("/session/:sid/plugins", handlers.SessionCheck, handlers.ListUserPlugins)
	api.Post("/session/:sid/plugins", handlers.SessionCheck, handlers.AddUserPlugin)
	api.Delete("/session/:sid/plugins/:pluginid", handlers.SessionCheck, handlers.DeleteUserPlugin)
	api.Get("/session/:sid/plugins/:pluginid", handlers.SessionCheck, handlers.GetUserPlugin)
}
