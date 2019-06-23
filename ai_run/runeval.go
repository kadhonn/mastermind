package main

import (
	"fmt"
	"github.com/kadhonn/mastermind/ai"
	"github.com/kadhonn/mastermind/ai_evo"
	"github.com/kadhonn/mastermind/mastermind"
	"os"
)

func main() {
	//run(func() ai.Evaluator { return ai.CompleteRandom }, 10, 100000)
	//run(func() ai.Evaluator { return ai.RandomWithSafeGuard }, 100, 10)

	i := 0
	for true {
		dna := ai_evo.CreateRandomDNA(6, 10, 10)
		statistics := run(ai_evo.EvoEval(dna), 10000, 10000)
		if statistics.Won >= 1 {
			for _, game := range statistics.Games {
				ai.PrintGame(game)
			}
			os.Exit(0)
		}
		i++
		if i%100 == 0 {
			ai.P(fmt.Sprintf("%d tries\n", i))
		}
	}

}

func run(evaluatorCreator ai.EvaluatorCreator, times int, everyNTimes int) ai.Statistics {
	statistics := ai.StartEvaluationWithTime(evaluatorCreator, times, everyNTimes)
	if statistics.Won != 0 {
		ai.P(fmt.Sprintf("Won: %d Total: %d Avg Rounds: %2.1f\n", statistics.Won, statistics.Total, getAvg(statistics.Games)))
	}
	return statistics
}

func getAvg(games []mastermind.Game) float64 {
	sum := 0.0
	for _, game := range games {
		sum += float64(getMoveCount(game.GetMoves()))
	}
	return sum / float64(len(games))
}

func getMoveCount(moves [][]int) int {
	for i, move := range moves {
		if move == nil {
			return i
		}
	}
	return len(moves)
}

func prefixes(moves [][]int, evaluator ai.Evaluator) ai.EvaluatorCreator {
	return func() ai.Evaluator {
		i := 0
		return func(game mastermind.Game) []int {
			if i >= len(moves) {
				return evaluator(game)
			} else {
				i++
				return moves[i-1]
			}
		}
	}
}
