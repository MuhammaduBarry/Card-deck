package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// Which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	// Allows us to loop through both variables and make combinations
	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// Each return statement here has to return a new instance of deck
	// If we were to assign only one variable to this function, only the first statement would return
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",") // - Outputs a slice of string
	return deck(s)                      // - This returns us a deck by using type conversion
	// I already defined deck []string so i cant do a type conversion with the [], because go already knows it is a slice.
}

func (d deck) shuffleCards() {
	source := rand.NewSource(time.Now().UnixNano()) // - This generates a different int64 number every time we start this program
	// And we use it as the seed to generate a source object
	r := rand.New(source) // - Create a new random generator using that source

	// It insures that we have complete randomness and isnt generating the sameshuffled card deck
	r.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
