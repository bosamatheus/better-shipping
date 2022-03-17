package router

import (
	"testing"

	shippingMocks "github.com/bosamatheus/better-shipping/shipping-recommendations/internal/usecase/shipping/mocks"
	logwrapperMocks "github.com/bosamatheus/better-shipping/shipping-recommendations/pkg/logwrapper/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestShippingRecommendationsRouter(t *testing.T) {
	app := fiber.New()
	serviceMock := new(shippingMocks.UseCase)
	loggerMock := new(logwrapperMocks.Logger)

	ShippingRecommendationsRouter(app, serviceMock, loggerMock)

	assert.Equal(t, uint32(2), app.HandlersCount())
}
