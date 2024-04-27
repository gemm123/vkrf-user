package controller

import (
	"github.com/gemm123/vkrf-user/internal/model"
	"github.com/gemm123/vkrf-user/internal/service"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	CreateUser(ctx *fiber.Ctx) error
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService: userService}
}

func (c *userController) CreateUser(ctx *fiber.Ctx) error {
	image, err := ctx.FormFile("profile_pic")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse form",
			"status":  fiber.StatusBadRequest,
			"error":   err.Error(),
		})
	}

	var user model.UserRequest
	user.Name = ctx.FormValue("name")
	user.Email = ctx.FormValue("email")

	err = c.userService.CreateUser(user, image)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create user",
			"status":  fiber.StatusInternalServerError,
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created",
		"status":  fiber.StatusCreated,
	})
}
