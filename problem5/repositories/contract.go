package repositories

import "go-kbs-notification/entities/models"

// FeatureFlagRepository ...
type FeatureFlagRepository interface {
	Get() ([]models.FeatureFlag, error)
	Store(featureFlag []models.FeatureFlag) (interface{}, error)
	Delete() (interface{}, error)
}
