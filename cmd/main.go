package main

import (
	"os"

	"github.com/kataras/iris"
	"github.com/xuebing1110/notice_wx/crontask"
	"github.com/xuebing1110/notice_wx/handlers"
	"github.com/xuebing1110/notice_wx/router/v1"
	"github.com/xuebing1110/noticeplat/storage/redis"
)

func main() {
	// task schedule
	crontask.Schedual(redis.Client)

	// http server
	app := iris.New()
	app.Use(handlers.SessionStorage)

	v1.InitRouter(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Run(iris.Addr(":" + port))
}
