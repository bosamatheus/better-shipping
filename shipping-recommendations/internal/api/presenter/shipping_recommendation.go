package presenter

import (
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"
)

type ShippingRecommendationsResponse struct {
	ShippingRecommendations []ShippingOptionResponse `json:"shipping_recommendations"`
}

func NewShippingRecommendationsResponse(data []vo.ShippingOption) *ShippingRecommendationsResponse {
	recommendations := make([]ShippingOptionResponse, len(data))
	for i, v := range data {
		recommendations[i] = *NewShippingOptionResponse(&v)
	}
	return &ShippingRecommendationsResponse{
		ShippingRecommendations: recommendations,
	}
}
