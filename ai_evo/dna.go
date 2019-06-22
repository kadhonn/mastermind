package ai_evo

import (
	"github.com/kadhonn/mastermind/ai"
	"github.com/kadhonn/mastermind/mastermind"
	"log"
)

type Field struct {
	number int
}

type FixedColor struct {
	color int
}

type ColorCompare struct {
	skip   int
	equals bool
	first  interface{}
	second interface{}
}

type Action struct {
	field Field
	color interface{}
}

type DNA struct {
	nucl []interface{}
}

func EvoEval(dna DNA) ai.Evaluator {
	return func(game mastermind.Game) []int {
		return doEvoEval(game, dna)
	}
}

func doEvoEval(game mastermind.Game, dna DNA) []int {
	fields := make([]int, game.GetColorCount()*game.GetMoveSize()*2)
	for i := range fields {
		fields[i] = -1
	}

	evalMove(dna, fields)
	for _, move := range game.GetMoves() {
		if move == nil {
			break
		}
		for i := range move {
			fields[i+len(move)] = move[i]
		}
		evalMove(dna, fields)
	}

	result := make([]int, game.GetMoveSize())
	for i := range result {
		if fields[i] == -1 {
			result[i] = 0
		} else {
			result[i] = fields[i]
		}
	}
	return result
}

func evalMove(dna DNA, fields []int) {
	i := 0
	forward := 0
	for i = 0; i < len(dna.nucl); i += forward {
		forward = evalNucl(dna.nucl[i], fields)
	}
}

func evalNucl(nucl interface{}, fields []int) int {
	action, ok := nucl.(Action)
	if ok {
		evalAction(action, fields)
		return 1
	}
	colorCompare, ok := nucl.(ColorCompare)
	if ok {
		return evalColorCompare(colorCompare, fields)
	}
	log.Fatal("unknown value", nucl)
	return -1
}

func evalAction(action Action, fields []int) {
	color := resolveColor(action.color, fields)
	fields[action.field.number] = color
}

func resolveColor(color interface{}, fields []int) int {
	fixedColor, ok := color.(FixedColor)
	if ok {
		return fixedColor.color
	}
	field, ok := color.(Field)
	if ok {
		return fields[field.number]
	}
	log.Fatal("unknown color", color)
	return -1
}

func evalColorCompare(compare ColorCompare, fields []int) int {
	firstColor := resolveColor(compare.first, fields)
	secondColor := resolveColor(compare.second, fields)
	if (compare.equals && firstColor == secondColor) || (!compare.equals && firstColor != secondColor) {
		return compare.skip + 1
	}
	return 1
}
