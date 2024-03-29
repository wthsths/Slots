package slots

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wthsths/slots/pkg/random"
)

func TestPlay(t *testing.T) {
	tests := []struct {
		name       string
		variation  *Variation
		randomizer random.Random
		bet        float64
		lines      int
		want       Result
		wantErr    error
	}{
		{
			name:       "should win 50 from line 0",
			variation:  getTestVariation(),
			randomizer: random.NewFakeRandom(0),
			bet:        1,
			lines:      1,
			want: Result{
				Bet:      1,
				Lines:    1,
				LastSpin: []int{0, 0, 0, 0, 0},
				Spins:    []int{999, 999, 999, 999, 999},
				WinLines: []ResultWinLines{
					{
						LineIndex:          0,
						WinAmount:          50,
						SymbolMatchesCount: 5,
					},
				},
				TotalWinAmount: 50,
			},
		},
		{
			name: "should win 50 with wild from line 0",
			variation: &Variation{
				Symbols: []symbol{
					{Payouts: []float64{0, 0, 0, 0, 0}},
					{Payouts: []float64{0, 0, 0, 0, 50}},
					{Payouts: []float64{0, 0, 0, 0, 0}},
					{Payouts: []float64{0, 0, 0, 0, 0}, Wild: true},
				},
				Reels: [][]int{
					{0, 1, 2},
					{0, 3, 2},
					{0, 3, 2},
					{0, 3, 2},
					{0, 3, 2},
				},
				Lines: [][]int{
					{1, 1, 1, 1, 1},
				},
			},
			randomizer: random.NewFakeRandom(0),
			bet:        1,
			lines:      1,
			want: Result{
				Bet:      1,
				Lines:    1,
				LastSpin: []int{0, 0, 0, 0, 0},
				Spins:    []int{999, 999, 999, 999, 999},
				WinLines: []ResultWinLines{
					{
						LineIndex:          0,
						WinAmount:          50,
						SymbolMatchesCount: 5,
					},
				},
				TotalWinAmount: 50,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, _ := NewGame(tt.variation, tt.randomizer)
			got, gotErr := g.Play(tt.bet, tt.lines)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, gotErr)
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
