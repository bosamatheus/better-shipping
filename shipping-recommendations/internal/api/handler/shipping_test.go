package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	shippingMocks "github.com/bosamatheus/better-shipping/shipping-recommendations/internal/usecase/shipping/mocks"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"
	logwrapperMocks "github.com/bosamatheus/better-shipping/shipping-recommendations/pkg/logwrapper/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestListShippingRecommendations(t *testing.T) {
	path := "/api/v1/shipping-recommendations"

	t.Run("Success", func(t *testing.T) {
		want := newFixtureShippingRecommendations()
		serviceMock := new(shippingMocks.UseCase)
		serviceMock.On("GetShippingRecommendations").Return(want, nil)
		loggerMock := new(logwrapperMocks.Logger)
		app := fiber.New()
		app.Get(path, ListShippingRecommendations(serviceMock, loggerMock))

		req, _ := http.NewRequest("GET", path, nil)
		res, err := app.Test(req, -1)
		body, _ := ioutil.ReadAll(res.Body)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Contains(t, string(body), "Delivery")
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		serviceMock := new(shippingMocks.UseCase)
		serviceMock.On("GetShippingRecommendations").Return(nil, errors.New("error"))
		loggerMock := new(logwrapperMocks.Logger)
		loggerMock.On("DefaultError", errors.New("error")).Return()
		app := fiber.New()
		app.Get(path, ListShippingRecommendations(serviceMock, loggerMock))

		req, _ := http.NewRequest("GET", path, nil)
		res, err := app.Test(req, -1)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.NoError(t, err)
	})
}

func newFixtureShippingRecommendations() []vo.ShippingOption {
	return []vo.ShippingOption{
		{Name: "Option 1", ShippingType: "Custom", Cost: 5.0, EstimatedDays: 4},
		{Name: "Option 2", ShippingType: "Pickup", Cost: 7.0, EstimatedDays: 1},
		{Name: "Option 3", ShippingType: "Delivery", Cost: 10.0, EstimatedDays: 3},
	}
}
