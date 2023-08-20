package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nabil-y/zerogaspi-back/controllers/handlers"
)

func InitGlobalMiddlewares(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(recover.New())
}

func InitRoutes(app *fiber.App, uc handlers.UserController) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/users", uc.CreateUser)
	v1.Get("/users", uc.GetUsers)
	v1.Get("/users/:userId", uc.GetUserById)
	v1.Patch("/users/:userId", uc.UpdateUserById)
	v1.Delete("/users/:userId", uc.DeleterUserById)
}
