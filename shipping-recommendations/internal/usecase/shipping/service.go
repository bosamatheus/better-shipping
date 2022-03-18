package shipping

import (
	"sort"

	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"
)

// Service is the use case for Shipping Recommendations.
type Service struct {
	repo Repository
}

// NewService returns a new Shipping Recommendations use case.
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetShippingRecommendations returns Shipping Recommendations.
func (s *Service) GetShippingRecommendations() ([]vo.ShippingOption, error) {
	options, err := s.getShippingOptions()
	if err != nil {
		return nil, err
	}

	recommendations := s.sortShippingOptions(options)

	return recommendations, nil
}

func (s *Service) getShippingOptions() ([]vo.ShippingOption, error) {
	return s.repo.GetShippingOptions()
}

func (s *Service) sortShippingOptions(options []vo.ShippingOption) []vo.ShippingOption {
	recommendations := make([]vo.ShippingOption, len(options))
	copy(recommendations, options)

	sort.SliceStable(recommendations, func(i, j int) bool {
		sameCosts := recommendations[i].Cost == recommendations[j].Cost
		lowerCosts := recommendations[i].Cost < recommendations[j].Cost
		lowerEstimatedDays := recommendations[i].EstimatedDays < recommendations[j].EstimatedDays
		return sameCosts && lowerEstimatedDays || lowerCosts
	})

	return recommendations
}
