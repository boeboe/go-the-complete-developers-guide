package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range numbers {
		if n%2 == 1 {
			fmt.Printf("number %d is odd\n", n)
		} else {
			fmt.Printf("number %d is even\n", n)
		}
	}
}
