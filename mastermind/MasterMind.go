package mastermind

import (
	"errors"
	"fmt"
)

type Game interface {
	GetMoveSize() int
	GetColorCount() int
	GetMoves() [][]int
	GetSecret() []int
	MakeMove(move []int) error
}

type GameData struct {
	moveSize   int
	colorCount int
	moves      [][]int
	secret     []int
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

	return nil
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
	return &GameData{
		moveSize:   6,
		colorCount: 10,
		moves:      make([][]int, 15),
	}
}
