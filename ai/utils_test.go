package ai

import (
	"mastermind/mastermind"
	"testing"
)

func TestStartEvaluationEverytimeRight(t *testing.T) {
	creator := func() mastermind.Game {
		game := createTestGame()
		game.Secret = []int{0, 1}
		return game
	}

	solver := func(data mastermind.Game) []int {
		return []int{0, 1}
	}

	statistics := startEvaluation(solver, creator)

	if statistics.total != statistics.won {
		t.Fatal("should have won all games!")
	}
}

func TestStartEvaluationEverytimeWrong(t *testing.T) {
	creator := func() mastermind.Game {
		game := createTestGame()
		game.Secret = []int{0, 1}
		return game
	}

	solver := func(data mastermind.Game) []int {
		return []int{0, 2}
	}

	statistics := startEvaluation(solver, creator)

	if statistics.won != 0 {
		t.Fatal("should have lost all games!")
	}
}

func createTestGame() *mastermind.GameData {
	return &mastermind.GameData{
		MoveSize:   2,
		ColorCount: 4,
		Moves:      make([][]int, 2),
		Points:     make([]mastermind.Points, 2),
		Secret:     []int{1, 1},
	}
}
