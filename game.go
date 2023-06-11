package slots

import (
	"github.com/wthsths/slots/pkg/random"
)

type Game struct {
	variation *Variation
	rand      random.Random
	bet       float64
	lines     int
}

func NewGame(variation *Variation, rand random.Random, bet float64, lines int) (*Game, error) {
	return &Game{
		variation: variation,
		rand:      rand,
		bet:       bet,
		lines:     lines,
	}, nil
}

func (g *Game) Play() (Result, error) {
	result := Result{}

	slotMachine, err := NewSlotMachine(g.variation, g.rand)
	if err != nil {
		return result, err
	}

	result.Spins = slotMachine.Spin()

	for lineIndex, line := range g.variation.Lines {
		if lineIndex >= g.lines {
			break
		}

		lineWinAmount, symbolMatchesCount := g.calculateLineWinAmount(slotMachine, line)
		if lineWinAmount > 0 {
			resultWinLine := ResultWinLines{}
			resultWinLine.LineIndex = lineIndex
			resultWinLine.WinAmount = lineWinAmount
			resultWinLine.SymbolMatchesCount = symbolMatchesCount

			result.WinLines = append(result.WinLines, resultWinLine)
		}
		result.TotalWinAmount += lineWinAmount
	}

	return result, nil
}

func (g *Game) calculateLineWinAmount(slotMachine *slotMachine, line []int) (float64, int) {
	lineSymbolIndexes := slotMachine.GetLineSymbolIndexes(line)
	firstSymbolIndex := lineSymbolIndexes[0]
	symbolMatchesCount := g.countSymbolMatches(lineSymbolIndexes)

	return g.variation.Symbols[firstSymbolIndex].Payouts[symbolMatchesCount-1], symbolMatchesCount
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
