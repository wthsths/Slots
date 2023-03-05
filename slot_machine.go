package slots

import (
	"math"

	"github.com/samber/lo"
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

	s.reels = lo.Map(variation.Reels, func(r []int, _ int) int {
		return len(r)
	})

	return s, nil
}

func (s *slotMachine) Spin() {
	s.spins = lo.Map(s.spins, func(v int, _ int) int {
		return s.rand.Intn(MAX_SPIN-MIN_SPIN) + MIN_SPIN
	})
}

func (s *slotMachine) GetSpins() []int {
	return s.spins
}

func (s *slotMachine) GetReels() []int {
	return s.reels
}

func (s *slotMachine) GetReelsPositions() []int {
	return lo.Map(
		lo.Zip2(s.reels, s.spins),
		func(v lo.Tuple2[int, int], _ int) int {
			if v.A != v.B && math.Min(float64(v.A), float64(v.B)) > 0 {
				return int(math.Max(float64(v.A), float64(v.B))) %
					int(math.Min(float64(v.A), float64(v.B)))
			}

			return 0
		})
}

func (s *slotMachine) GetLineReelPositions(line []int) []int {
	return lo.Map(
		lo.Zip2(
			s.GetReelsPositions(),
			line,
		), func(v lo.Tuple2[int, int], index int) int {
			reelSymbolsCount := len(s.variation.Reels[index])
			returnValue := v.A + v.B

			if returnValue > reelSymbolsCount-1 {
				return returnValue - reelSymbolsCount
			}

			return returnValue
		})
}

func (s *slotMachine) GetLineSymbolIndexes(line []int) []int {
	lineReelPositions := s.GetLineReelPositions(line)

	return lo.Map(
		lineReelPositions, func(v int, index int) int {
			return s.variation.Reels[index][v]
		})
}
