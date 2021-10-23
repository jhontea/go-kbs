package exceptions

import "errors"

var (
	// request
	ErrTeamNameRequired   = errors.New("team name required")
	ErrPlayerNameRequired = errors.New("player name required")

	// db
	ErrTeamNotFound     = errors.New("team not found")
	ErrTeamAlreadyExist = errors.New("team already exist")

	ErrPlayerNotFound           = errors.New("player not found")
	ErrPlayerAlreadyExistInTeam = errors.New("player already exist in team")
)
