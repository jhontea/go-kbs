package db

import (
	"go-kbs-soccer/entities/models"
	"sync"
)

var (
	Meiwa    = "Meiwa"
	Nankatsu = "Nankatsu"

	teamMeiwa = models.Team{
		Name: Meiwa,
	}
	teamNankatsu = models.Team{
		Name: Nankatsu,
	}

	dummyTeam = []models.Team{
		teamMeiwa,
		teamNankatsu,
	}
	indexTeam = map[string]models.Team{
		Meiwa:    teamMeiwa,
		Nankatsu: teamNankatsu,
	}

	hyuga            = "Kojiro Hyuga"
	wakashimazu      = "Ken Wakashimazu"
	sawada           = "Takeshi Sawada"
	tsubasa          = "Tsubasa Ozora"
	misaki           = "Taro Misaki"
	wakabayashi      = "Genzo Wakabayashi"
	playerMeiwaHyuga = models.Player{
		TeamName: Meiwa,
		Name:     hyuga,
	}
	playerMeiwaWakashimazu = models.Player{
		TeamName: Meiwa,
		Name:     wakashimazu,
	}
	playerMeiwaSawada = models.Player{
		TeamName: Meiwa,
		Name:     sawada,
	}
	playerNankatsuTsubasa = models.Player{
		TeamName: Nankatsu,
		Name:     tsubasa,
	}
	playerNankatsuMisaki = models.Player{
		TeamName: Nankatsu,
		Name:     misaki,
	}
	playerNankatsuWakabayashi = models.Player{
		TeamName: Nankatsu,
		Name:     wakabayashi,
	}

	dummyPlayer = []models.Player{
		playerMeiwaHyuga,
		playerMeiwaWakashimazu,
		playerMeiwaSawada,
		playerNankatsuTsubasa,
		playerNankatsuMisaki,
		playerNankatsuWakabayashi,
	}

	indexTeamPlayer = map[string][]models.Player{
		Meiwa: {
			playerMeiwaHyuga,
			playerMeiwaSawada,
			playerMeiwaWakashimazu,
		},
		Nankatsu: {
			playerNankatsuMisaki,
			playerNankatsuTsubasa,
			playerNankatsuWakabayashi,
		},
	}

	indexPlayer = map[string][]models.Player{
		hyuga:       {playerMeiwaHyuga},
		wakashimazu: {playerMeiwaWakashimazu},
		sawada:      {playerMeiwaSawada},
		tsubasa:     {playerNankatsuTsubasa},
		misaki:      {playerNankatsuMisaki},
		wakabayashi: {playerNankatsuWakabayashi},
	}
)

var (
	dbConnOnce sync.Once
	inmemory   *InmemoryDB
)

type InmemoryDB struct {
	Teams           []models.Team
	IndexTeam       map[string]models.Team
	IndexTeamPlayer map[string][]models.Player
	Players         []models.Player
	IndexPlayer     map[string][]models.Player
}

func Inmemory() *InmemoryDB {
	dbConnOnce.Do(func() {
		inmemory = &InmemoryDB{
			Teams:           dummyTeam,
			IndexTeam:       indexTeam,
			IndexTeamPlayer: indexTeamPlayer,
			Players:         dummyPlayer,
			IndexPlayer:     indexPlayer,
		}
	})
	return inmemory
}
