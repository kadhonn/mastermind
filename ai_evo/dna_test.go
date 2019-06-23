package ai_evo

import (
	"github.com/kadhonn/mastermind/ai"
	"github.com/kadhonn/mastermind/mastermind"
	"testing"
)

func TestDoEvoEval1(t *testing.T) {
	game := createTestGame()

	result := doEvoEval(game, DNA{
		Nucl: []interface{}{
			&Action{
				Field: &Field{0},
				Color: &FixedColor{1},
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
		Nucl: []interface{}{
			&ColorCompare{
				Skip:   1,
				Equals: true,
				First:  &FixedColor{1},
				Second: &FixedColor{1},
			}, &Action{
				Field: &Field{0},
				Color: &FixedColor{1},
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
		Nucl: []interface{}{
			&ColorCompare{
				Skip:   1,
				Equals: true,
				First:  &FixedColor{1},
				Second: &FixedColor{2},
			}, &Action{
				Field: &Field{0},
				Color: &FixedColor{1},
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
		Nucl: []interface{}{
			&ColorCompare{
				Skip:   1,
				Equals: false,
				First:  &FixedColor{1},
				Second: &FixedColor{1},
			}, &Action{
				Field: &Field{0},
				Color: &FixedColor{1},
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
		Nucl: []interface{}{
			&ColorCompare{
				Skip:   1,
				Equals: false,
				First:  &FixedColor{1},
				Second: &FixedColor{1},
			}, &Action{
				Field: &Field{0},
				Color: &FixedColor{1},
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
		Nucl: []interface{}{
			&ColorCompare{
				Skip:   1,
				Equals: true,
				First:  &Field{7},
				Second: &FixedColor{1},
			}, &Action{
				Field: &Field{0},
				Color: &FixedColor{1},
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
		Nucl: []interface{}{
			&ColorCompare{
				Skip:   2,
				Equals: true,
				First:  &Field{6},
				Second: &FixedColor{0},
			}, &ColorCompare{
				Skip:   1,
				Equals: true,
				First:  &Field{6},
				Second: &FixedColor{-1},
			}, &Action{
				Field: &Field{0},
				Color: &FixedColor{1},
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
		Nucl: []interface{}{
			&ColorCompare{
				Skip:   2,
				Equals: true,
				First:  &Field{6},
				Second: &FixedColor{0},
			}, &ColorCompare{
				Skip:   1,
				Equals: true,
				First:  &Field{6},
				Second: &FixedColor{0},
			}, &Action{
				Field: &Field{0},
				Color: &FixedColor{1},
			},
		},
	})
	if result[0] != 1 {
		t.Error("result wrong")
	}
}
func TestDoEvoEval9(t *testing.T) {
	game := createTestGame()
	game.GetMoves()[0] = []int{
		0, 1, 0, 0, 0, 0,
	}
	game.GetPoints()[0] = &mastermind.PointsData{Black: 0, White: 0}

	result := doEvoEval(game, DNA{
		Nucl: []interface{}{
			&PointsCompare{
				Skip:   1,
				Mode:   0,
				Blacks: true,
				Count:  -1,
			}, &PointsCompare{
				Skip:   1,
				Mode:   0,
				Blacks: true,
				Count:  0,
			}, &Action{
				Field: &Field{0},
				Color: &FixedColor{1},
			},
		},
	})
	if result[0] != 0 {
		t.Error("result wrong")
	}
}
func TestDoEvoEval10(t *testing.T) {
	game := createTestGame()
	game.GetMoves()[0] = []int{
		0, 1, 0, 0, 0, 0,
	}
	game.GetPoints()[0] = &mastermind.PointsData{Black: 0, White: 0}

	result := doEvoEval(game, DNA{
		Nucl: []interface{}{
			&PointsCompare{
				Skip:   1,
				Mode:   0,
				Blacks: true,
				Count:  -1,
			}, &Action{
				Field: &Field{0},
				Color: &FixedColor{1},
			},
		},
	})
	if result[0] != 1 {
		t.Error("result wrong")
	}
}
func TestDoEvoEval11(t *testing.T) {
	game := createTestGame()
	game.GetMoves()[0] = []int{
		0, 1, 0, 0, 0, 0,
	}
	game.GetPoints()[0] = &mastermind.PointsData{Black: 0, White: 2}

	result := doEvoEval(game, DNA{
		Nucl: []interface{}{
			&PointsCompare{
				Skip:   1,
				Mode:   1,
				Blacks: false,
				Count:  1,
			}, &Action{
				Field: &Field{0},
				Color: &FixedColor{1},
			},
		},
	})
	if result[0] != 0 {
		t.Error("result wrong")
	}
}
func TestMarshalUnmarshal(t *testing.T) {
	for i := 0; i < 100; i++ {
		ai.StartEvaluationWithTime(EvoEval(Load(Save(CreateRandomDNA(6, 10, 1000)))), 1, 1)
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
