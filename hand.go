package main

import (
	"fmt"
	pc "github.com/daman1807/deck"
	"strings"
)

type Hand struct {
	Name  string
	Cards []pc.Card
}

func (hand *Hand) MinScore() int{
	score := 0
	for _, card := range hand.Cards{
		// adding score of 10 for face cards
		if card.Rank() == pc.King || card.Rank() == pc.Queen || card.Rank() == pc.Jack{
			score += 10
		} else if card.Rank() == pc.Ace{
			score += 1
		}else{
			score += int(card.Rank().Value())
		}
	}
	return score
}
func (hand *Hand) Add(card pc.Card) {
	hand.Cards = append(hand.Cards, card)
}

func (hand Hand) String() string {
	cardStrings := make([]string, len(hand.Cards))

	for i := 0; i < len(hand.Cards); i++ {
		if hand.Name == "Dealer" && i == len(hand.Cards)-1 {
			cardStrings[i] = "**HIDDEN**"
		} else {
			cardStrings[i] = hand.Cards[i].String()
		}
	}
	return fmt.Sprintf("<Hand: %s| %s>", hand.Name, strings.Join(cardStrings, ", "))
}
