package main

import (
	"bufio"
	"errors"
	"log"
	"mastermind/mastermind"
	"os"
	"strconv"
)

var game mastermind.Game
var in *bufio.Reader

func main() {
	in = bufio.NewReader(os.Stdin)
	game = mastermind.StartGame()

	printGame()

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

		printGame()
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

func printGame() {
	printHeader()
	moves := game.GetMoves()
	for i := len(moves) - 1; i >= 0; i-- {
		move := moves[i]
		if move != nil {
			for _, guess := range move {
				p(" ")
				color(guess)
				p("*")
			}
			p("\n")
		}
	}
	printFooter()
}

func printFooter() {
	p("\n\n")
}

func printHeader() {
	white()
	for i := 0; i < game.GetMoveSize(); i++ {
		p(" *")
	}
	p("\n")
	for i := 0; i < game.GetMoveSize(); i++ {
		p("--")
	}
	p("-")
	p("\n")
}

func color(guess int) {
	switch guess {
	case 0:
		red()
	case 1:
		green()
	case 2:
		yellow()
	case 3:
		blue()
	case 4:
		magenta()
	case 5:
		cyan()
	case 6:
		lightgray()
	case 7:
		darkgray()
	case 8:
		black()
	case 9:
		white()
	}
}

func green() {
	p("\u001b[32m")
}

func yellow() {
	p("\u001b[33m")
}

func blue() {
	p("\u001b[34m")
}

func magenta() {
	p("\u001b[35m")
}

func cyan() {
	p("\u001b[36m")
}

func lightgray() {
	p("\u001b[37m")
}

func darkgray() {
	p("\u001b[90m")
}

func black() {
	p("\u001b[30m")
}

func white() {
	p("\u001b[39m")
}

func red() {
	p("\u001b[31m")
}

func p(s string) {
	os.Stdout.WriteString(s)
}
