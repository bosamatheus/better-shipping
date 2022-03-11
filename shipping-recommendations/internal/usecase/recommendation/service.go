package recommendation

import "github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetRecommendations() ([]vo.ShippingOption, error) {
	return s.getShippingOptions()
}

func (s *Service) getShippingOptions() ([]vo.ShippingOption, error) {
	return s.repo.GetShippingOptions()
}
