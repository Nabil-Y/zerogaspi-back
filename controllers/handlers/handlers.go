package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nabil-y/zerogaspi-back/model"
	userRepository "github.com/nabil-y/zerogaspi-back/repository"
)

type UserController struct {
	userRepository userRepository.UserRepository
}

func NewUserController(userRepository *userRepository.UserRepository) UserController {
	return UserController{userRepository: *userRepository}
}

func (controller *UserController) CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
		return err
	}
	newUser, err := controller.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(newUser)
}

func (controller *UserController) GetUsers(c *fiber.Ctx) error {
	users, err := controller.userRepository.GetUsers()
	if err != nil {
		return err
	}
	return c.Status(200).JSON(users)
}

func (controller *UserController) GetUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")
	user, err := controller.userRepository.GetUser(userId)
	if err != nil {
		return err
	}
	if user.ID == uuid.Nil {
		response := fmt.Sprintf("No user found with id %s", userId)
		return c.Status(404).SendString(response)
	}
	return c.Status(200).JSON(user)
}

func (controller *UserController) UpdateUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
		return err
	}
	updatedUser, err := controller.userRepository.UpdateUser(userId, user)
	if err != nil {
		return err
	}
	if updatedUser.ID == uuid.Nil {
		response := fmt.Sprintf("No user found with id %s", userId)
		return c.Status(404).SendString(response)
	}
	return c.Status(200).JSON(updatedUser)
}

func (controller *UserController) DeleteUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")
	isDeleted := controller.userRepository.DeleteUser(userId)

	if isDeleted != nil {
		response := fmt.Sprintf("Can't find user with id %s", userId)
		return c.Status(404).SendString(response)
	}

	response := fmt.Sprintf("User with id %s successfuly deleted", userId)
	return c.Status(200).SendString(response)
}

func (controller *UserController) DeleteUsers(c *fiber.Ctx) error {
	// TODO: Add flag, only in test mode
	isDeleted := controller.userRepository.DeleteUsers()

	if isDeleted != nil {
		return c.Status(500).SendString("Can't delete users")
	}

	return c.Status(200).SendString("Users deleted")
}
