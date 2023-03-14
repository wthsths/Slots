package slots

import (
	"encoding/json"
	"errors"
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
	Wild    bool      `json:"wild"`
}

var (
	errReadingConfigFile = errors.New("error reading config file")
	errInvalidJson       = errors.New("invalid json")
)

// NewVariationFromConfig reading file and returns variation struct.
func NewVariationFromConfig(file string) (*Variation, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errReadingConfigFile
	}

	variation := &Variation{}
	err = json.Unmarshal([]byte(b), &variation)
	if err != nil {
		return nil, errInvalidJson
	}

	return variation, nil
}

// NewVariationFromString creating variation struct from string.
func NewVariationFromString(s string) (*Variation, error) {
	variation := &Variation{}
	if err := json.Unmarshal([]byte(s), &variation); err != nil {
		return nil, errInvalidJson
	}

	return variation, nil
}
