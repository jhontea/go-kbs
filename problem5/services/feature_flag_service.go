package services

import (
	"go-kbs-notification/entities/models"
	"go-kbs-notification/repositories"
)

type featureFlagService struct {
	repository repositories.FeatureFlagRepository
}

// NewFeatureFlagService ...
func NewFeatureFlagService(repository repositories.FeatureFlagRepository) FeatureFlagService {
	return &featureFlagService{
		repository: repository,
	}
}

func (s *featureFlagService) GetFeatureFlag() ([]models.FeatureFlag, error) {
	return s.repository.Get()
}

func (s *featureFlagService) StoreFeatureFlag(req []models.FeatureFlag) error {
	_, err := s.repository.Store(req)
	if err != nil {
		return err
	}

	return nil
}

func (s *featureFlagService) DeleteFeatureFlag() (interface{}, error) {
	return s.repository.Delete()
}
