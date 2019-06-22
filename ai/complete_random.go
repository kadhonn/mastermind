package ai

import (
	"github.com/kadhonn/mastermind/mastermind"
)

func CompleteRandom(game mastermind.Game) []int {

	guess := make([]int, game.GetMoveSize())

	for index := range guess {
		guess[index] = r1.Intn(game.GetColorCount())
	}

	return guess
}
