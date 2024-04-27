package main

import (
	"github.com/gemm123/vkrf-user/config"
	"github.com/gemm123/vkrf-user/internal/controller"
	grpcserver "github.com/gemm123/vkrf-user/internal/grpc"
	"github.com/gemm123/vkrf-user/internal/repository"
	"github.com/gemm123/vkrf-user/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
)

func runGrpcServer(db *pgxpool.Pool) {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	userRepository := repository.NewUserRepository(db)

	s := grpcserver.Server{
		UserRepository: userRepository,
	}

	grpcServer := grpc.NewServer()

	grpcserver.RegisterUserServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc server: %v", err)
	}
}

func runRestApiServer(db *pgxpool.Pool) {
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

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}

	db := config.InitConnPool()
	defer db.Close()

	go runGrpcServer(db)
	go runRestApiServer(db)

	select {}
}
