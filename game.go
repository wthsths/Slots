package slots

import (
	"sync"

	"github.com/wthsths/slots/pkg/random"
)

type Game struct {
	sync.RWMutex
	slotMachine *slotMachine
	variation   *Variation
}

func NewGame(variation *Variation, rand random.Random) (*Game, error) {
	slotMachine, err := NewSlotMachine(variation, rand)
	if err != nil {
		return nil, err
	}

	return &Game{
		slotMachine: slotMachine,
		variation:   variation,
	}, nil
}

func (g *Game) Play(bet float64, lines int) (Result, error) {
	g.Lock()
	defer g.Unlock()

	result := Result{}
	result.Bet = bet
	result.Lines = lines
	result.LastSpin = g.slotMachine.GetLastSpins()
	result.Spins = g.slotMachine.Spin()

	for lineIndex, line := range g.variation.Lines {
		if lineIndex >= lines {
			break
		}

		lineWinAmount, symbolMatchesCount := g.calculateLineWinAmount(g.slotMachine, line)
		if lineWinAmount > 0 {
			resultWinLine := ResultWinLines{}
			resultWinLine.LineIndex = lineIndex
			resultWinLine.WinAmount = lineWinAmount
			resultWinLine.SymbolMatchesCount = symbolMatchesCount

			result.WinLines = append(result.WinLines, resultWinLine)
		}
		result.TotalWinAmount += lineWinAmount * bet
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
