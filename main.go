package main

import (
	"fmt"
	"strings"
)

func repetition(st string) map[string]int {

	// using strings.Field Function
	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}

// Main function
func main() {

	input := "I love to eat apples and I also like to eat rice I like spicy food and also like sweet drinks."

	for index, element := range repetition(input) {
		fmt.Println(index, "=", element)
	}
}
