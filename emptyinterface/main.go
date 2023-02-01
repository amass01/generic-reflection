package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// We define a struct named PlayingCard with the properties Suit and Rank, to
// represent the cards from a deck of 52 playing cards. The Suit will be one of
// Diamonds, Hearts, Clubs, or Spades, and the Rank will be A, 2, 3, and so on
// through K.

type PlayingCard struct {
	Suit string
	Rank string
}

// Define a NewPlayingCard function to act as the constructor for the
// PlayingCard struct, and a String method, which will return the rank and
// suit of the card using fmt.Sprintf.

func NewPlayingCard(suit string, card string) *PlayingCard {
	return &PlayingCard{Suit: suit, Rank: card}
}

func (pc *PlayingCard) String() string {
	return fmt.Sprintf("%s of %s", pc.Rank, pc.Suit)
}

// Define the Deck struct with a field called cards to hold a slice of cards.
// Since you want the deck to be able to hold multiple different types of cards,
// you can’t just define it as []*PlayingCard, though. Define it as
// []interface{} so it can hold any type of card you may create in the future.

type Deck struct {
	cards []interface{}
}

// Create an AddCard method that accepts the same interface{} type to append a
// card to the Deck’s cards field.

func (d *Deck) AddCard(card interface{}) {
	d.cards = append(d.cards, card)
}

// Also, create a RandomCard method that will return a random card from the
// Deck’s cards slice. Use the math/rand package to generate a random number
// between 0 and the number of cards in the cards slice.

func (d *Deck) RandomCard() interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cardIdx := r.Intn(len(d.cards))
	return d.cards[cardIdx]
}

// Finally, create a NewPlayingCardDeck function which returns a *Deck value
// populated with all the cards in a playing card deck. Use the already defined
// two slices `suits` & `ranks`.

func NewPlayingCardDeck() *Deck {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	deck := &Deck{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.AddCard(NewPlayingCard(suit, rank))
		}
	}
	return deck
}

// Now you can use the pre-defined main func
func main() {
	deck := NewPlayingCardDeck()

	fmt.Printf("--- drawing playing card ---\n")
	card := deck.RandomCard()
	fmt.Printf("drew card: %s\n", card)

	playingCard, ok := card.(*PlayingCard)
	if !ok {
		fmt.Printf("card received wasn't a playing card!")
		os.Exit(1)
	}
	fmt.Printf("card suit: %s\n", playingCard.Suit)
	fmt.Printf("card rank: %s\n", playingCard.Rank)
}

// That's nice, but to access specific information about the *PlayingCard value
// we drew, we needed to do some extra work to convert the interface{} type
// into a  *PlayingCard type with access to the Suit and Rank fields. Using the
// Deck this way will work, but can also result in errors if a value other than
// *PlayingCard is added to the Deck. By updating our Deck to use interface,
// you can benefit from Go’s strong types and static type checking while still
// having the flexibility accepting interface{} values provides.
