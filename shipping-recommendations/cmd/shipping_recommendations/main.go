package main

import (
	"io"
	"log"
	"os"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/api/router"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/infrastructure/repository"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/usecase/shipping"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/pkg/date"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/pkg/logwrapper"
	fiber "github.com/gofiber/fiber/v2"
	fiberCors "github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"

	_ "github.com/bosamatheus/better-shipping/shipping-recommendations/api"
)

const logFilePermission = 0o666

// @title        Shipping Recommendations API
// @version      0.1
// @description  A RESTful API able to recommend the same shipping options ordered by the best combination of cost and time.
// @basePath     /api
func main() {
	// load environment variables
	loadEnvVariables()

	// logger
	logger := logwrapper.NewStandardLogger(viper.GetString("env"))
	f, err := os.OpenFile(getLogFilename(), os.O_APPEND|os.O_CREATE|os.O_RDWR, logFilePermission)

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
	app := setupApp(service, logger)

	// middlewares
	app.Use(fiberCors.New())
	app.Use(fiberLogger.New(fiberLogger.Config{
		Output:     mw,
		TimeFormat: time.RFC3339,
	}))

	// start server
	port := viper.GetString("server.port")
	err = app.Listen(":" + port)

	if err != nil {
		logger.Fatalf("error starting server: %v", err)
	}
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

func setupApp(service shipping.UseCase, logger logwrapper.Logger) *fiber.App {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Shipping Recommendations API")
	})
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
	app.Get("/swagger/*", swagger.HandlerDefault)
	router.ShippingRecommendationsRouter(app, service, logger)

	return app
}
