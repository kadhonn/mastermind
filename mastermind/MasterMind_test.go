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
	if game.GetSecret() == nil || len(game.GetSecret()) != game.GetMoveSize() {
		t.Error("default game should have valid secret after start")
	}
	if game.GetPoints() == nil || len(game.GetPoints()) != len(game.GetMoves()) {
		t.Error("default game should have valid points array")
	}
}

func TestMakeInvalidMove(t *testing.T) {
	game := createTestGame()

	err := game.MakeMove([]int{1})
	if err == nil {
		t.Error("should have returned an error when called with invalid move size")
	}

	err = game.MakeMove(nil)
	if err == nil {
		t.Error("should have returned an error when called with nil move")
	}

	err = game.MakeMove([]int{0, 4})
	if err == nil {
		t.Error("should have returned an error when called with out of bounds move")
	}
}

func TestMakeTooManyMoves(t *testing.T) {
	game := createTestGame()

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
	game := createTestGame()

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

func TestWinning(t *testing.T) {
	game := createTestGame()
	game.Secret = []int{1, 1}

	if game.HasWon() {
		t.Fatal("should not have won the game")
	}

	err := game.MakeMove([]int{1, 0})
	if err != nil {
		t.Fatal("unwanted error", err)
	}

	if game.HasWon() {
		t.Fatal("should not have won the game")
	}

	err = game.MakeMove([]int{1, 1})
	if err != nil {
		t.Fatal("unwanted error", err)
	}

	if !game.HasWon() {
		t.Fatal("should have won game")
	}

	if game.HasLost() {
		t.Fatal("should have not lost the game")
	}
}

func TestLosing(t *testing.T) {
	game := createTestGame()
	game.Secret = []int{1, 1}

	if game.HasLost() {
		t.Fatal("should not have lost the game")
	}

	err := game.MakeMove([]int{1, 0})
	if err != nil {
		t.Fatal("unwanted error", err)
	}

	if game.HasLost() {
		t.Fatal("should not have lost the game")
	}

	err = game.MakeMove([]int{1, 0})
	if err != nil {
		t.Fatal("unwanted error", err)
	}

	if !game.HasLost() {
		t.Fatal("should have lost the game")
	}

	if game.HasWon() {
		t.Fatal("should have not won the game")
	}
}

func TestAllBlacks(t *testing.T) {
	testCounting(t, []int{1, 1}, []int{1, 1}, 2, 0)
}

func TestFail(t *testing.T) {
	testCounting(t, []int{1, 1}, []int{0, 0}, 0, 0)
}

func TestOneBlack(t *testing.T) {
	testCounting(t, []int{1, 1}, []int{0, 1}, 1, 0)
}

func TestOneWhite(t *testing.T) {
	testCounting(t, []int{0, 1}, []int{1, 2}, 0, 1)
}

func TestAllWhite(t *testing.T) {
	testCounting(t, []int{0, 1}, []int{1, 0}, 0, 2)
}

func testCounting(t *testing.T, secret []int, move []int, blacks int, whites int) {
	game := createTestGame()
	game.Secret = secret

	err := game.MakeMove(move)
	if err != nil {
		t.Fatal("unwanted error", err)
	}

	points := game.GetPoints()[0]
	if points == nil {
		t.Fatal("points are not allowed to be nil after move")
	}

	if points.GetBlack() != blacks || points.GetWhite() != whites {
		t.Fatalf("points are wrong, got %d blacks and %d whites", points.GetBlack(), points.GetWhite())
	}
}

func createTestGame() *GameData {
	return &GameData{
		MoveSize:   2,
		ColorCount: 4,
		Moves:      make([][]int, 2),
		Points:     make([]Points, 2),
		Secret:     []int{1, 1},
	}
}
