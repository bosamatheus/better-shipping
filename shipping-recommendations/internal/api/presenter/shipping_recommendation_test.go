package presenter

import (
	"encoding/json"
	"testing"

	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"
	"github.com/stretchr/testify/assert"
)

func TestNewShippingRecommendationsResponse(t *testing.T) {
	got := NewShippingRecommendationsResponse(newFixtureShippingOptions())
	firstOne := got.ShippingRecommendations[0]

	assert.Equal(t, 1, len(got.ShippingRecommendations))
	assert.Equal(t, "Option 1", firstOne.Name)
	assert.Equal(t, "Delivery", firstOne.ShippingType)
	assert.Equal(t, 10.0, firstOne.Cost)
	assert.Equal(t, 3, firstOne.EstimatedDays)
}

func TestShippingRecommendationsResponseMarshal(t *testing.T) {
	want := `{"shipping_recommendations":[{"name":"Option 1","type":"Delivery","cost":10,"estimated_days":3}]}`

	res := NewShippingRecommendationsResponse(newFixtureShippingOptions())
	got, err := json.Marshal(res)

	assert.NoError(t, err)
	assert.Equal(t, want, string(got))
}

func newFixtureShippingOptions() []vo.ShippingOption {
	return []vo.ShippingOption{
		*newFixtureShippingOption(),
	}
}
