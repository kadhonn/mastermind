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
	black int
	white int
}

type GameData struct {
	moveSize   int
	colorCount int
	moves      [][]int
	secret     []int
	points     []Points
}

func (game *GameData) GetMoves() [][]int {
	return game.moves
}

func (game *GameData) GetSecret() []int {
	return game.secret
}

func (game *GameData) GetMoveSize() int {
	return game.moveSize
}

func (game *GameData) GetColorCount() int {
	return game.colorCount
}

func (points *PointsData) GetBlack() int {
	return points.black
}

func (points *PointsData) GetWhite() int {
	return points.white
}

func (game *GameData) MakeMove(move []int) error {
	if move == nil {
		return errors.New("move cannot be nil")
	}
	if len(move) != game.moveSize {
		return errors.New(fmt.Sprintf("move is size %d but move size is %d", len(move), game.moveSize))
	}
	for index, element := range move {
		if element < 0 || element >= game.colorCount {
			return errors.New(fmt.Sprintf("color at index %d is invalid: %d", index, element))
		}
	}

	lastMoveIndex, err := findNextIndex(game)
	if err != nil {
		return err
	}

	game.moves[lastMoveIndex] = move
	game.points[lastMoveIndex] = calcPoints(game.secret, move)

	return nil
}

func calcPoints(origsecret []int, origmove []int) Points {
	secret := make([]int, len(origsecret))
	copy(secret, origsecret)
	move := make([]int, len(origmove))
	copy(move, origmove)
	points := PointsData{
		black: 0,
		white: 0,
	}

	for i := 0; i < len(secret); i++ {
		if secret[i] == move[i] {
			secret[i] = -1
			move[i] = -1
			points.black++
		}
	}

	for _, guess := range move {
		if guess != -1 {
			for secretI, color := range secret {
				if color == guess {
					secret[secretI] = -1
					points.white++
					break
				}
			}
		}
	}
	return &points
}

func (game *GameData) GetPoints() []Points {
	return game.points
}

func (game *GameData) HasWon() bool {
	index, err := findNextIndex(game)
	if err != nil {
		index = len(game.moves)
	}
	index--
	if index < 0 {
		return false
	}
	return game.points[index].GetBlack() == game.moveSize
}

func (game *GameData) HasLost() bool {
	if game.HasWon() {
		return false
	}
	if game.moves[len(game.moves)-1] != nil {
		return true
	}
	return false
}

func findNextIndex(data *GameData) (int, error) {
	for index, element := range data.moves {
		if element == nil {
			return index, nil
		}
	}
	return 0, errors.New("out of moves")
}

func StartGame() Game {
	moveSize := 6
	colorCount := 10
	moveCount := 15
	return &GameData{
		moveSize:   moveSize,
		colorCount: colorCount,
		moves:      make([][]int, moveCount),
		secret:     getSecret(moveSize, colorCount),
		points:     make([]Points, moveCount),
	}
}

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

func getSecret(moveSize int, colors int) []int {
	secret := make([]int, moveSize)
	for index := range secret {
		secret[index] = r1.Intn(colors)
	}
	return secret
}
