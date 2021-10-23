package services

import "go-kbs-notification/entities/models"

// FeatureFlagService ...
type FeatureFlagService interface {
	GetFeatureFlag() ([]models.FeatureFlag, error)
	StoreFeatureFlag(req []models.FeatureFlag) error
	DeleteFeatureFlag() (interface{}, error)
}

type SMSNotificationService interface {
	SendSMS(req models.SMS) error
}
