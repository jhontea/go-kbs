package repositories

import (
	"context"

	"go-kbs-soccer/entities/models"
	"go-kbs-soccer/exceptions"
	db "go-kbs-soccer/infra/inmemoryDB"
)

type playerRepository struct {
	db *db.InmemoryDB
}

func NewPlayerRepository() PlayerRepository {
	return &playerRepository{
		db: db.Inmemory(),
	}
}

// StorePlayer store player to the team, same name just for different team
func (r *playerRepository) StorePlayer(ctx context.Context, player models.Player) (models.Player, error) {
	if _, exist := r.db.IndexTeam[player.TeamName]; !exist {
		return models.Player{}, exceptions.ErrTeamNotFound
	}

	r.db.IndexTeamPlayer[player.TeamName] = append(r.db.IndexTeamPlayer[player.TeamName], player)
	players, exist := r.db.IndexPlayer[player.Name]
	if exist {
		for _, v := range players {
			if v.TeamName == player.TeamName {
				return models.Player{}, exceptions.ErrPlayerAlreadyExistInTeam
			}
		}
	}

	r.db.IndexPlayer[player.Name] = append(r.db.IndexPlayer[player.Name], player)
	r.db.Players = append(r.db.Players, player)

	return player, nil
}

func (r *playerRepository) GetPlayer(ctx context.Context, name string) ([]models.Player, error) {
	players, exist := r.db.IndexPlayer[name]
	if !exist {
		return []models.Player{}, exceptions.ErrPlayerNotFound
	}

	return players, nil
}
