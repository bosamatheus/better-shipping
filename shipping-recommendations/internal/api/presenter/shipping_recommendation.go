package presenter

import (
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"
)

// ShippingRecommendationsResponse is the presentation layer for Shipping Recommendations.
type ShippingRecommendationsResponse struct {
	ShippingRecommendations []ShippingOptionResponse `json:"shipping_recommendations"`
}

// NewShippingRecommendationsResponse creates a new Shipping Recommendations response.
func NewShippingRecommendationsResponse(data []vo.ShippingOption) *ShippingRecommendationsResponse {
	recommendations := make([]ShippingOptionResponse, len(data))
	for i := range data {
		recommendations[i] = *NewShippingOptionResponse(&data[i])
	}

	return &ShippingRecommendationsResponse{
		ShippingRecommendations: recommendations,
	}
}
