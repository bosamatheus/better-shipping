package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShippingOptionsAPIClient_GetShippingOptions(t *testing.T) {
	repo := NewShippingOptionsAPIClient("shipping-options-api")

	got, err := repo.GetShippingOptions()

	assert.Equal(t, 4, len(got))
	assert.NoError(t, err)
}
