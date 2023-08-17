package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nabil-y/zerogaspi-back/database"
	"github.com/nabil-y/zerogaspi-back/model"
)

func main() {
	app := fiber.New()

	database.InitDatabase()
	initGlobalMiddlewares(app)
	initRoutes(app)

	app.Listen(":3000")
}

func initRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/json", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello",
			"name":    "Blob",
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		db := database.DB
		user := new(model.User)

		if err := c.BodyParser(user); err != nil {
			c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
			return err
		}

		db.Create(&user)
		return c.Status(201).JSON(user)
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		db := database.DB
		var users []model.User
		db.Find(&users)
		c.JSON(users)
		return c.JSON(users)
	})
}

func initGlobalMiddlewares(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(recover.New())
}
