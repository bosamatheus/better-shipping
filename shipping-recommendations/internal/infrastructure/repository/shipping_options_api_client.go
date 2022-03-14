package repository

import "github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"

type ShippingOptionsAPIClient struct {
	externalAPI string
}

func NewShippingOptionsAPIClient(externalAPI string) *ShippingOptionsAPIClient {
	return &ShippingOptionsAPIClient{
		externalAPI: externalAPI,
	}
}

func (r *ShippingOptionsAPIClient) GetShippingOptions() (*[]vo.ShippingOption, error) {
	return &[]vo.ShippingOption{}, nil
}
