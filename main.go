package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"github.com/nabil-y/zerogaspi-back/database"
	"github.com/nabil-y/zerogaspi-back/model"
	userRepository "github.com/nabil-y/zerogaspi-back/repository"
)

func main() {
	app := fiber.New()

	database.InitDatabase()
	initGlobalMiddlewares(app)
	userRepository := userRepository.NewUserRepository(database.DB)

	initRoutes(app, userRepository)

	app.Listen(":3000")
}

func initRoutes(app *fiber.App, ur userRepository.UserRepository) {
	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(model.User)

		if err := c.BodyParser(user); err != nil {
			c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
			return err
		}
		newUser, err := ur.CreateUser(user)
		if err != nil {
			return err
		}
		return c.Status(201).JSON(newUser)
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		users, err := ur.GetUsers()
		if err != nil {
			return err
		}
		return c.Status(200).JSON(users)
	})

	app.Get("/users/:userId", func(c *fiber.Ctx) error {
		userIdString := c.Params("userId")
		userId, err := uuid.Parse(userIdString)
		if err != nil {
			return err
		}
		user, err := ur.GetUser(userId)
		if err != nil {
			return err
		}
		if user.ID == uuid.Nil {
			response := fmt.Sprintf("No user found with id %s", userIdString)
			return c.Status(404).SendString(response)
		}
		return c.Status(200).JSON(user)
	})

	app.Patch("/users/:userId", func(c *fiber.Ctx) error {
		userIdString := c.Params("userId")
		userId, err := uuid.Parse(userIdString)
		if err != nil {
			return err
		}
		user := new(model.User)
		if err := c.BodyParser(user); err != nil {
			c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
			return err
		}
		updatedUser, err := ur.UpdateUser(userId, user)
		if err != nil {
			return err
		}
		if updatedUser.ID == uuid.Nil {
			response := fmt.Sprintf("No user found with id %s", userIdString)
			return c.Status(404).SendString(response)
		}
		return c.Status(200).JSON(updatedUser)
	})

	app.Delete("/users/:userId", func(c *fiber.Ctx) error {
		userIdString := c.Params("userId")
		userId, err := uuid.Parse(userIdString)
		if err != nil {
			return err
		}
		isDeleted := ur.DeleteUser(userId)

		if isDeleted != nil {
			response := fmt.Sprintf("Can't find user with id %s", userIdString)
			return c.Status(404).SendString(response)
		}

		response := fmt.Sprintf("User with id %s successfuly deleted", userIdString)
		return c.Status(200).SendString(response)
	})

}

func initGlobalMiddlewares(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(recover.New())
}
