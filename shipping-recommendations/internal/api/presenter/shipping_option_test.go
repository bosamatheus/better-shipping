package presenter

import (
	"encoding/json"
	"testing"

	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"
	"github.com/stretchr/testify/assert"
)

func TestNewShippingOptionResponse(t *testing.T) {
	got := NewShippingOptionResponse(newFixtureShippingOption())

	assert.Equal(t, "Option 1", got.Name)
	assert.Equal(t, "Delivery", got.ShippingType)
	assert.Equal(t, 10.0, got.Cost)
	assert.Equal(t, 3, got.EstimatedDays)
}

func TestShippingOptionsResponseMarshal(t *testing.T) {
	want := `{"name":"Option 1","type":"Delivery","cost":10,"estimated_days":3}`

	res := NewShippingOptionResponse(newFixtureShippingOption())
	got, err := json.Marshal(res)

	assert.NoError(t, err)
	assert.Equal(t, want, string(got))
}

func newFixtureShippingOption() *vo.ShippingOption {
	return &vo.ShippingOption{
		Name:          "Option 1",
		ShippingType:  "Delivery",
		Cost:          10.0,
		EstimatedDays: 3,
	}
}
