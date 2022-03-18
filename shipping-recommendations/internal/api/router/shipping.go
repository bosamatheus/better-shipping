package router

import (
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/api/handler"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/usecase/shipping"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/pkg/logwrapper"
	"github.com/gofiber/fiber/v2"
)

const v1BasePath = "/api/v1"

// ShippingRecommendationsRouter is the router for shipping recommendations.
func ShippingRecommendationsRouter(app *fiber.App, service shipping.UseCase, logger logwrapper.Logger) {
	v1 := app.Group(v1BasePath)
	v1.Get("/shipping-recommendations", handler.ListShippingRecommendations(service, logger))
}
