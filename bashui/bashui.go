package main

import (
	"log"
	"mastermind/mastermind"
	"os"
)

var game mastermind.Game

func main() {
	game = mastermind.StartGame()

	printGame()

	for true {
		game.MakeMove([]int{0, 1, 2, 3, 4, 5})
		err := game.MakeMove([]int{4, 5, 6, 7, 8, 9})
		if err != nil {
			log.Fatal(err)
		}

		printGame()
	}
}

func printGame() {
	printHeader()
	for _, move := range game.GetMoves() {
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
