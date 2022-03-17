package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/api/router"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/infrastructure/repository"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/usecase/shipping"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/pkg/date"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/pkg/logwrapper"
	"github.com/gofiber/fiber/v2"
	fiberCors "github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func main() {
	// load environment variables
	loadEnvVariables()

	// logger
	logger := logwrapper.NewStandardLogger(viper.GetString("env"))
	f, err := os.OpenFile(
		getLogFilename(),
		os.O_APPEND|os.O_CREATE|os.O_RDWR,
		0666,
	)
	if err != nil {
		logger.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	mw := io.MultiWriter(os.Stdout, f)
	logger.SetOutput(mw)

	// repositories
	shippingOptionsBaseURL := viper.GetString("client.shippingOptions.baseUrl")
	repo := repository.NewShippingOptionsAPIClient(shippingOptionsBaseURL)

	// use cases
	service := shipping.NewService(repo)

	// routers
	app := fiber.New()
	app.Use(fiberCors.New())
	app.Use(fiberLogger.New(fiberLogger.Config{
		Output:     mw,
		TimeFormat: time.RFC3339,
	}))
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Shipping Recommendations API")
	})
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
	v1 := app.Group("/api/v1")
	router.ShippingRecommendationsRouter(v1, service, logger)

	// start server
	port := viper.GetString("server.port")
	log.Fatal(app.Listen(":" + port))
}

func loadEnvVariables() {
	viper.SetConfigName("shipping_recommendations_development")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error while reading config file: %s", err)
	}
}

func getLogFilename() string {
	return "logs/shipping_recommendations_" + date.FormatDate(time.Now()) + ".log"
}
