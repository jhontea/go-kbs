package services

import (
	"context"
	"errors"
	"fmt"
	"go-kbs-soccer/entities/models"
	db "go-kbs-soccer/infra/inmemoryDB"
	"go-kbs-soccer/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewTeamService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockTeamRepository(ctrl)

	serviceObj := NewTeamService(repository)
	if serviceObj == nil {
		t.Errorf("NewTeamService return nil, expected return TeamService")
	}
	if fmt.Sprintf("%T", serviceObj) != "*services.teamService" {
		t.Errorf("NewTeamService return wrong type, expected '*services.teamService'")
	}
	if _, ok := serviceObj.(TeamService); !ok {
		t.Errorf("NewTeamService returns not implements TeamService interface")
	}
}

func TestTeamServiceStoreTeam(t *testing.T) {
	type input struct {
		ctx  context.Context
		team models.Team
	}
	type output struct {
		team models.Team
		err  error
	}
	type mockConfig struct {
		in         input
		out        output
		repository *mocks.MockTeamRepository
	}

	var (
		defaultErr = errors.New("default error")
		team       = models.Team{
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
			name: "failed, default error",
			given: input{
				team: team,
				ctx:  context.TODO(),
			},
			expected: output{
				err:  defaultErr,
				team: models.Team{},
			},
			configureMock: func(config *mockConfig) {
				config.repository.EXPECT().
					StoreTeam(config.in.ctx, config.in.team).
					Return(models.Team{}, defaultErr)
			},
		},
		{
			name: "success, store team",
			given: input{
				team: team,
				ctx:  context.TODO(),
			},
			expected: output{
				err:  nil,
				team: team,
			},
			configureMock: func(config *mockConfig) {
				config.repository.EXPECT().
					StoreTeam(config.in.ctx, config.in.team).
					Return(config.in.team, nil)
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockTeamRepository(ctrl)

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.configureMock(&mockConfig{
				in:         test.given,
				out:        test.expected,
				repository: repository,
			})

			r := teamService{
				repository: repository,
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

func TestTeamServiceGetTeam(t *testing.T) {
	type input struct {
		ctx  context.Context
		name string
	}
	type output struct {
		teamPlayers models.TeamPlayers
		err         error
	}
	type mockConfig struct {
		in         input
		out        output
		repository *mocks.MockTeamRepository
	}

	var (
		defaultErr   = errors.New("default error")
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
			name: "failed, default error",
			given: input{
				ctx:  context.TODO(),
				name: notExistTeam.Name,
			},
			expected: output{
				err:         defaultErr,
				teamPlayers: models.TeamPlayers{},
			},
			configureMock: func(config *mockConfig) {
				config.repository.EXPECT().
					GetTeam(gomock.Any(), config.in.name).
					Return(models.TeamPlayers{}, defaultErr)
			},
		},
		{
			name: "success, get team players",
			given: input{
				ctx:  context.TODO(),
				name: team.Name,
			},
			expected: output{
				err:         nil,
				teamPlayers: teamPlayers,
			},
			configureMock: func(config *mockConfig) {
				config.repository.EXPECT().
					GetTeam(gomock.Any(), config.in.name).
					Return(teamPlayers, nil)
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockTeamRepository(ctrl)

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.configureMock(&mockConfig{
				in:         test.given,
				out:        test.expected,
				repository: repository,
			})

			r := teamService{
				repository: repository,
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
