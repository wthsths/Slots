package slots

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wthsths/slots/pkg/random"
)

func TestPlay(t *testing.T) {
	tests := []struct {
		name          string
		variation     *Variation
		randomizer    random.Random
		bet           float64
		lines         int
		wantWinAmount float64
		wantWinLines  []int
		wantErr       error
	}{
		{
			name:          "should win 50 from line 0",
			variation:     getTestVariation(),
			randomizer:    random.NewFakeRandom(0),
			bet:           1,
			lines:         1,
			wantWinAmount: 50,
			wantWinLines:  []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, _ := NewGame(tt.variation, tt.randomizer, tt.bet, tt.lines)
			err := g.Play()
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantWinAmount, g.GetWinAmount())
			assert.Equal(t, tt.wantWinLines, g.GetWinLines())
		})
	}
}

func getTestVariation() *Variation {
	// test variation
	// one	one	 one  one  one
	// two  two  two  two  two
	// thr  thr  thr  thr  thr
	return &Variation{
		Slug: "test_game",
		Symbols: []symbol{
			{Payouts: []float64{0, 0, 0, 0, 20}, Slug: "one"},
			{Payouts: []float64{0, 0, 0, 0, 50}, Slug: "two"},
			{Payouts: []float64{0, 0, 0, 0, 100}, Slug: "thr"},
		},
		Reels: [][]int{
			{0, 1, 2},
			{0, 1, 2},
			{0, 1, 2},
			{0, 1, 2},
			{0, 1, 2},
		},
		Lines: [][]int{
			{1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0},
			{2, 2, 2, 2, 2},
		},
	}
}
