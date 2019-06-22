package main

import (
	"bufio"
	"errors"
	"github.com/kadhonn/mastermind/ai"
	"github.com/kadhonn/mastermind/mastermind"
	"log"
	"os"
	"strconv"
)

var game mastermind.Game
var in *bufio.Reader

func main() {
	in = bufio.NewReader(os.Stdin)
	game = mastermind.StartGame()

	ai.PrintGame(game)

	for true {
		move, err := readMove()
		if err == nil {
			err = game.MakeMove(move)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Print(err)
		}

		ai.PrintGame(game)

		if game.HasWon() {
			ai.Red()
			ai.P("YOU WON!")
			os.Exit(0)
		}
		if game.HasLost() {
			ai.Red()
			ai.P("YOU LOOOSE!")
			os.Exit(0)
		}
	}
}

func readMove() ([]int, error) {
	line, err := in.ReadString('\n')
	if err != nil {
		log.Fatal("could not read input", err)
	}
	if len(line)-1 != game.GetMoveSize() {
		return nil, errors.New("line is not same size as word size")
	}
	move := make([]int, game.GetMoveSize())
	for i := 0; i < game.GetMoveSize(); i++ {
		move[i], err = strconv.Atoi(line[i : i+1])
		if err != nil {
			return nil, err
		}
	}
	return move, nil
}
