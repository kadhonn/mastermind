package mastermind

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Game interface {
	GetMoveSize() int
	GetColorCount() int
	GetMoves() [][]int
	GetSecret() []int
	MakeMove(move []int) error
	GetPoints() []Points
	HasWon() bool
	HasLost() bool
}

type Points interface {
	GetBlack() int
	GetWhite() int
}

type PointsData struct {
	Black int
	White int
}

type GameData struct {
	MoveSize   int
	ColorCount int
	Moves      [][]int
	Secret     []int
	Points     []Points
}

func (game *GameData) GetMoves() [][]int {
	return game.Moves
}

func (game *GameData) GetSecret() []int {
	return game.Secret
}

func (game *GameData) GetMoveSize() int {
	return game.MoveSize
}

func (game *GameData) GetColorCount() int {
	return game.ColorCount
}

func (points *PointsData) GetBlack() int {
	return points.Black
}

func (points *PointsData) GetWhite() int {
	return points.White
}

func (game *GameData) MakeMove(move []int) error {
	if move == nil {
		return errors.New("move cannot be nil")
	}
	if len(move) != game.MoveSize {
		return errors.New(fmt.Sprintf("move is size %d but move size is %d", len(move), game.MoveSize))
	}
	for index, element := range move {
		if element < 0 || element >= game.ColorCount {
			return errors.New(fmt.Sprintf("color at index %d is invalid: %d", index, element))
		}
	}

	lastMoveIndex, err := findNextIndex(game)
	if err != nil {
		return err
	}

	game.Moves[lastMoveIndex] = move
	game.Points[lastMoveIndex] = calcPoints(game.Secret, move)

	return nil
}

func calcPoints(origsecret []int, origmove []int) Points {
	secret := make([]int, len(origsecret))
	copy(secret, origsecret)
	move := make([]int, len(origmove))
	copy(move, origmove)
	points := PointsData{
		Black: 0,
		White: 0,
	}

	for i := 0; i < len(secret); i++ {
		if secret[i] == move[i] {
			secret[i] = -1
			move[i] = -1
			points.Black++
		}
	}

	for _, guess := range move {
		if guess != -1 {
			for secretI, color := range secret {
				if color == guess {
					secret[secretI] = -1
					points.White++
					break
				}
			}
		}
	}
	return &points
}

func (game *GameData) GetPoints() []Points {
	return game.Points
}

func (game *GameData) HasWon() bool {
	index, err := findNextIndex(game)
	if err != nil {
		index = len(game.Moves)
	}
	index--
	if index < 0 {
		return false
	}
	return game.Points[index].GetBlack() == game.MoveSize
}

func (game *GameData) HasLost() bool {
	if game.HasWon() {
		return false
	}
	if game.Moves[len(game.Moves)-1] != nil {
		return true
	}
	return false
}

func findNextIndex(data *GameData) (int, error) {
	for index, element := range data.Moves {
		if element == nil {
			return index, nil
		}
	}
	return 0, errors.New("out of moves")
}

func StartGame() Game {
	moveSize := 6
	colorCount := 10
	moveCount := 1000
	return &GameData{
		MoveSize:   moveSize,
		ColorCount: colorCount,
		Moves:      make([][]int, moveCount),
		Secret:     getSecret(moveSize, colorCount),
		Points:     make([]Points, moveCount),
	}
}

var s1 = rand.NewSource(time.Now().UnixNano() - 100)
var r1 = rand.New(s1)

func getSecret(moveSize int, colors int) []int {
	secret := make([]int, moveSize)
	for index := range secret {
		secret[index] = r1.Intn(colors)
	}
	return secret
}
