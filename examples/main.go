package main

import (
	"log"

	"github.com/wthsths/slots"
	"github.com/wthsths/slots/pkg/random"
)

func main() {
	rand := random.NewSimpleRandom()
	variation, err := slots.NewVariationFromConfig("./fruits.json")
	if err != nil {
		log.Fatal(err)
	}

	game, err := slots.NewGame(variation, rand)
	if err != nil {
		log.Fatal(err)
	}

	result, err := game.Play(100, 20)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Game Result: %+v\n", result)
}
