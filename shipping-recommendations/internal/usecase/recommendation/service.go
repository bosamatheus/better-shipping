package recommendation

import (
	"sort"

	"github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetRecommendations() ([]vo.ShippingOption, error) {
	options, err := s.getShippingOptions()
	if err != nil {
		return nil, err
	}

	sort.SliceStable(options, func(i, j int) bool {
		return options[i].EstimatedDays < options[j].EstimatedDays
	})
	return options, nil
}

func (s *Service) getShippingOptions() ([]vo.ShippingOption, error) {
	return s.repo.GetShippingOptions()
}
