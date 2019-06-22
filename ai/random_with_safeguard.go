package ai

import (
	"github.com/kadhonn/mastermind/mastermind"
)

func RandomWithSafeGuard(game mastermind.Game) []int {
	guess := make([]int, game.GetMoveSize())

	for true {
		createGuess(guess, game.GetColorCount())
		if isCompatible(game, guess) {
			break
		}
	}

	return guess
}

func createGuess(guess []int, colorCount int) {
	for index := range guess {
		guess[index] = r1.Intn(colorCount)
	}
}
