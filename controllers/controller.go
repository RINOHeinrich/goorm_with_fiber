package controllers

import (
	config "github.com/RINOHeinrich/goorm_with_fiber/configs"
	"github.com/RINOHeinrich/goorm_with_fiber/models"
	"github.com/gofiber/fiber/v2"
)

func GetDogs(c *fiber.Ctx) error {
	var dogs []models.Dog

	config.Database.Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	id := c.Params("id")
	var dog models.Dog

	result := config.Database.Find(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	dog := new(models.Dog)

	if err := c.BodyParser(dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	dog := new(models.Dog)
	id := c.Params("id")

	if err := c.BodyParser(dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	id := c.Params("id")
	var dog models.Dog

	result := config.Database.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
