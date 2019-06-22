package ai

import (
	"github.com/kadhonn/mastermind/mastermind"
	"log"
	"math/rand"
	"time"
)

type EvaluatorCreator func() Evaluator
type Evaluator func(mastermind.Game) []int

type GameCreator func() mastermind.Game

type Statistics struct {
	Won    int
	Total  int
	Rounds []int
}

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

func StartEvaluationWithTime(eval EvaluatorCreator, n_games int, every_n_games int) Statistics {
	return startEvaluationWithTime(eval, mastermind.StartGame, n_games, every_n_games)
}

func startEvaluationWithCreator(eval EvaluatorCreator, creator GameCreator) Statistics {
	return startEvaluationWithTime(eval, creator, 100, 10)
}

func startEvaluationWithTime(evaluatorCreator EvaluatorCreator, creator GameCreator, n_games int, every_n_games int) Statistics {
	statistics := Statistics{Won: 0, Total: 0, Rounds: make([]int, n_games)}

	for i := 0; i < n_games; i++ {
		game := creator()

		evalOneGame(evaluatorCreator(), game)

		if game.HasWon() {
			statistics.Won++
			PrintGame(game)
		} else {
			PrintGame(game)
		}
		statistics.Total++
		statistics.Rounds[i] = getRounds(game)

		if i%every_n_games == 0 {
			log.Printf("%d/%d = %3.2f%%", i, n_games, float32(i)/float32(n_games)*100.0)
		}
	}

	return statistics
}

func getRounds(game mastermind.Game) int {
	for i, move := range game.GetMoves() {
		if move == nil {
			return i
		}
	}
	return len(game.GetMoves())
}

func evalOneGame(e Evaluator, game mastermind.Game) {
	for !game.HasWon() && !game.HasLost() {
		err := game.MakeMove(e(game))
		if err != nil {
			log.Fatal("invalid move from creator", err)
		}
	}
}

func isCompatible(game mastermind.Game, move []int) bool {
	if len(move) != game.GetMoveSize() {
		log.Fatal("moves are not same size!")
	}
	for i, oldMove := range game.GetMoves() {
		if oldMove != nil {
			if !isGuessCompatible(oldMove, game.GetPoints()[i], move) {
				return false
			}
		}
	}
	return true
}

func isGuessCompatible(origOldMove []int, p mastermind.Points, origMove []int) bool {
	oldMove := make([]int, len(origOldMove))
	copy(oldMove, origOldMove)
	move := make([]int, len(origMove))
	copy(move, origMove)
	if len(oldMove) != len(move) {
		log.Fatal("moves are not same size!")
	}
	//TODO REMOVE
	layout := make([]int, len(oldMove))
	for index := range layout {
		layout[index] = 0
	}

	return layoutBlacks(oldMove, &mastermind.PointsData{Black: p.GetBlack(), White: p.GetWhite()}, move)
}

func layoutBlacks(oldMove []int, p *mastermind.PointsData, move []int) bool {
	blackCount := 0
	for i := 0; i < len(oldMove); i++ {
		if oldMove[i] == move[i] {
			move[i] = -1
			oldMove[i] = -1
			blackCount++
		}
	}
	if p.Black != blackCount {
		return false
	}
	return layoutWhites(0, oldMove, p, move)
}

func layoutWhites(startIndex int, oldMove []int, p *mastermind.PointsData, move []int) bool {
	if p.White > 0 {
		p.White--
		for i := startIndex; i < len(oldMove); i++ {
			for j := range move {
				if move[j] != oldMove[j] && oldMove[i] != -1 && move[j] != -1 && oldMove[i] == move[j] {
					guess := move[j]
					oldGuess := oldMove[i]
					move[j] = -1
					oldMove[i] = -1

					ret := layoutWhites(i+1, oldMove, p, move)
					if ret {
						return true
					}

					move[j] = guess
					oldMove[i] = oldGuess
				}
			}
		}
		p.White++
		return false
	} else {
		return layoutClears(oldMove, p, move)
	}
}

func layoutClears(oldMove []int, p *mastermind.PointsData, move []int) bool {
	for _, guess := range move {
		if guess != -1 {
			for _, oldGuess := range oldMove {
				if guess == oldGuess {
					return false
				}
			}
		}
	}
	return true
}
