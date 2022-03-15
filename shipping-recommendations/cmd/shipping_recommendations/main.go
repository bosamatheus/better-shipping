package main

import (
	"log"

	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/api/router"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/infrastructure/repository"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/usecase/shipping"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func main() {
	loadEnv()
	shippingOptionsBaseURL := viper.GetString("client.shippingOptions.baseUrl")
	repo := repository.NewShippingOptionsAPIClient(shippingOptionsBaseURL)
	service := shipping.NewService(repo)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
	}))
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Shipping Recommendations API")
	})
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
	v1 := app.Group("/api/v1")
	router.ShippingRecommendationsRouter(v1, service)

	port := viper.GetString("server.port")
	log.Fatal(app.Listen(":" + port))
}

func loadEnv() {
	viper.SetConfigName("shipping_recommendations_default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error reading config file: %s", err.Error())
	}
}
