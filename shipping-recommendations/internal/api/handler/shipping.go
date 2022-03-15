package handler

import (
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/api/presenter"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/usecase/shipping"
	"github.com/gofiber/fiber/v2"
)

// ListShippingRecommendations returns a list of recommended shipping options
func ListShippingRecommendations(service shipping.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		recommendations, err := service.GetShippingRecommendations()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
		response := presenter.NewShippingRecommendationsResponse(recommendations)
		return c.JSON(response)
	}
}
