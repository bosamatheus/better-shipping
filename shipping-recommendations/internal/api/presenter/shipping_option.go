package presenter

import (
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"
)

// ShippingOptionResponse is the presentation layer for Shipping Option.
type ShippingOptionResponse struct {
	Name          string  `json:"name"`
	ShippingType  string  `json:"type"`
	Cost          float64 `json:"cost"`
	EstimatedDays int     `json:"estimated_days"`
}

// NewShippingOptionResponse creates a new Shipping Option response.
func NewShippingOptionResponse(data *vo.ShippingOption) *ShippingOptionResponse {
	return &ShippingOptionResponse{
		Name:          data.Name,
		ShippingType:  data.ShippingType,
		Cost:          data.Cost,
		EstimatedDays: data.EstimatedDays,
	}
}
