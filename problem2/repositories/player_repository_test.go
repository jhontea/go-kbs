package repositories

import (
	"context"
	"errors"
	"fmt"
	"go-kbs-soccer/entities/models"
	"go-kbs-soccer/exceptions"
	db "go-kbs-soccer/infra/inmemoryDB"
	"reflect"
	"testing"
)

func TestNewPlayerRepository(t *testing.T) {
	repositoryObj := NewPlayerRepository()
	if repositoryObj == nil {
		t.Errorf("NewPlayerRepository return nil, expected return PlayerRepository")
	}
	if fmt.Sprintf("%T", repositoryObj) != "*repositories.playerRepository" {
		t.Errorf("NewPlayerRepository return wrong type, expected '*repositories.playerRepository'")
	}
	if _, ok := repositoryObj.(PlayerRepository); !ok {
		t.Errorf("NewPlayerRepository returns not implements PlayerRepository interface")
	}
}

func TestPlayerRepositoryStorePlayer(t *testing.T) {
	type input struct {
		ctx    context.Context
		player models.Player
	}
	type output struct {
		player models.Player
		err    error
	}
	type mockConfig struct {
		in  input
		out output
	}

	var (
		notExistTeam = models.Player{
			TeamName: "test not exist team name",
			Name:     "Test player name",
		}
		existPlayer = models.Player{
			TeamName: "Nankatsu",
			Name:     "Tsubasa Ozora",
		}
		player = models.Player{
			TeamName: "Nankatsu",
			Name:     "Test player name",
		}
	)

	testCases := []struct {
		name          string
		given         input
		expected      output
		configureMock func(*mockConfig)
	}{
		{
			name: "failed, team already exist",
			given: input{
				player: notExistTeam,
			},
			expected: output{
				err: exceptions.ErrTeamNotFound,
			},
			configureMock: func(config *mockConfig) {},
		},
		{
			name: "failed, player already exist in team",
			given: input{
				player: existPlayer,
			},
			expected: output{
				err: exceptions.ErrPlayerAlreadyExistInTeam,
			},
			configureMock: func(config *mockConfig) {},
		},
		{
			name: "success, store player",
			given: input{
				player: player,
			},
			expected: output{
				err:    nil,
				player: player,
			},
			configureMock: func(config *mockConfig) {},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.configureMock(&mockConfig{
				in:  test.given,
				out: test.expected,
			})

			r := playerRepository{
				db: db.Inmemory(),
			}

			result, err := r.StorePlayer(context.TODO(), test.given.player)
			if expected := test.expected.err; (err == nil && expected != nil) || (err != nil && !errors.Is(err, expected)) {
				t.Errorf("error:\nEXPECTED: %v\nGOT: %v\n",
					expected, err)
			}

			if expected := test.expected.player; !reflect.DeepEqual(result, expected) {
				t.Errorf("result:\nEXPECTED:\n%+v\nGOT:\n%+v\n",
					expected, result)
			}
		})
	}
}

func TestPlayerRepositoryGetPlayer(t *testing.T) {
	type input struct {
		ctx  context.Context
		name string
	}
	type output struct {
		players []models.Player
		err     error
	}
	type mockConfig struct {
		in  input
		out output
	}

	var (
		playerNotExistName = "test not exist player name"
		playerName         = "Tsubasa Ozora"
		players            = db.Inmemory().IndexPlayer[playerName]
	)

	testCases := []struct {
		name          string
		given         input
		expected      output
		configureMock func(*mockConfig)
	}{
		{
			name: "failed, player not found",
			given: input{
				name: playerNotExistName,
			},
			expected: output{
				err:     exceptions.ErrPlayerNotFound,
				players: []models.Player{},
			},
			configureMock: func(config *mockConfig) {},
		},
		{
			name: "success, get players",
			given: input{
				name: playerName,
			},
			expected: output{
				err:     nil,
				players: players,
			},
			configureMock: func(config *mockConfig) {},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.configureMock(&mockConfig{
				in:  test.given,
				out: test.expected,
			})

			r := playerRepository{
				db: db.Inmemory(),
			}

			result, err := r.GetPlayer(context.TODO(), test.given.name)
			if expected := test.expected.err; (err == nil && expected != nil) || (err != nil && !errors.Is(err, expected)) {
				t.Errorf("error:\nEXPECTED: %v\nGOT: %v\n",
					expected, err)
			}

			if expected := test.expected.players; !reflect.DeepEqual(result, expected) {
				t.Errorf("result:\nEXPECTED:\n%+v\nGOT:\n%+v\n",
					expected, result)
			}
		})
	}
}
