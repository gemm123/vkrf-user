package controller

import (
	"github.com/gemm123/vkrf-user/internal/model"
	"github.com/gemm123/vkrf-user/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	userService service.UserService
	validate    *validator.Validate
}

type UserController interface {
	CreateUser(ctx *fiber.Ctx) error
}

func NewUserController(userService service.UserService, validate *validator.Validate) UserController {
	return &userController{userService: userService, validate: validate}
}

func (c *userController) CreateUser(ctx *fiber.Ctx) error {
	var user model.UserRequest
	user.Name = ctx.FormValue("name")
	user.Email = ctx.FormValue("email")

	err := c.validate.Struct(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user is required",
			"status":  fiber.StatusBadRequest,
			"error":   err.Error(),
		})
	}

	image, err := ctx.FormFile("profile_pic")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to parse form",
			"status":  fiber.StatusBadRequest,
			"error":   err.Error(),
		})
	}
	err = c.validate.Var(image, "required")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "profile_pic is required",
			"status":  fiber.StatusBadRequest,
			"error":   err.Error(),
		})
	}

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
