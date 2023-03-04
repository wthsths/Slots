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

	if err := game.Play(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Win Amount: %v\n", game.GetWinAmount())
	log.Printf("Win Lines: %v\n", game.GetWinLines())
}
