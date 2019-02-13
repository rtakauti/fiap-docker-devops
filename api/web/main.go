package main

import (
	"../repositories"
	"../services"
	"../shared"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	conf := &shared.DatabaseConf{
		DBDriver: "mysql",
		DBUser:   "root",
		DBPass:   "123456",
		DBName:   "go",
	}
	repo := repositories.NewStatusRepository(repositories.DBConn(conf))

	statusService := services.NewStatusService(repo)

	app.Get("/status", func(ctx iris.Context) { // go > 1.9
		ctx.JSON(statusService.GetAll())
	})

	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}
