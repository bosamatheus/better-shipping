package repository

import "github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"

type ShippingOptionsAPIClient struct {
	baseURL string
}

func NewShippingOptionsAPIClient(baseURL string) *ShippingOptionsAPIClient {
	return &ShippingOptionsAPIClient{
		baseURL: baseURL,
	}
}

func (r *ShippingOptionsAPIClient) GetShippingOptions() ([]vo.ShippingOption, error) {
	//nolint:gomnd // mock implementation
	options := []vo.ShippingOption{
		{Name: "Option 4", ShippingType: "Delivery", Cost: 10.0, EstimatedDays: 3},
		{Name: "Option 1", ShippingType: "Delivery", Cost: 10.0, EstimatedDays: 5},
		{Name: "Option 2", ShippingType: "Custom", Cost: 5.0, EstimatedDays: 4},
		{Name: "Option 3", ShippingType: "Pickup", Cost: 7.0, EstimatedDays: 1},
	}

	return options, nil
}
