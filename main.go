package main

import (
	"fmt"
	pc "github.com/daman1807/deck"
)

func main() {
	deck := pc.NewDeck("standard")

	// add two more decks to the main deck
	for i := 0; i < 2; i++ {
		deck.Add(pc.NewDeck("standard"))
	}

	// shuffle the deck and print first 10 members
	deck.Shuffle()

	dealer := Hand{Name: "Dealer"}
	player := Hand{Name: "Player"}

	var card pc.Card

	// Initialize hands
	for i := 0; i < 4; i += 2 {
		card = deck.Draw()
		player.Add(card)

		card = deck.Draw()
		dealer.Add(card)
	}

	fmt.Println("Initial Hands")

	// Let the game begin!
	var input string
	for input != "s"{

		fmt.Println("Player")
		fmt.Printf("\tScore: %d\n\tHand:%v \n", player.MinScore(), player)
		fmt.Println("Dealer")
		fmt.Printf("\tScore: **HIDDEN**\n\tHand:%v \n", dealer)
		fmt.Println()
		fmt.Println("What will you do (h)it or (s)stand?")
		fmt.Scanf("%s", &input)

		switch input {
		case "h":
			card = deck.Draw()
			fmt.Println("Adding ", card)
			fmt.Println()
			player.Add(card)
		case "s":

		default:
			fmt.Println("Enter either 'h' or 's'")
		}
	}
}
