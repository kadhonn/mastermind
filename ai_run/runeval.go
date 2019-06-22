package main

import (
	"fmt"
	"github.com/kadhonn/mastermind/ai"
	"github.com/kadhonn/mastermind/ai_evo"
	"github.com/kadhonn/mastermind/mastermind"
)

func main() {
	//run(func() ai.Evaluator { return ai.CompleteRandom }, 10, 100000)
	//run(func() ai.Evaluator { return ai.RandomWithSafeGuard }, 100, 10)

	//for true {
	run(ai_evo.EvoEval(ai_evo.CreateRandomDNA(6, 10, 1000)), 100, 1000)
	//}

}

func run(evaluatorCreator ai.EvaluatorCreator, times int, everyNTimes int) {
	statistics := ai.StartEvaluationWithTime(evaluatorCreator, times, everyNTimes)
	//if statistics.Won != 0 {
	fmt.Printf("Won: %d Total: %d Avg Rounds: %2.1f\n", statistics.Won, statistics.Total, getAvg(statistics.Rounds))
	//log.Fatal("woaaah")
	//}
}

func getAvg(rounds []int) float64 {
	sum := 0.0
	for _, round := range rounds {
		sum += float64(round)
	}
	return sum / float64(len(rounds))
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
