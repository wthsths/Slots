package main

import (
	"log"

	"github.com/wthsths/slots"
	"github.com/wthsths/slots/pkg/random"
)

func main() {
	rand := random.NewSimpleRandom()
	variation, _ := slots.NewVariationFromConfig("../variations/fruits.json")
	game, err := slots.NewGame(variation, rand, 100, 20)
	if err != nil {
		log.Fatal(err)
	}

	result, err := game.Play()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Game Result: %+v\n", result)
}
