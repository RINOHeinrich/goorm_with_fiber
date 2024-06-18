package main

import (
	"log"

	config "github.com/RINOHeinrich/goorm_with_fiber/configs"
	"github.com/RINOHeinrich/goorm_with_fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Connect()

	app.Get("/dogs", controllers.GetDogs)
	app.Get("/dogs/:id", controllers.GetDog)
	app.Post("/dogs", controllers.AddDog)
	app.Put("/dogs/:id", controllers.UpdateDog)
	app.Delete("/dogs/:id", controllers.RemoveDog)
	log.Fatal(app.Listen(":3000"))
}
