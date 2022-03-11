package recommendation

import (
	"testing"

	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/usecase/recommendation/mocks"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"
	"github.com/stretchr/testify/assert"
)

func TestService_GetRecommendations(t *testing.T) {
	t.Run("No shipping options available", func(t *testing.T) {
		want := []vo.ShippingOption{}
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return([]vo.ShippingOption{}, nil)
		s := NewService(repoMock)

		got, err := s.GetRecommendations()

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})

	t.Run("Same shipping costs and estimated delivery dates", func(t *testing.T) {
		want := newFixtureShippingOptionsSameCostsAndEstimatedDays()
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return(newFixtureShippingOptionsSameCostsAndEstimatedDays(), nil)
		s := NewService(repoMock)

		got, err := s.GetRecommendations()

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})
}

func newFixtureShippingOptionsSameCostsAndEstimatedDays() []vo.ShippingOption {
	return []vo.ShippingOption{
		{
			Name:          "Option 1",
			ShippingType:  "Delivery",
			Cost:          10.0,
			EstimatedDays: 3,
		},
		{
			Name:          "Option 2",
			ShippingType:  "Custom",
			Cost:          10.0,
			EstimatedDays: 3,
		},
		{
			Name:          "Option 3",
			ShippingType:  "Pickup",
			Cost:          10.0,
			EstimatedDays: 3,
		},
	}
}
