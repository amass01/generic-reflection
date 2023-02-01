package main

import (
	"fmt"
	"math/rand"
	"time"
)

type PlayingCard struct {
	Suit string
	Rank string
}

func NewPlayingCard(suit string, card string) *PlayingCard {
	return &PlayingCard{Suit: suit, Rank: card}
}

func (pc *PlayingCard) String() string {
	return fmt.Sprintf("%s of %s", pc.Rank, pc.Suit)
}

// Update the Deck struct to be generic type

type Deck[C any] struct {
	cards []C
}

// Update AddCard to use the generic type

func (d *Deck[C]) AddCard(card C) {
	d.cards = append(d.cards, card)
}

// Update the RandomCard method to use the C generic type as well

func (d *Deck[C]) RandomCard() C {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cardIdx := r.Intn(len(d.cards))
	return d.cards[cardIdx]
}

// Update NewPlayingCardDeck function to use the generic Deck type for
// *PlayingCard types

func NewPlayingCardDeck() *Deck[*PlayingCard] {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	deck := &Deck[*PlayingCard]{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.AddCard(NewPlayingCard(suit, rank))
		}
	}
	return deck
}

// Now update the main func to use the generic types

func main() {
	deck := NewPlayingCardDeck()

	fmt.Printf("--- drawing playing card ---\n")
	playingCard := deck.RandomCard()
	fmt.Printf("drew card: %s\n", playingCard)
	// Code removed
	fmt.Printf("card suit: %s\n", playingCard.Suit)
	fmt.Printf("card rank: %s\n", playingCard.Rank)
}

// We no longer need to assert an interface{} value into a *PlayingCard value.
// When you updated Deck’s RandomCard method to return C and updated
// NewPlayingCardDeck to return *Deck[*PlayingCard], it changed RandomCard to
// return a *PlayingCard value instead of interface{}. When RandomCard returns
// *PlayingCard, it means the type of playingCard is also *PlayingCard instead
// of interface{} and you can access the Suit or Rank fields right away.
