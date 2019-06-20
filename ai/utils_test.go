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

	statistics := startEvaluationWithCreator(func() Evaluator { return solver }, creator)

	if statistics.Total != statistics.Won {
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

	statistics := startEvaluationWithCreator(func() Evaluator { return solver }, creator)

	if statistics.Won != 0 {
		t.Fatal("should have lost all games!")
	}
}

func createTestGame() *mastermind.GameData {
	return &mastermind.GameData{
		MoveSize:   2,
		ColorCount: 4,
		Moves:      make([][]int, 4),
		Points:     make([]mastermind.Points, 4),
		Secret:     []int{1, 1},
	}
}

func TestIsGuessCompatible(t *testing.T) {
	if !isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 0, White: 0}, []int{3, 3, 3}) {
		t.Error("should be compatible!")
	}
	if !isGuessCompatible([]int{0, 1}, &mastermind.PointsData{Black: 0, White: 1}, []int{1, 2}) {
		t.Error("should be compatible!")
	}
	if !isGuessCompatible([]int{0, 1}, &mastermind.PointsData{Black: 1, White: 0}, []int{2, 1}) {
		t.Error("should be compatible!")
	}
	if !isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 1, White: 1}, []int{0, 2, 2}) {
		t.Error("should be compatible!")
	}
	if !isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 1, White: 1}, []int{0, 2, 3}) {
		t.Error("should be compatible!")
	}
	if !isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 1, White: 2}, []int{0, 2, 1}) {
		t.Error("should be compatible!")
	}
	if !isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 2, White: 0}, []int{3, 1, 2}) {
		t.Error("should be compatible!")
	}
	if !isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 3, White: 0}, []int{0, 1, 2}) {
		t.Error("should be compatible!")
	}
	if !isGuessCompatible([]int{0, 0, 2}, &mastermind.PointsData{Black: 1, White: 0}, []int{0, 1, 3}) {
		t.Error("should be compatible!")
	}
	if !isGuessCompatible([]int{0, 0, 2}, &mastermind.PointsData{Black: 1, White: 0}, []int{1, 0, 3}) {
		t.Error("should be compatible!")
	}
	if !isGuessCompatible([]int{0, 0, 2}, &mastermind.PointsData{Black: 1, White: 0}, []int{1, 1, 2}) {
		t.Error("should be compatible!")
	}

	if isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 0, White: 0}, []int{3, 0, 3}) {
		t.Error("should NOT be compatible!")
	}
	if isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 0, White: 1}, []int{3, 3, 3}) {
		t.Error("should NOT be compatible!")
	}
	if isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 0, White: 1}, []int{3, 2, 0}) {
		t.Error("should NOT be compatible!")
	}
	if isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 0, White: 1}, []int{3, 1, 3}) {
		t.Error("should NOT be compatible!")
	}
	if isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 1, White: 0}, []int{3, 0, 3}) {
		t.Error("should NOT be compatible!")
	}
	if isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 1, White: 0}, []int{3, 3, 3}) {
		t.Error("should NOT be compatible!")
	}
	if isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 1, White: 0}, []int{0, 1, 3}) {
		t.Error("should NOT be compatible!")
	}
	if isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 1, White: 1}, []int{0, 1, 3}) {
		t.Error("should NOT be compatible!")
	}
	if isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 1, White: 1}, []int{0, 0, 3}) {
		t.Error("should NOT be compatible!")
	}
	if isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 1, White: 1}, []int{3, 3, 1}) {
		t.Error("should NOT be compatible!")
	}
	if isGuessCompatible([]int{0, 1, 2}, &mastermind.PointsData{Black: 3, White: 0}, []int{3, 3, 1}) {
		t.Error("should NOT be compatible!")
	}
	if isGuessCompatible([]int{0, 0, 0}, &mastermind.PointsData{Black: 1, White: 0}, []int{0, 0, 1}) {
		t.Error("should NOT be compatible!")
	}
	if isGuessCompatible([]int{0, 0, 1, 2}, &mastermind.PointsData{Black: 0, White: 2}, []int{0, 0, 3, 4}) {
		t.Error("should NOT be compatible!")
	}
}

func TestIsCompatible(t *testing.T) {
	game := createTestGame()
	game.Secret = []int{1, 2}
	game.Moves = [][]int{
		{0, 0},
		{1, 3},
	}
	game.Points = []mastermind.Points{
		&mastermind.PointsData{Black: 0, White: 0},
		&mastermind.PointsData{Black: 1, White: 0},
	}

	if !isCompatible(game, []int{1, 2}) {
		t.Error("should be compatible!")
	}
	if !isCompatible(game, []int{2, 3}) {
		t.Error("should be compatible!")
	}

	if isCompatible(game, []int{3, 1}) {
		t.Error("should NOT be compatible!")
	}
	if isCompatible(game, []int{0, 1}) {
		t.Error("should NOT be compatible!")
	}
}
