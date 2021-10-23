package repositories

import (
	"context"

	"go-kbs-soccer/entities/models"
	"go-kbs-soccer/exceptions"
	db "go-kbs-soccer/infra/inmemoryDB"
)

type teamRepository struct {
	db *db.InmemoryDB
}

func NewTeamRepository() TeamRepository {
	return &teamRepository{
		db: db.Inmemory(),
	}
}

func (r *teamRepository) StoreTeam(ctx context.Context, team models.Team) (models.Team, error) {
	if _, exist := r.db.IndexTeam[team.Name]; exist {
		return models.Team{}, exceptions.ErrTeamAlreadyExist
	}

	r.db.IndexTeam[team.Name] = team
	r.db.IndexTeamPlayer[team.Name] = make([]models.Player, 0)
	r.db.Teams = append(r.db.Teams, team)

	return team, nil
}

func (r *teamRepository) GetTeam(ctx context.Context, name string) (models.TeamPlayers, error) {
	team, exist := r.db.IndexTeamPlayer[name]
	if !exist {
		return models.TeamPlayers{}, exceptions.ErrTeamNotFound
	}

	return models.TeamPlayers{
		Name:    name,
		Players: team,
	}, nil
}
