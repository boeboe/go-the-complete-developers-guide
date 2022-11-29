package main

import (
	"fmt"
	"strings"
)

/* Create a new type of 'card' which is a struct. */
type card struct {
	suit  string
	value string
}

/* Convert card to a single string */
func (c card) toString() string {
	return fmt.Sprintf("%v of %v", c.value, c.suit)
}

/* Convert string to a card */
func cardFromString(s string) card {
	sp := strings.Split(s, " of ")
	return card{
		suit:  sp[0],
		value: sp[1],
	}
}
