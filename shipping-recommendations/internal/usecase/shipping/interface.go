package shipping

import "github.com/bosamatheus/better-shipping/shipping-recommendations/internal/vo"

// UseCase is the interface for the shipping recommendations use case
type UseCase interface {
	GetShippingRecommendations() ([]vo.ShippingOption, error)
}

// Repository is the interface for shipping options repository
type Repository interface {
	GetShippingOptions() ([]vo.ShippingOption, error)
}
