package main

import (
	"flag"
	"fmt"
	"github.com/kadhonn/mastermind/ai"
	"github.com/kadhonn/mastermind/ai_evo"
	"github.com/kadhonn/mastermind/mastermind"
	"os"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	//run(func() ai.Evaluator { return ai.CompleteRandom }, 10, 100000)
	//run(func() ai.Evaluator { return ai.RandomWithSafeGuard }, 100, 10)
	//flag.Parse()
	//if *cpuprofile != "" {
	//	f, err := os.Create(*cpuprofile)
	//	if err != nil {
	//		log.Fatal("could not create CPU profile: ", err)
	//	}
	//	defer f.Close()
	//	if err := pprof.StartCPUProfile(f); err != nil {
	//		log.Fatal("could not start CPU profile: ", err)
	//	}
	//	defer pprof.StopCPUProfile()
	//}

	runEvo()
	// ... rest of the program ...

	//if *memprofile != "" {
	//	f, err := os.Create(*memprofile)
	//	if err != nil {
	//		log.Fatal("could not create memory profile: ", err)
	//	}
	//	defer f.Close()
	//	runtime.GC() // get up-to-date statistics
	//	if err := pprof.WriteHeapProfile(f); err != nil {
	//		log.Fatal("could not write memory profile: ", err)
	//	}
	//}
}

func runEvo() {
	i := 0
	for true {
		dna := ai_evo.CreateRandomDNA(10, 6, 1000)
		statistics := ai.StartEvaluationWithTime(ai_evo.EvoEval(dna), 1, 999999)

		if statistics.Won >= 0 {
			for j, game := range statistics.Games {
				if game.HasWon() || j < 10 {
					ai.PrintGame(game)
				}
			}
			ai.P(fmt.Sprintf("Won: %d Total: %d Avg Rounds: %2.1f\n", statistics.Won, statistics.Total, getAvg(statistics.Games)))
			os.Exit(0)
		}
		i++
		if i%1 == 0 {
			ai.P(fmt.Sprintf("%d tries\n", i))
		}
	}
}

func run(evaluatorCreator ai.EvaluatorCreator, times int, everyNTimes int) ai.Statistics {
	statistics := ai.StartEvaluationWithTime(evaluatorCreator, times, everyNTimes)
	ai.P(fmt.Sprintf("Won: %d Total: %d Avg Rounds: %2.1f\n", statistics.Won, statistics.Total, getAvg(statistics.Games)))
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
