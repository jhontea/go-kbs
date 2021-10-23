package repositories

import (
	"encoding/json"

	"go-kbs-notification/constants"
	"go-kbs-notification/entities/models"
	"go-kbs-notification/infra/redis"
)

type featureFlagRepository struct {
	db        redis.Client
	prefixKey string
}

// NewFeatureFlagRepository ...
func NewFeatureFlagRepository(db redis.Client) FeatureFlagRepository {
	return &featureFlagRepository{
		db:        db,
		prefixKey: "feature_flag",
	}
}

func (r *featureFlagRepository) Get() ([]models.FeatureFlag, error) {
	result := []models.FeatureFlag{}

	featureFlag, err := r.db.Cmd(
		constants.RedisKBS,
		redis.CmdGet, r.prefixKey,
	)
	if err != nil {
		return result, err
	}

	if featureFlag != nil {
		json.Unmarshal(featureFlag.([]byte), &result)
	}

	return result, nil
}

func (r *featureFlagRepository) Store(featureFlag []models.FeatureFlag) (interface{}, error) {
	jsonValue, _ := json.Marshal(featureFlag)

	return r.db.Cmd(
		constants.RedisKBS,
		redis.CmdSet, r.prefixKey, jsonValue,
	)
}

func (r *featureFlagRepository) Delete() (interface{}, error) {
	return r.db.Cmd(
		constants.RedisKBS,
		redis.CmdDel, r.prefixKey,
	)
}
