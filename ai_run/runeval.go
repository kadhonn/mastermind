package main

import (
	"fmt"
	"mastermind/ai"
)

func main() {
	//statistics := ai.StartEvaluationWithTime(ai.CompleteRandom,10000000, 100000)
	statistics := ai.StartEvaluationWithTime(ai.RandomWithSafeGuard, 100, 10)

	fmt.Printf("%+v\n", statistics)
}
