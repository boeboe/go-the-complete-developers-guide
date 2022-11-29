package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff00000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	var col map[string]string
	// this would fail (panic: assignment to entry in nil map)
	// col["white"] = "#ffffff"
	// col["black"] = "#000000"

	c := make(map[string]string)
	c["white"] = "#ffffff"
	c["black"] = "#000000"

	fmt.Println(colors)
	fmt.Println(col)
	fmt.Println(c)

	delete(colors, "red")
	fmt.Println(colors)

	printMap(c)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %q is %q\n", color, hex)
	}
}
