package slots

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wthsths/slots/pkg/random"
)

func TestPlay(t *testing.T) {
	var variation string = `
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

	v, _ := NewVariationFromString(variation)
	r := random.NewFakeRandom(-1)

	g, err := NewGame(v, r, 1, 20)
	assert.NoError(t, err)
	assert.NotNil(t, g)

	err = g.Play()
	assert.NoError(t, err)
	assert.Equal(t, float64(5), g.GetWinAmount())
	assert.Equal(t, []int{16}, g.GetWinLines())
}
