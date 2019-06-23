package ai_evo

import (
	"bytes"
	"encoding/gob"
	"github.com/kadhonn/mastermind/ai"
	"github.com/kadhonn/mastermind/mastermind"
	"log"
	"math/rand"
	"time"
)

type Field struct {
	Number int
}

type FixedColor struct {
	Color int
}

type ColorCompare struct {
	Skip   int
	Equals bool
	First  interface{}
	Second interface{}
}

type Action struct {
	Field *Field
	Color interface{}
}

type PointsCompare struct {
	Skip   int
	Mode   int
	Blacks bool
	Count  int
}

type DNA struct {
	Nucl []interface{}
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

	evalMove(dna, fields, &mastermind.PointsData{0, 0})
	for i, move := range game.GetMoves() {
		if move == nil {
			break
		}
		for i := range move {
			fields[i+len(move)] = move[i]
		}
		evalMove(dna, fields, game.GetPoints()[i])
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

func evalMove(dna DNA, fields []int, points mastermind.Points) {
	i := 0
	forward := 0
	for i = 0; i < len(dna.Nucl); i += forward {
		forward = evalNucl(dna.Nucl[i], fields, points)
	}
}

func evalNucl(nucl interface{}, fields []int, points mastermind.Points) int {
	action, ok := nucl.(*Action)
	if ok {
		evalAction(action, fields)
		return 1
	}
	colorCompare, ok := nucl.(*ColorCompare)
	if ok {
		return evalColorCompare(colorCompare, fields)
	}
	pointsCompare, ok := nucl.(*PointsCompare)
	if ok {
		return evalPointsCompare(pointsCompare, points)
	}
	log.Fatal("unknown value", nucl)
	return -1
}

func evalPointsCompare(compare *PointsCompare, points mastermind.Points) int {
	var count int
	if compare.Blacks {
		count = points.GetBlack()
	} else {
		count = points.GetWhite()
	}

	var machted bool
	switch compare.Mode {
	case 0:
		machted = count == compare.Count
	case 1:
		machted = count != compare.Count
	case 2:
		machted = count < compare.Count
	case 3:
		machted = count > compare.Count
	default:
		log.Fatal("wrong Mode", compare.Mode)
	}
	if machted {
		return compare.Skip + 1
	} else {
		return 1
	}
}

func evalAction(action *Action, fields []int) {
	color := resolveColor(action.Color, fields)
	fields[action.Field.Number] = color
}

func resolveColor(color interface{}, fields []int) int {
	fixedColor, ok := color.(*FixedColor)
	if ok {
		return fixedColor.Color
	}
	field, ok := color.(*Field)
	if ok {
		return fields[field.Number]
	}
	log.Fatal("unknown color", color)
	return -1
}

func evalColorCompare(compare *ColorCompare, fields []int) int {
	firstColor := resolveColor(compare.First, fields)
	secondColor := resolveColor(compare.Second, fields)
	if (compare.Equals && firstColor == secondColor) || (!compare.Equals && firstColor != secondColor) {
		return compare.Skip + 1
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

	return DNA{Nucl: nucl}
}

var s1 = rand.NewSource(time.Now().UnixNano() + 1245)
var r1 = rand.New(s1)

func createRandomNucl(sizes DNACreationSizes) interface{} {
	switch r1.Intn(5) {
	case 0, 1, 2:
		return createRandomAction(sizes)
	case 3:
		return createRandomColorCompare(sizes)
	case 4:
		return createRandomPointsCompare(sizes)
	}
	log.Fatal("fuck me")
	return nil
}

func createRandomPointsCompare(sizes DNACreationSizes) *PointsCompare {
	return &PointsCompare{
		Skip:   r1.Intn(50),
		Mode:   r1.Intn(4),
		Blacks: r1.Intn(2) == 0,
		Count:  r1.Intn(sizes.colorCount),
	}
}

func createRandomColorCompare(sizes DNACreationSizes) *ColorCompare {
	return &ColorCompare{
		Skip:   r1.Intn(50),
		Equals: r1.Intn(2) == 0,
		First:  createRandomColor(sizes),
		Second: createRandomColor(sizes),
	}
}

func createRandomAction(sizes DNACreationSizes) *Action {
	return &Action{
		Field: createRandomField(sizes),
		Color: createRandomColor(sizes),
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

func createRandomFixedColor(sizes DNACreationSizes) *FixedColor {
	return &FixedColor{r1.Intn(sizes.colorCount+1) - 1}
}

func createRandomField(sizes DNACreationSizes) *Field {
	return &Field{r1.Intn(sizes.fieldsSize)}
}

func Save(dna DNA) *bytes.Buffer {
	gob.Register(&Field{})
	gob.Register(&ColorCompare{})
	gob.Register(&PointsCompare{})
	gob.Register(&Action{})
	gob.Register(&FixedColor{})

	buffer := &bytes.Buffer{}
	enc := gob.NewEncoder(buffer)
	err := enc.Encode(dna)
	if err != nil {
		log.Fatal("encode:", err)
	}
	return buffer
}

func Load(buffer *bytes.Buffer) DNA {
	gob.Register(&Field{})
	gob.Register(&ColorCompare{})
	gob.Register(&PointsCompare{})
	gob.Register(&Action{})
	gob.Register(&FixedColor{})

	dec := gob.NewDecoder(buffer)
	var dna DNA
	err := dec.Decode(&dna)
	if err != nil {
		log.Fatal("decode:", err)
	}
	return dna
}
