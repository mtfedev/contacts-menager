package main

import "github.com/gofiber/fiber/v2"

var (
	userHandler = api.NewUserHandler(userStore)
	User        = userStore
	app         = fiber.New(config)
	apiv1       = app.Group("/api/v1")
)

func main() {
	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

}
