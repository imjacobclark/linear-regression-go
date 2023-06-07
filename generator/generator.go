package generator

import (
	"math/rand"
	"time"
)

type Genre int

const (
	RPG Genre = iota
	FPS
	Puzzle
	Adventure
	Strategy
	Racing
	Simulator
	GenreCount
)

func (g Genre) String() string {
	return [...]string{"RPG", "FPS", "Puzzle", "Adventure", "Strategy", "Racing", "Simulator"}[g]
}

type Game struct {
	ID							int
	Genre						Genre
	DeveloperReputationScore	float64
	MarketingBudget				int
	PreReleaseHypeScore			float64
	PlatformCount				int
	GameLength					int
	CopiesSold					int
}

func GenerateData() []Game {
	rand.Seed(time.Now().UnixNano())

	var games []Game

	for i := 1; i <= 1; i++ {
		game := Game{
			ID:                   	  i,
			Genre:                    Genre(rand.Intn(int(GenreCount))),
			DeveloperReputationScore: rand.Float64() * 5,
			MarketingBudget:          rand.Intn(3000000) + 500000,
			PreReleaseHypeScore:      rand.Float64() * 10,
			PlatformCount:            rand.Intn(7),
			GameLength:               rand.Intn(100) + 10,
			CopiesSold:               rand.Intn(2000000) + 100000,
		}
		
		games = append(games, game)
	}

	return games
}