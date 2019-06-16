package main

import (
	"fmt"
	"mastermind/ai"
	"mastermind/mastermind"
)

func main() {
	//statistics := ai.StartEvaluationWithTime(ai.CompleteRandom,10000000, 100000)
	//statistics := ai.StartEvaluationWithTime(ai.RandomWithSafeGuard, 10, 10)
	statistics := ai.StartEvaluationWithTime(prefixes(
		[][]int{
			{1, 1, 1, 2, 2, 2},
			{3, 3, 3, 4, 4, 4},
			{5, 5, 5, 6, 6, 6},
			{7, 7, 7, 8, 8, 8},
			{9, 9, 9, 0, 0, 0},
		},
		ai.RandomWithSafeGuard), 10, 10)

	fmt.Printf("%+v\n", statistics)
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
