package ai_evo

import (
	"github.com/kadhonn/mastermind/ai"
	"github.com/kadhonn/mastermind/mastermind"
	"log"
	"math/rand"
	"time"
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

func EvoEval(dna DNA) ai.EvaluatorCreator {
	return func() ai.Evaluator {
		return func(game mastermind.Game) []int {
			return doEvoEval(game, dna)
		}
	}
}

func doEvoEval(game mastermind.Game, dna DNA) []int {
	fields := make([]int, getFieldsSize(game.GetColorCount(), game.GetMoveSize()))
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

func getFieldsSize(colorCount int, moveSize int) int {
	return colorCount * moveSize * 2
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

type DNACreationSizes struct {
	colorCount int
	moveSize   int
	fieldsSize int
	dnaSize    int
}

func CreateRandomDNA(colorCount int, moveSize int, size int) DNA {
	sizes := DNACreationSizes{
		colorCount: colorCount,
		moveSize:   moveSize,
		fieldsSize: getFieldsSize(colorCount, moveSize),
		dnaSize:    size,
	}
	nucl := make([]interface{}, size)

	for i := range nucl {
		nucl[i] = createRandomNucl(sizes)
	}

	return DNA{nucl: nucl}
}

var s1 = rand.NewSource(time.Now().UnixNano() + 1245)
var r1 = rand.New(s1)

func createRandomNucl(sizes DNACreationSizes) interface{} {
	switch r1.Intn(3) {
	case 0, 1:
		return createRandomAction(sizes)
	case 2:
		return createRandomColorCompare(sizes)
	}
	log.Fatal("fuck me")
	return nil
}

func createRandomColorCompare(sizes DNACreationSizes) ColorCompare {
	return ColorCompare{
		skip:   r1.Intn(10),
		equals: r1.Intn(2) == 0,
		first:  createRandomColor(sizes),
		second: createRandomColor(sizes),
	}
}

func createRandomAction(sizes DNACreationSizes) Action {
	return Action{
		field: createRandomField(sizes),
		color: createRandomColor(sizes),
	}
}

func createRandomColor(sizes DNACreationSizes) interface{} {
	switch r1.Intn(2) {
	case 0:
		return createRandomField(sizes)
	case 1:
		return createRandomFixedColor(sizes)
	}
	log.Fatal("fuck me")
	return nil
}

func createRandomFixedColor(sizes DNACreationSizes) FixedColor {
	return FixedColor{r1.Intn(sizes.colorCount+1) - 1}
}

func createRandomField(sizes DNACreationSizes) Field {
	return Field{r1.Intn(sizes.fieldsSize)}
}
