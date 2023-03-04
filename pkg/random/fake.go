package random

type FakeRandom struct {
	shiftVal int
}

func NewFakeRandom(shiftVal int) FakeRandom {
	return FakeRandom{shiftVal}
}

func (r FakeRandom) Intn(n int) int {
	return n + r.shiftVal
}
