package vo

// ShippingOption is a value object for shipping option
type ShippingOption struct {
	Name          string
	ShippingType  string
	Cost          float64
	EstimatedDays int
}
