package ai

import (
	"github.com/kadhonn/mastermind/mastermind"
	"os"
	"strconv"
)

func PrintGame(game mastermind.Game) {
	printHeader(game)
	moves := game.GetMoves()
	for i := len(moves) - 1; i >= 0; i-- {
		move := moves[i]
		if move != nil {
			printMove(move)
			White()
			P(" | ")
			PrintBlacks(game.GetPoints()[i].GetBlack())
			PrintWhites(game.GetPoints()[i].GetWhite())
			P("\n")
		}
	}
	printFooter(game)
}

func printFooter(game mastermind.Game) {
	P("\n\n")
}

func printHeader(game mastermind.Game) {
	if game.HasLost() || game.HasWon() {
		printMove(game.GetSecret())
	} else {
		White()
		for i := 0; i < game.GetMoveSize(); i++ {
			P(" *")
		}
	}
	White()
	P("\n")
	for i := 0; i < game.GetMoveSize(); i++ {
		P("--")
	}
	P("-")
	P("\n")
}

func printMove(move []int) {
	for _, guess := range move {
		P(" ")
		printGuess(guess)
	}
}

func printGuess(guess int) {
	color(guess)
	P(strconv.Itoa(guess))
}

func PrintWhites(whites int) {
	White()
	for j := 0; j < whites; j++ {
		P("*")
	}
}

func PrintBlacks(blacks int) {
	Red()
	for j := 0; j < blacks; j++ {
		P("*")
	}
}

func color(guess int) {
	switch guess {
	case 0:
		Red()
	case 1:
		Green()
	case 2:
		Yellow()
	case 3:
		Blue()
	case 4:
		Magenta()
	case 5:
		Cyan()
	case 6:
		Lightgray()
	case 7:
		Darkgray()
	case 8:
		Black()
	case 9:
		White()
	}
}

func Green() {
	P("\u001b[32m")
}

func Yellow() {
	P("\u001b[33m")
}

func Blue() {
	P("\u001b[34m")
}

func Magenta() {
	P("\u001b[35m")
}

func Cyan() {
	P("\u001b[36m")
}

func Lightgray() {
	P("\u001b[37m")
}

func Darkgray() {
	P("\u001b[90m")
}

func Black() {
	P("\u001b[30m")
}

func White() {
	P("\u001b[39m")
}

func Red() {
	P("\u001b[31m")
}

func P(s string) {
	_, _ = os.Stdout.WriteString(s)
}
