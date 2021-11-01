package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// func toBytes(s string) []byte {
// 	return []byte(s)
// }

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// func readFromFile(filename string) ([]byte, error) {
// 	return ioutil.ReadFile(filename)
// }

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) // bs - byte slice
	if err != nil {
		// Option 1 - log the error and return a call to newDeck()
		// Option 2 - log the error and entirely quit hte programm
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",") // slice of strings
	return deck(s)
}

func (d deck) shuffle() {

	/* time.Now package returns an t time type.
	UnixNano needs a t time type to return int64 type.
	We needed that to create
	a new int64 for source each time we run our program*/

	source := rand.NewSource(time.Now().UnixNano()) // creating of source of randome numbers (seed)
	r := rand.New(source)                           // creating of new Rand type that returns a randome nums from source

	for i := range d { // i - index
		// newPosition := rand.Intn(len(d) - 1)
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
