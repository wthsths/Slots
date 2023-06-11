package slots

import (
	"github.com/wthsths/slots/pkg/random"
)

type Game struct {
	variation *Variation
	rand      random.Random
	bet       float64
	lines     int
	winAmount float64
	winLines  []int
}

func NewGame(variation *Variation, rand random.Random, bet float64, lines int) (*Game, error) {
	return &Game{
		variation: variation,
		rand:      rand,
		bet:       bet,
		lines:     lines,
	}, nil
}

func (g *Game) Play() error {
	slotMachine, _ := NewSlotMachine(g.variation, g.rand)
	slotMachine.Spin()

	for lineIndex, line := range g.variation.Lines {
		if lineIndex >= g.lines {
			break
		}

		lineWinAmount := g.calculateLineWinAmount(slotMachine, line)
		if lineWinAmount > 0 {
			g.winLines = append(g.winLines, lineIndex)
		}
		g.winAmount += lineWinAmount
	}

	return nil
}

func (g *Game) calculateLineWinAmount(slotMachine *slotMachine, line []int) float64 {
	lineSymbolIndexes := slotMachine.GetLineSymbolIndexes(line)
	firstSymbolIndex := lineSymbolIndexes[0]
	symbolMatchesCount := g.countSymbolMatches(lineSymbolIndexes)

	return g.variation.Symbols[firstSymbolIndex].Payouts[symbolMatchesCount-1]
}

func (g *Game) countSymbolMatches(lineSymbolIndexes []int) int {
	firstSymbolIndex := lineSymbolIndexes[0]
	symbolMatchesCount := 1

	for _, symbolIndex := range lineSymbolIndexes[1:] {
		if g.isMatchingSymbol(firstSymbolIndex, symbolIndex) {
			symbolMatchesCount++
		}
	}

	return symbolMatchesCount
}

func (g *Game) isMatchingSymbol(firstSymbolIndex, symbolIndex int) bool {
	return firstSymbolIndex == symbolIndex || g.variation.Symbols[symbolIndex].Wild
}

func (g *Game) GetWinAmount() float64 {
	return g.winAmount
}

func (g *Game) GetWinLines() []int {
	return g.winLines
}
