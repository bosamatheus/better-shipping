package shipping

import (
	"errors"
	"testing"

	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/usecase/shipping/mocks"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"
	"github.com/stretchr/testify/assert"
)

func TestService_GetShippingRecommendations(t *testing.T) {
	t.Run("NoShippingOptionsAvailable", func(t *testing.T) {
		want := []vo.ShippingOption{}
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return([]vo.ShippingOption{}, nil)
		s := NewService(repoMock)

		got, err := s.GetShippingRecommendations()

		assert.Equal(t, want, got)
		assert.NoError(t, err)
	})

	t.Run("SameShippingCostsAndSameEstimatedDeliveryDates", func(t *testing.T) {
		options, want := newFixtureSameCostsAndSameDaysShippingOptions()
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return(options, nil)
		s := NewService(repoMock)

		got, err := s.GetShippingRecommendations()

		assert.Equal(t, want, got)
		assert.NoError(t, err)
	})

	t.Run("SameShippingCostsAndDifferentEstimatedDeliveryDates", func(t *testing.T) {
		options, want := newFixtureSameCostsAndDifferentDaysShippingOptions()
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return(options, nil)
		s := NewService(repoMock)

		got, err := s.GetShippingRecommendations()

		assert.Equal(t, want, got)
		assert.NoError(t, err)
	})

	t.Run("DifferentShippingCostsAndSameEstimatedDeliveryDates", func(t *testing.T) {
		options, want := newFixtureDifferentCostsAndSameDaysShippingOptions()
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return(options, nil)
		s := NewService(repoMock)

		got, err := s.GetShippingRecommendations()

		assert.Equal(t, want, got)
		assert.NoError(t, err)
	})

	t.Run("DifferentShippingCostsAndDifferentEstimatedDeliveryDates", func(t *testing.T) {
		options, want := newFixtureDifferentCostsAndDifferentDaysShippingOptions()
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return(options, nil)
		s := NewService(repoMock)

		got, err := s.GetShippingRecommendations()

		assert.Equal(t, want, got)
		assert.NoError(t, err)
	})

	t.Run("ErrorWhileGettingShippingOptions", func(t *testing.T) {
		repoMock := new(mocks.Repository)
		repoMock.On("GetShippingOptions").Return(nil, errors.New("error"))
		s := NewService(repoMock)

		got, err := s.GetShippingRecommendations()

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
