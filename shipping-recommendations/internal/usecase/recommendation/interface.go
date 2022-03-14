package recommendation

import "github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"

type UseCase interface {
	GetRecommendations() ([]vo.ShippingOption, error)
}

type Repository interface {
	GetShippingOptions() ([]vo.ShippingOption, error)
}
