package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type SimpleRandom struct {
}

func NewSimpleRandom() SimpleRandom {
	return SimpleRandom{}
}

func (r SimpleRandom) Intn(n int) int {
	return rand.Intn(n)
}
