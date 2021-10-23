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

func TestNewTeamRepository(t *testing.T) {
	repositoryObj := NewTeamRepository()
	if repositoryObj == nil {
		t.Errorf("NewTeamRepository return nil, expected return TeamRepository")
	}
	if fmt.Sprintf("%T", repositoryObj) != "*repositories.teamRepository" {
		t.Errorf("NewTeamRepository return wrong type, expected '*repositories.teamRepository'")
	}
	if _, ok := repositoryObj.(TeamRepository); !ok {
		t.Errorf("NewTeamRepository returns not implements TeamRepository interface")
	}
}

func TestTeamRepositoryStoreTeam(t *testing.T) {
	type input struct {
		ctx  context.Context
		team models.Team
	}
	type output struct {
		team models.Team
		err  error
	}
	type mockConfig struct {
		in  input
		out output
	}

	var (
		existTeam = models.Team{
			Name: "Nankatsu",
		}
		team = models.Team{
			Name: "test store team name",
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
				team: existTeam,
			},
			expected: output{
				err: exceptions.ErrTeamAlreadyExist,
			},
			configureMock: func(config *mockConfig) {},
		},
		{
			name: "success, store team",
			given: input{
				team: team,
			},
			expected: output{
				err:  nil,
				team: team,
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

			r := teamRepository{
				db: db.Inmemory(),
			}

			result, err := r.StoreTeam(context.TODO(), test.given.team)
			if expected := test.expected.err; (err == nil && expected != nil) || (err != nil && !errors.Is(err, expected)) {
				t.Errorf("error:\nEXPECTED: %v\nGOT: %v\n",
					expected, err)
			}

			if expected := test.expected.team; !reflect.DeepEqual(result, expected) {
				t.Errorf("result:\nEXPECTED:\n%+v\nGOT:\n%+v\n",
					expected, result)
			}
		})
	}
}

func TestTeamRepositoryGetTeam(t *testing.T) {
	type input struct {
		ctx  context.Context
		name string
	}
	type output struct {
		teamPlayers models.TeamPlayers
		err         error
	}
	type mockConfig struct {
		in  input
		out output
	}

	var (
		notExistTeam = models.Team{
			Name: "test not exist team name",
		}
		team = models.Team{
			Name: "Nankatsu",
		}
		teamPlayers = models.TeamPlayers{
			Name:    team.Name,
			Players: db.Inmemory().IndexTeamPlayer[team.Name],
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
				name: notExistTeam.Name,
			},
			expected: output{
				err: exceptions.ErrTeamNotFound,
			},
			configureMock: func(config *mockConfig) {},
		},
		{
			name: "success, get team players",
			given: input{
				name: team.Name,
			},
			expected: output{
				err:         nil,
				teamPlayers: teamPlayers,
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

			r := teamRepository{
				db: db.Inmemory(),
			}

			result, err := r.GetTeam(context.TODO(), test.given.name)
			if expected := test.expected.err; (err == nil && expected != nil) || (err != nil && !errors.Is(err, expected)) {
				t.Errorf("error:\nEXPECTED: %v\nGOT: %v\n",
					expected, err)
			}

			if expected := test.expected.teamPlayers; !reflect.DeepEqual(result, expected) {
				t.Errorf("result:\nEXPECTED:\n%+v\nGOT:\n%+v\n",
					expected, result)
			}
		})
	}
}
