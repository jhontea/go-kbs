package models

import "go-kbs-soccer/exceptions"

type Team struct {
	Name string `json:"name"`
}

func (t Team) Validate() error {
	if t.Name == "" {
		return exceptions.ErrTeamNameRequired
	}

	return nil
}

type TeamPlayers struct {
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}
