package router

import (
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/api/handler"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/usecase/shipping"
	"github.com/gofiber/fiber/v2"
)

// ShippingRecommendationsRouter is the router for shipping recommendations
func ShippingRecommendationsRouter(r fiber.Router, service shipping.UseCase) {
	r.Get("/shipping-recommendations", handler.ListShippingRecommendations(service))
}
