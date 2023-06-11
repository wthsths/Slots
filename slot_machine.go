package slots

import (
	"math"

	"github.com/wthsths/slots/pkg/random"
)

const MIN_SPIN = 1
const MAX_SPIN = 999

type slotMachine struct {
	variation *Variation
	rand      random.Random
	reels     []int
	spins     []int
}

func NewSlotMachine(variation *Variation, rand random.Random) (*slotMachine, error) {
	s := &slotMachine{
		variation: variation,
		rand:      rand,
		reels:     []int{0, 0, 0, 0, 0},
		spins:     []int{0, 0, 0, 0, 0},
	}

	for i := range variation.Reels {
		s.reels[i] = len(variation.Reels[i])
	}

	return s, nil
}

func (s *slotMachine) Spin() {
	for i := range s.spins {
		s.spins[i] = s.rand.Intn(MAX_SPIN-MIN_SPIN) + MIN_SPIN
	}
}

func (s *slotMachine) GetSpins() []int {
	return s.spins
}

func (s *slotMachine) GetReels() []int {
	return s.reels
}

func (s *slotMachine) GetReelsPositions() []int {
	positions := make([]int, len(s.reels))

	for i := range s.reels {
		v := s.reels[i]
		spin := s.spins[i]

		if v != spin && float64(v) > 0 && float64(spin) > 0 {
			positions[i] = int(math.Max(float64(v), float64(spin))) % int(math.Min(float64(v), float64(spin)))
		} else {
			positions[i] = 0
		}
	}

	return positions
}

func (s *slotMachine) GetLineReelPositions(line []int) []int {
	positions := make([]int, len(s.reels))

	for i, reelIndex := range line {
		reelSymbolsCount := len(s.variation.Reels[i])
		returnValue := s.GetReelsPositions()[i] + reelIndex

		if returnValue > reelSymbolsCount-1 {
			positions[i] = returnValue - reelSymbolsCount
		} else {
			positions[i] = returnValue
		}
	}

	return positions
}

func (s *slotMachine) GetLineSymbolIndexes(line []int) []int {
	lineReelPositions := s.GetLineReelPositions(line)
	symbolIndexes := make([]int, len(lineReelPositions))

	for i, pos := range lineReelPositions {
		symbolIndexes[i] = s.variation.Reels[i][pos]
	}

	return symbolIndexes
}
