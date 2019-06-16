package main

import (
	"fmt"
	"mastermind/ai"
)

func main() {
	statistics := ai.StartEvaluation(ai.CompleteRandom)

	fmt.Printf("%+v\n", statistics)
}
