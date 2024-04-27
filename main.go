package main

import (
	"github.com/gemm123/vkrf-user/config"
	"github.com/gemm123/vkrf-user/internal/controller"
	"github.com/gemm123/vkrf-user/internal/repository"
	"github.com/gemm123/vkrf-user/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}

	db := config.InitConnPool()
	defer db.Close()

	userRepository := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepository)

	userController := controller.NewUserController(userService)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")

	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.Post("/create", userController.CreateUser)

	app.Listen(":3000")
}
