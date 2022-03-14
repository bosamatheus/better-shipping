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

	t.Run("Same shipping costs and same estimated delivery dates", func(t *testing.T) {
		want := newFixtureExpectedSameCostsAndSameDaysShippingOptions()
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return(newFixtureMockSameCostsAndSameDaysShippingOptions(), nil)
		s := NewService(repoMock)

		got, err := s.GetRecommendations()

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})

	t.Run("Same shipping costs and different estimated delivery dates", func(t *testing.T) {
		want := newFixtureExpectedSameCostsAndDifferentDaysShippingOptions()
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return(newFixtureMockSameCostsAndDifferentDaysShippingOptions(), nil)
		s := NewService(repoMock)

		got, err := s.GetRecommendations()

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})
}

func newFixtureMockSameCostsAndSameDaysShippingOptions() []vo.ShippingOption {
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

func newFixtureExpectedSameCostsAndSameDaysShippingOptions() []vo.ShippingOption {
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

func newFixtureMockSameCostsAndDifferentDaysShippingOptions() []vo.ShippingOption {
	return []vo.ShippingOption{
		{
			Name:          "Option 1",
			ShippingType:  "Delivery",
			Cost:          10.0,
			EstimatedDays: 5,
		},
		{
			Name:          "Option 2",
			ShippingType:  "Custom",
			Cost:          10.0,
			EstimatedDays: 2,
		},
		{
			Name:          "Option 3",
			ShippingType:  "Pickup",
			Cost:          10.0,
			EstimatedDays: 3,
		},
	}
}

func newFixtureExpectedSameCostsAndDifferentDaysShippingOptions() []vo.ShippingOption {
	return []vo.ShippingOption{
		{
			Name:          "Option 2",
			ShippingType:  "Custom",
			Cost:          10.0,
			EstimatedDays: 2,
		},
		{
			Name:          "Option 3",
			ShippingType:  "Pickup",
			Cost:          10.0,
			EstimatedDays: 3,
		},
		{
			Name:          "Option 1",
			ShippingType:  "Delivery",
			Cost:          10.0,
			EstimatedDays: 5,
		},
	}
}
