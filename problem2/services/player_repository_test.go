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

func TestNewPlayerService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockPlayerRepository(ctrl)

	serviceObj := NewPlayerService(repository)
	if serviceObj == nil {
		t.Errorf("NewPlayerService return nil, expected return PlayerService")
	}
	if fmt.Sprintf("%T", serviceObj) != "*services.playerService" {
		t.Errorf("NewPlayerService return wrong type, expected '*services.playerService'")
	}
	if _, ok := serviceObj.(PlayerService); !ok {
		t.Errorf("NewPlayerService returns not implements PlayerService interface")
	}
}

func TestPlayerServiceStorePlayer(t *testing.T) {
	type input struct {
		ctx    context.Context
		player models.Player
	}
	type output struct {
		player models.Player
		err    error
	}
	type mockConfig struct {
		in         input
		out        output
		repository *mocks.MockPlayerRepository
	}

	var (
		defaultErr = errors.New("default error")
		player     = models.Player{
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
			name: "failed, default error",
			given: input{
				player: player,
			},
			expected: output{
				err:    defaultErr,
				player: models.Player{},
			},
			configureMock: func(config *mockConfig) {
				config.repository.EXPECT().
					StorePlayer(gomock.Any(), player).
					Return(models.Player{}, defaultErr)
			},
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
			configureMock: func(config *mockConfig) {
				config.repository.EXPECT().
					StorePlayer(gomock.Any(), player).
					Return(player, nil)
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockPlayerRepository(ctrl)

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.configureMock(&mockConfig{
				in:         test.given,
				out:        test.expected,
				repository: repository,
			})

			s := playerService{
				repository: repository,
			}

			result, err := s.StorePlayer(context.TODO(), test.given.player)
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

func TestPlayerServiceGetPlayer(t *testing.T) {
	type input struct {
		ctx  context.Context
		name string
	}
	type output struct {
		players []models.Player
		err     error
	}
	type mockConfig struct {
		in         input
		out        output
		repository *mocks.MockPlayerRepository
	}

	var (
		defaultErr = errors.New("default error")
		playerName = "Tsubasa Ozora"
		players    = db.Inmemory().IndexPlayer[playerName]
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
				name: playerName,
			},
			expected: output{
				err:     defaultErr,
				players: []models.Player{},
			},
			configureMock: func(config *mockConfig) {
				config.repository.EXPECT().
					GetPlayer(gomock.Any(), config.in.name).
					Return([]models.Player{}, defaultErr)
			},
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
			configureMock: func(config *mockConfig) {
				config.repository.EXPECT().
					GetPlayer(gomock.Any(), config.in.name).
					Return(players, nil)
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockPlayerRepository(ctrl)

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.configureMock(&mockConfig{
				in:         test.given,
				out:        test.expected,
				repository: repository,
			})

			r := playerService{
				repository: repository,
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
