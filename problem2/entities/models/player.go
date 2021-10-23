package models

import "go-kbs-soccer/exceptions"

type Player struct {
	TeamName string `json:"team_name"`
	Name     string `json:"name"`
}

func (p Player) Validate() error {
	if p.Name == "" {
		return exceptions.ErrPlayerNameRequired
	}

	if p.TeamName == "" {
		return exceptions.ErrTeamNameRequired
	}

	return nil
}
