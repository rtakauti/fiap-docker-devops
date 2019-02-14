package api

import (
	"api/datamodels"
	"api/repositories"
	"api/services"
	"api/shared"
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
	repo := repositories.NewContactRepository(repositories.DBConn(conf))

	contactService := services.NewContactService(repo)

	app.Get("/contatos", func(ctx iris.Context) {
		ctx.JSON(contactService.GetAll())
	})

	app.Get("/contatos/{id:int}", func(ctx iris.Context) {
		contact := contactService.Get(ctx.Params().GetInt("id"))
		if contact != nil {
			ctx.JSON(contact)
		}
		ctx.StatusCode(iris.StatusNotFound)
	})

	app.Delete("/contatos/{id:int}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		exists := contactService.Delete(id)
		if !exists {
			ctx.StatusCode(iris.StatusNotFound)
		} else {
			ctx.StatusCode(iris.StatusNoContent)
		}
		return
	})

	app.Post("/contatos", func(ctx iris.Context) {
		var contact datamodels.Contact
		ctx.ReadJSON(&contact)

		inserterContact, err := contactService.Insert(contact)
		if err == nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		}
		ctx.JSON(inserterContact)
	})

	app.Put("/contatos/{id:int}", func(ctx iris.Context) {
		var contact datamodels.Contact
		ctx.ReadJSON(&contact)
		id, _ := ctx.Params().GetInt("id")

		updatedContact, err := contactService.Update(id,contact)
		if err == nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		}

		if updatedContact == nil {
			ctx.StatusCode(iris.StatusNotFound)
			return
		}

		ctx.JSON(updatedContact)
	})

	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("localhost:80"),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}
