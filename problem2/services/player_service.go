package services

import (
	"context"

	"go-kbs-soccer/entities/models"
	"go-kbs-soccer/repositories"
)

type playerService struct {
	repository repositories.PlayerRepository
}

func NewPlayerService(repository repositories.PlayerRepository) PlayerService {
	return &playerService{
		repository: repository,
	}
}

func (s *playerService) StorePlayer(ctx context.Context, player models.Player) (models.Player, error) {
	return s.repository.StorePlayer(ctx, player)
}

func (s *playerService) GetPlayer(ctx context.Context, name string) ([]models.Player, error) {
	return s.repository.GetPlayer(ctx, name)
}
