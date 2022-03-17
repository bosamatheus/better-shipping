package presenter

import (
	"encoding/json"
	"testing"

	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"
	"github.com/stretchr/testify/assert"
)

func TestNewShippingRecommendationsResponse(t *testing.T) {
	got := NewShippingRecommendationsResponse(newFixtureShippingOptions())
	first := got.ShippingRecommendations[0]

	assert.Equal(t, "Option 1", first.Name)
	assert.Equal(t, "Delivery", first.ShippingType)
	assert.Equal(t, 10.0, first.Cost)
	assert.Equal(t, 3, first.EstimatedDays)
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
