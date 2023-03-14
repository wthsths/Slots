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

	for index, value := range g.variation.Lines {
		if index < g.lines {
			lineSymbolIndexes := slotMachine.GetLineSymbolIndexes(value)

			firstSymbolIndex := lineSymbolIndexes[0]
			symbolMatchesCount := 1

			for _, symbolIndex := range lineSymbolIndexes[1:] {
				if firstSymbolIndex == symbolIndex || g.variation.Symbols[symbolIndex].Wild {
					symbolMatchesCount++
				}
			}

			lineWinAmount := g.variation.Symbols[firstSymbolIndex].Payouts[symbolMatchesCount-1]

			if lineWinAmount > 0 {
				g.winLines = append(g.winLines, index)
			}

			g.winAmount += lineWinAmount
		}
	}

	return nil
}

func (g *Game) GetWinAmount() float64 {
	return g.winAmount
}

func (g *Game) GetWinLines() []int {
	return g.winLines
}
