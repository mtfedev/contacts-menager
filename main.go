package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	api "github.com/mtfedev/contacts-menager/Api"
	"github.com/mtfedev/contacts-menager/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	var (
		userStore = db.NewMongoUser(client)
		_         = &db.Store{
			User: userStore}

		userHandler = api.NewUserHandler(userStore)

		app   = fiber.New(config)
		apiv1 = app.Group("/api/v1")
	)

	//User

	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

}
