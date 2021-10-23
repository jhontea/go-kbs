package models

// FeatureFlag ...
type FeatureFlag struct {
	Vendor string `json:"vendor"`
	Price  int    `json:"price"`
	Active bool   `json:"active"`
}
