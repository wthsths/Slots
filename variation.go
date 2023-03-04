package slots

import (
	"encoding/json"
	"io/ioutil"
)

type Variation struct {
	Slug    string   `json:"slug"`
	Symbols []symbol `json:"symbols"`
	Reels   [][]int  `json:"reels"`
	Lines   [][]int  `json:"lines"`
}

type symbol struct {
	Slug    string    `json:"slug"`
	Payouts []float64 `json:"payouts"`
}

func NewVariationFromConfig(file string) (*Variation, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	variation := &Variation{}
	err = json.Unmarshal([]byte(b), &variation)
	if err != nil {
		return nil, err
	}

	return variation, nil
}

func NewVariationFromString(s string) (*Variation, error) {
	variation := &Variation{}
	_ = json.Unmarshal([]byte(s), &variation)

	return variation, nil
}
