package ai

import (
	"log"
	"mastermind/mastermind"
	"math/rand"
	"time"
)

type evaluator func(mastermind.Game) []int

type gameCreator func() mastermind.Game

type Statistics struct {
	won   int
	total int
}

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

func StartEvaluationWithTime(eval evaluator, n_games int, every_n_games int) Statistics {
	return startEvaluationWithTime(eval, mastermind.StartGame, n_games, every_n_games)
}

func startEvaluationWithCreator(eval evaluator, creator gameCreator) Statistics {
	return startEvaluationWithTime(eval, creator, 100, 10)
}

func startEvaluationWithTime(eval evaluator, creator gameCreator, n_games int, every_n_games int) Statistics {
	statistics := Statistics{0, 0}

	for i := 0; i < n_games; i++ {
		game := creator()

		evalOneGame(eval, game)
		PrintGame(game)

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

	return layoutBlacks(0, oldMove, &mastermind.PointsData{Black: p.GetBlack(), White: p.GetWhite()}, move)
}

//LAYOUT: 0 = clear, 1 = black, 2 = white
func layoutBlacks(startIndex int, oldMove []int, p *mastermind.PointsData, move []int) bool {
	if p.Black > 0 {
		p.Black--
		for i := startIndex; i < len(oldMove); i++ {
			if oldMove[i] == move[i] {
				guess := move[i]
				move[i] = -1
				oldMove[i] = -1

				ret := layoutBlacks(i+1, oldMove, p, move)
				if ret {
					return true
				}

				move[i] = guess
				oldMove[i] = guess
			}
		}
		p.Black++
		return false
	} else {
		return layoutWhites(0, oldMove, p, move)
	}
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