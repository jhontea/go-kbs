package services

import (
	"context"

	"go-kbs-soccer/entities/models"
	"go-kbs-soccer/repositories"
)

type teamService struct {
	repository repositories.TeamRepository
}

func NewTeamService(repository repositories.TeamRepository) TeamService {
	return &teamService{
		repository: repository,
	}
}

func (s *teamService) StoreTeam(ctx context.Context, team models.Team) (models.Team, error) {
	return s.repository.StoreTeam(ctx, team)
}

func (s *teamService) GetTeam(ctx context.Context, name string) (models.TeamPlayers, error) {
	return s.repository.GetTeam(ctx, name)
}
