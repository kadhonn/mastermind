package ai

import (
	"log"
	"mastermind/mastermind"
)

type evaluator func(mastermind.Game) []int

type gameCreator func() mastermind.Game

type Statistics struct {
	won   int
	total int
}

func StartEvaluation(eval evaluator) Statistics {
	return startEvaluation(eval, mastermind.StartGame)
}

func startEvaluation(eval evaluator, creator gameCreator) Statistics {
	statistics := Statistics{0, 0}

	n_games := 10000000
	every_n_games := 100000

	for i := 0; i < n_games; i++ {
		game := creator()

		evalOneGame(eval, game)
		//PrintGame(game)

		if game.HasWon() {
			statistics.won++
		}
		statistics.total++

		if i%every_n_games == 0 {
			log.Printf("%d/%d = %3.2f%%", i, n_games, float32(i)/float32(n_games)*100.0)
		}
	}

	return statistics
}

func evalOneGame(e evaluator, game mastermind.Game) {
	for !game.HasWon() && !game.HasLost() {
		err := game.MakeMove(e(game))
		if err != nil {
			log.Fatal("invalid move from creator", err)
		}
	}
}
