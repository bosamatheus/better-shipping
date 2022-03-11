package recommendation

import (
	"testing"

	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/infrastructure/repository"
	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"
	"github.com/stretchr/testify/assert"
)

func TestService_GetRecommendations(t *testing.T) {
	externalAPI := "/api/v1/shipping-options"
	s := &Service{
		repo: repository.NewShippingOptionsAPIClient(externalAPI),
	}

	testCases := []struct {
		name    string
		want    []vo.ShippingOption
		wantErr error
	}{
		{
			"No shipping options available", []vo.ShippingOption{}, nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := s.GetRecommendations()

			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
