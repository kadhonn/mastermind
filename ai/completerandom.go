package ai

import (
	"mastermind/mastermind"
	"math/rand"
	"time"
)

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

func CompleteRandom(game mastermind.Game) []int {
	guess := make([]int, game.GetMoveSize())

	for index := range guess {
		guess[index] = r1.Intn(game.GetColorCount())
	}

	return guess
}
