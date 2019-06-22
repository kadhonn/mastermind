package main

import (
	"fmt"
	"github.com/kadhonn/mastermind/ai"
)

func main() {
	//run(func() ai.Evaluator { return ai.CompleteRandom }, 10000000, 100000)
	//run(func() ai.Evaluator { return ai.RandomWithSafeGuard }, 100, 10)
	run(prefixes(
		[][]int{
			{1, 1, 1, 2, 2, 2},
			{3, 3, 3, 4, 4, 4},
			{5, 5, 5, 6, 6, 6},
			{7, 7, 7, 8, 8, 8},
			{9, 9, 9, 0, 0, 0},
		},
		ai.RandomWithSafeGuard), 100, 10)
	run(prefixes(
		[][]int{
			{1, 1, 2, 2, 3, 3},
			{4, 4, 5, 5, 6, 6},
			{8, 8, 9, 9, 0, 0},
		},
		ai.RandomWithSafeGuard), 100, 10)
	run(prefixes(
		[][]int{
			{1, 2, 3, 4, 5, 6},
			{7, 8, 9, 0, 5, 6},
		},
		ai.RandomWithSafeGuard), 100, 10)

}

func run(evaluatorCreator ai.EvaluatorCreator, times int, everyNTimes int) {
	statistics := ai.StartEvaluationWithTime(evaluatorCreator, times, everyNTimes)
	fmt.Printf("Won: %d Total: %d Avg Rounds: %2.1f", statistics.Won, statistics.Total, getAvg(statistics.Rounds))
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
