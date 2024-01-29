package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/zulhwk/go-restaurant/internal/controller"
	"github.com/zulhwk/go-restaurant/internal/repository"
	"github.com/zulhwk/go-restaurant/internal/usecase"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	app, _ := BuildServer()

	// Registry Restaurant.
	restaurantRepo := repository.NewRestaurantRepository()
	restaurantUsecase := usecase.NewRestaurantUsecase(restaurantRepo)
	restaurantController := controller.NewRestaurantController(restaurantUsecase)

	controller.CreateRestaurantRoutes(app, restaurantController)

	app.Listen(":3000")
}

func BuildServer() (*fiber.App, error) {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	return app, nil
}
