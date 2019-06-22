package ai_evo

import (
	"github.com/kadhonn/mastermind/mastermind"
	"testing"
)

func TestDoEvoEval1(t *testing.T) {
	game := createTestGame()

	result := doEvoEval(game, DNA{
		nucl: []interface{}{
			Action{
				field: Field{0},
				color: FixedColor{1},
			},
		},
	})
	if result[0] != 1 {
		t.Error("result should be 1")
	}
}
func TestDoEvoEval2(t *testing.T) {
	game := createTestGame()

	result := doEvoEval(game, DNA{
		nucl: []interface{}{
			ColorCompare{
				skip:   1,
				equals: true,
				first:  FixedColor{1},
				second: FixedColor{1},
			}, Action{
				field: Field{0},
				color: FixedColor{1},
			},
		},
	})
	if result[0] != 0 {
		t.Error("result should be 0")
	}
}

func TestDoEvoEval3(t *testing.T) {
	game := createTestGame()

	result := doEvoEval(game, DNA{
		nucl: []interface{}{
			ColorCompare{
				skip:   1,
				equals: true,
				first:  FixedColor{1},
				second: FixedColor{2},
			}, Action{
				field: Field{0},
				color: FixedColor{1},
			},
		},
	})
	if result[0] != 1 {
		t.Error("result wrong")
	}
}

func TestDoEvoEval4(t *testing.T) {
	game := createTestGame()

	result := doEvoEval(game, DNA{
		nucl: []interface{}{
			ColorCompare{
				skip:   1,
				equals: false,
				first:  FixedColor{1},
				second: FixedColor{1},
			}, Action{
				field: Field{0},
				color: FixedColor{1},
			},
		},
	})
	if result[0] != 1 {
		t.Error("result wrong")
	}
}
func TestDoEvoEval5(t *testing.T) {
	game := createTestGame()

	result := doEvoEval(game, DNA{
		nucl: []interface{}{
			ColorCompare{
				skip:   1,
				equals: false,
				first:  FixedColor{1},
				second: FixedColor{1},
			}, Action{
				field: Field{0},
				color: FixedColor{1},
			},
		},
	})
	if result[0] != 1 {
		t.Error("result wrong")
	}
}
func TestDoEvoEval6(t *testing.T) {
	game := createTestGame()
	game.GetMoves()[0] = []int{
		0, 1, 0, 0, 0, 0,
	}

	result := doEvoEval(game, DNA{
		nucl: []interface{}{
			ColorCompare{
				skip:   1,
				equals: true,
				first:  Field{7},
				second: FixedColor{1},
			}, Action{
				field: Field{0},
				color: FixedColor{1},
			},
		},
	})
	if result[0] != 1 {
		t.Error("result wrong")
	}
}
func TestDoEvoEval7(t *testing.T) {
	game := createTestGame()
	game.GetMoves()[0] = []int{
		0, 1, 0, 0, 0, 0,
	}

	result := doEvoEval(game, DNA{
		nucl: []interface{}{
			ColorCompare{
				skip:   2,
				equals: true,
				first:  Field{6},
				second: FixedColor{0},
			}, ColorCompare{
				skip:   1,
				equals: true,
				first:  Field{6},
				second: FixedColor{-1},
			}, Action{
				field: Field{0},
				color: FixedColor{1},
			},
		},
	})
	if result[0] != 0 {
		t.Error("result wrong")
	}
}
func TestDoEvoEval8(t *testing.T) {
	game := createTestGame()
	game.GetMoves()[0] = []int{
		0, 1, 0, 0, 0, 0,
	}

	result := doEvoEval(game, DNA{
		nucl: []interface{}{
			ColorCompare{
				skip:   2,
				equals: true,
				first:  Field{6},
				second: FixedColor{0},
			}, ColorCompare{
				skip:   1,
				equals: true,
				first:  Field{6},
				second: FixedColor{0},
			}, Action{
				field: Field{0},
				color: FixedColor{1},
			},
		},
	})
	if result[0] != 1 {
		t.Error("result wrong")
	}
}
func createTestGame() *mastermind.GameData {
	return &mastermind.GameData{
		MoveSize:   6,
		ColorCount: 10,
		Moves:      make([][]int, 15),
		Points:     make([]mastermind.Points, 15),
		Secret:     []int{0, 0, 0, 0, 0, 0},
	}
}
