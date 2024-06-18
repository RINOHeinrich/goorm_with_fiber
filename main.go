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

	app.Get("/Products", controllers.GetProducts)
	app.Get("/Products/:id", controllers.GetProduct)
	app.Post("/Products", controllers.AddProduct)
	app.Put("/Products/:id", controllers.UpdateProduct)
	app.Delete("/Products/:id", controllers.RemoveProduct)
	log.Fatal(app.Listen(":3000"))
}
