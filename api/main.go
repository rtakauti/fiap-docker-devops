package main

import (
	"os"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	conf := &DatabaseConf{
		DBDriver: "mysql",
		DBUser:   os.Getenv("DB_USER"),
		DBPass:   os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	repo := NewContactRepository(DBConn(conf))

	contactService := NewContactService(repo)

	app.Get("/contacts", func(ctx iris.Context) {
		ctx.JSON(contactService.GetAll())
	})

	app.Get("/contacts/{id:int}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		contact := contactService.GetById(id)
		if &contact != nil {
			ctx.JSON(contact)
		}
		ctx.StatusCode(iris.StatusNotFound)
	})

	app.Delete("/contacts/{id:int}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		exists := contactService.Delete(id)
		if !exists {
			ctx.StatusCode(iris.StatusNotFound)
		} else {
			ctx.StatusCode(iris.StatusNoContent)
		}
		return
	})

	app.Post("/contacts", func(ctx iris.Context) {
		var contact Contact
		ctx.ReadJSON(&contact)

		inserterContact, err := contactService.Insert(contact)
		if err == nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		}
		ctx.JSON(inserterContact)
	})

	app.Put("/contacts/{id:int}", func(ctx iris.Context) {
		var contact Contact
		ctx.ReadJSON(&contact)
		id, _ := ctx.Params().GetInt("id")

		updatedContact, err := contactService.Update(id,contact)
		if err == nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		}

		if &updatedContact == nil {
			ctx.StatusCode(iris.StatusNotFound)
			return
		}

		ctx.JSON(updatedContact)
	})

	app.Run(
		// Start the web server at localhost:8080
		iris.Addr(":80"),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}
