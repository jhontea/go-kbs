package repositories

import (
	"context"

	"go-kbs-soccer/entities/models"
)

type TeamRepository interface {
	StoreTeam(ctx context.Context, team models.Team) (models.Team, error)
	GetTeam(ctx context.Context, name string) (models.TeamPlayers, error)
}

type PlayerRepository interface {
	StorePlayer(ctx context.Context, player models.Player) (models.Player, error)
	GetPlayer(ctx context.Context, name string) ([]models.Player, error)
}
