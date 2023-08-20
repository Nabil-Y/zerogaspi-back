package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nabil-y/zerogaspi-back/controllers/handlers"
	"github.com/nabil-y/zerogaspi-back/controllers/routes"
	"github.com/nabil-y/zerogaspi-back/database"
	userRepositoryImpl "github.com/nabil-y/zerogaspi-back/repository/impl"
)

func main() {
	app := fiber.New()

	database.InitDatabase()
	routes.InitGlobalMiddlewares(app)

	userRepository := userRepositoryImpl.NewUserRepository(database.DB)

	userController := handlers.NewUserController(&userRepository)

	routes.InitRoutes(app, userController)

	app.Listen(":3000")
}
