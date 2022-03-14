package recommendation

import (
	"errors"
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
		options, want := newFixtureSameCostsAndSameDaysShippingOptions()
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return(options, nil)
		s := NewService(repoMock)

		got, err := s.GetRecommendations()

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})

	t.Run("Same shipping costs and different estimated delivery dates", func(t *testing.T) {
		options, want := newFixtureSameCostsAndDifferentDaysShippingOptions()
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return(options, nil)
		s := NewService(repoMock)

		got, err := s.GetRecommendations()

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})

	t.Run("Different shipping costs and same estimated delivery dates", func(t *testing.T) {
		options, want := newFixtureDifferentCostsAndSameDaysShippingOptions()
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return(options, nil)
		s := NewService(repoMock)

		got, err := s.GetRecommendations()

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})

	t.Run("Different shipping costs and different estimated delivery dates", func(t *testing.T) {
		options, want := newFixtureDifferentCostsAndDifferentDaysShippingOptions()
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return(options, nil)
		s := NewService(repoMock)

		got, err := s.GetRecommendations()

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})

	t.Run("Error while getting shipping options", func(t *testing.T) {
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return(nil, errors.New("error"))
		s := NewService(repoMock)

		got, err := s.GetRecommendations()

		assert.Nil(t, got)
		assert.Error(t, err)
	})
}

func newFixtureSameCostsAndSameDaysShippingOptions() ([]vo.ShippingOption, []vo.ShippingOption) {
	options := []vo.ShippingOption{
		{Name: "Option 1", ShippingType: "Delivery", Cost: 10.0, EstimatedDays: 3},
		{Name: "Option 2", ShippingType: "Custom", Cost: 10.0, EstimatedDays: 3},
		{Name: "Option 3", ShippingType: "Pickup", Cost: 10.0, EstimatedDays: 3},
	}
	recommendations := []vo.ShippingOption{
		{Name: "Option 1", ShippingType: "Delivery", Cost: 10.0, EstimatedDays: 3},
		{Name: "Option 2", ShippingType: "Custom", Cost: 10.0, EstimatedDays: 3},
		{Name: "Option 3", ShippingType: "Pickup", Cost: 10.0, EstimatedDays: 3},
	}
	return options, recommendations
}

func newFixtureSameCostsAndDifferentDaysShippingOptions() ([]vo.ShippingOption, []vo.ShippingOption) {
	options := []vo.ShippingOption{
		{Name: "Option 1", ShippingType: "Delivery", Cost: 10.0, EstimatedDays: 5},
		{Name: "Option 2", ShippingType: "Custom", Cost: 10.0, EstimatedDays: 2},
		{Name: "Option 3", ShippingType: "Pickup", Cost: 10.0, EstimatedDays: 3},
	}
	recommendations := []vo.ShippingOption{
		{Name: "Option 2", ShippingType: "Custom", Cost: 10.0, EstimatedDays: 2},
		{Name: "Option 3", ShippingType: "Pickup", Cost: 10.0, EstimatedDays: 3},
		{Name: "Option 1", ShippingType: "Delivery", Cost: 10.0, EstimatedDays: 5},
	}
	return options, recommendations
}

func newFixtureDifferentCostsAndSameDaysShippingOptions() ([]vo.ShippingOption, []vo.ShippingOption) {
	options := []vo.ShippingOption{
		{Name: "Option 1", ShippingType: "Delivery", Cost: 6.0, EstimatedDays: 3},
		{Name: "Option 2", ShippingType: "Custom", Cost: 5.0, EstimatedDays: 3},
		{Name: "Option 3", ShippingType: "Pickup", Cost: 10.0, EstimatedDays: 3},
	}
	recommendations := []vo.ShippingOption{
		{Name: "Option 2", ShippingType: "Custom", Cost: 5.0, EstimatedDays: 3},
		{Name: "Option 1", ShippingType: "Delivery", Cost: 6.0, EstimatedDays: 3},
		{Name: "Option 3", ShippingType: "Pickup", Cost: 10.0, EstimatedDays: 3},
	}
	return options, recommendations
}

func newFixtureDifferentCostsAndDifferentDaysShippingOptions() ([]vo.ShippingOption, []vo.ShippingOption) {
	options := []vo.ShippingOption{
		{Name: "Option 4", ShippingType: "Delivery", Cost: 10.0, EstimatedDays: 3},
		{Name: "Option 1", ShippingType: "Delivery", Cost: 10.0, EstimatedDays: 5},
		{Name: "Option 2", ShippingType: "Custom", Cost: 5.0, EstimatedDays: 4},
		{Name: "Option 3", ShippingType: "Pickup", Cost: 7.0, EstimatedDays: 1},
	}
	recommendations := []vo.ShippingOption{
		{Name: "Option 2", ShippingType: "Custom", Cost: 5.0, EstimatedDays: 4},
		{Name: "Option 3", ShippingType: "Pickup", Cost: 7.0, EstimatedDays: 1},
		{Name: "Option 4", ShippingType: "Delivery", Cost: 10.0, EstimatedDays: 3},
		{Name: "Option 1", ShippingType: "Delivery", Cost: 10.0, EstimatedDays: 5},
	}
	return options, recommendations
}
