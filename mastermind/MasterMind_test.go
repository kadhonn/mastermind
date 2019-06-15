package mastermind

import "testing"

func TestDefaultStartGameReturnsSomething(t *testing.T) {
	game := StartGame()
	if game == nil {
		t.Fatal("game is nil")
	}
	if game.GetMoveSize() != 6 {
		t.Error("default game size should be 6")
	}
	if game.GetColorCount() != 10 {
		t.Error("default game color count should be 10")
	}
	if len(game.GetMoves()) != 15 {
		t.Error("default game should have 15 rounds")
	}
	if game.GetMoves()[0] != nil {
		t.Error("default game should not have rounds after beginning")
	}
	if game.GetSecret() != nil {
		t.Error("default game should have secret after start")
	}
}

func TestMakeInvalidMove(t *testing.T) {
	game := &GameData{
		moveSize:   2,
		colorCount: 2,
		moves:      make([][]int, 2),
		secret:     []int{1, 1},
	}

	err := game.MakeMove([]int{1})
	if err == nil {
		t.Error("should have returned an error when called with invalid move size")
	}

	err = game.MakeMove(nil)
	if err == nil {
		t.Error("should have returned an error when called with nil move")
	}

	err = game.MakeMove([]int{0, 2})
	if err == nil {
		t.Error("should have returned an error when called with out of bounds move")
	}
}

func TestMakeTooManyMoves(t *testing.T) {
	game := &GameData{
		moveSize:   2,
		colorCount: 2,
		moves:      make([][]int, 2),
		secret:     []int{1, 1},
	}

	err := game.MakeMove([]int{0, 0})
	if err != nil {
		t.Fatal("unwanted error", err)
	}

	err = game.MakeMove([]int{0, 0})
	if err != nil {
		t.Fatal("unwanted error", err)
	}

	err = game.MakeMove([]int{0, 0})
	if err == nil {
		t.Fatal("wanted error but got none!")
	}
}

func TestMakeValidMove(t *testing.T) {
	game := &GameData{
		moveSize:   2,
		colorCount: 2,
		moves:      make([][]int, 2),
		secret:     []int{1, 1},
	}

	err := game.MakeMove([]int{1, 1})
	if err != nil {
		t.Fatal("unwanted error", err)
	}

	if game.GetMoves()[0] == nil || game.GetMoves()[1] != nil {
		t.Fatal("something wrong with array entries")
	}

	if game.GetMoves()[0][0] != 1 || game.GetMoves()[0][1] != 1 {
		t.Fatal("something wrong with saved move")
	}
}
