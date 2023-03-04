package slots

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wthsths/slots/pkg/random"
)

var testVariation string = `
{
	"slug": "fruits",
	"symbols": [
		{ "payouts": [0, 0, 5, 10, 25], "slug": "cherry" },
		{ "payouts": [0, 0, 5, 10, 25], "slug": "strawberry" },
		{ "payouts": [0, 0, 5, 10, 25], "slug": "lemon" },
		{ "payouts": [0, 0, 5, 10, 25], "slug": "orange" },
		{ "payouts": [0, 0, 10, 25, 50], "slug": "plum" },
		{ "payouts": [0, 0, 10, 25, 50], "slug": "grapes" },
		{ "payouts": [0, 0, 50, 100, 250], "slug": "bell" },
		{ "payouts": [0, 0, 50, 100, 250], "slug": "watermelon" },
		{ "payouts": [0, 0, 100, 250, 500], "slug": "seven" }
	],
	"reels": [
		[0, 1, 2, 3, 4, 5, 6, 7],
		[1, 2, 3, 4, 5, 6, 7, 0],
		[2, 3, 4, 5, 6, 7, 0, 1],
		[3, 4, 5, 6, 7, 0, 1, 2],
		[4, 5, 6, 7, 0, 1, 2, 3]
	],
	"lines": [
		[1, 1, 1, 1, 1],
		[0, 0, 0, 0, 0],
		[2, 2, 2, 2, 2],
		[1, 1, 0, 1, 2],
		[1, 1, 2, 1, 0],
		[1, 0, 1, 2, 1],
		[1, 0, 1, 2, 2],
		[1, 0, 0, 1, 2],
		[1, 2, 1, 0, 1],
		[1, 2, 2, 1, 0],
		[1, 2, 1, 0, 0],
		[0, 1, 2, 1, 0],
		[0, 1, 1, 1, 2],
		[0, 0, 1, 2, 2],
		[0, 0, 1, 2, 1],
		[0, 0, 0, 1, 2],
		[2, 1, 0, 1, 2],
		[2, 1, 1, 1, 0],
		[2, 2, 1, 0, 0],
		[2, 2, 1, 0, 1]
	]
}
`

var randShiftVal = -1

func TestNewSlotMachine(t *testing.T) {
	s, err := getSlotMachine()

	assert.NotNil(t, s)
	assert.NoError(t, err)
}

func TestSpin(t *testing.T) {
	s, _ := getSlotMachine()
	s.Spin()

	expect := MAX_SPIN + randShiftVal
	assert.Equal(t, s.GetSpins(), []int{expect, expect, expect, expect, expect})
}

func TestGetSpins(t *testing.T) {
	s, _ := getSlotMachine()
	assert.Equal(t, s.GetSpins(), []int{0, 0, 0, 0, 0})
}

func TestGetReels(t *testing.T) {
	s, _ := getSlotMachine()
	assert.Equal(t, s.GetReels(), []int{8, 8, 8, 8, 8})
}

func TestGetReelsPositions(t *testing.T) {
	s, _ := getSlotMachine()
	got := s.GetReelsPositions()
	assert.Equal(t, []int{0, 0, 0, 0, 0}, got)

	s.Spin()
	got = s.GetReelsPositions()
	assert.Equal(t, []int{6, 6, 6, 6, 6}, got)
}

func TestGetLineReelPositions(t *testing.T) {
	s, _ := getSlotMachine()
	s.Spin()

	got := s.GetLineReelPositions([]int{0, 0, 0, 0, 0})
	assert.Equal(t, []int{6, 6, 6, 6, 6}, got)

	got = s.GetLineReelPositions([]int{1, 1, 1, 1, 1})
	assert.Equal(t, []int{7, 7, 7, 7, 7}, got)

	got = s.GetLineReelPositions([]int{2, 1, 1, 1, 2})
	assert.Equal(t, []int{0, 7, 7, 7, 0}, got)
}

func TestGetLineSymbolIndexes(t *testing.T) {
	s, _ := getSlotMachine()
	s.Spin()

	got := s.GetLineSymbolIndexes([]int{0, 0, 0, 0, 0})
	assert.Equal(t, []int{6, 7, 0, 1, 2}, got)

	got = s.GetLineSymbolIndexes([]int{1, 1, 1, 1, 1})
	assert.Equal(t, []int{7, 0, 1, 2, 3}, got)
}

func getSlotMachine() (*slotMachine, error) {
	v, _ := NewVariationFromString(testVariation)

	fakeRandom := random.NewFakeRandom(randShiftVal)
	return NewSlotMachine(v, fakeRandom)
}
