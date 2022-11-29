package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/* Create a new type of 'deck' which is a slice of strings. */
type deck []string

/* Create a new deck */
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}
	return cards
}

/* Print the cards in a deck */
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/* Deal an amount of cards from a deck */
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/* Convert deck to a single string */
func (d deck) toString() string {
	return strings.Join(d, "\n")
}

/* Convert deck to a single string */
func fromString(deckString string) deck {
	return strings.Split(deckString, "\n")
}

/* Save a deck to a file */
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

/* Load a deck from a file */
func loadDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return fromString(string(bs))
}

/* Shuffle a deck */
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		j := r.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}
}
