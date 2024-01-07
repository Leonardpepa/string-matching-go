package main

import (
	"fmt"
	"log"
	"string-matching/internal/KMP"
)

func main() {
	input := `Project Gutenberg's International Short Stories: French, by Various

This eBook is for the use of anyone anywhere at no cost and with
almost no restrictions whatsoever.  You may copy it, give it away or
re-use it under the terms of the Project Gutenberg License included
with this eBook or online at www.gutenberg.net
`

	pattern := "the"

	indexes, err := KMP.MatchString(input, pattern)

	if err != nil {
		log.Fatal(err)
	}
	printMatches(indexes, input, len(pattern))
}

func printMatches(indexes []int, input string, patternLen int) {
	for _, val := range indexes {
		fmt.Printf("found at %d, %s\n", val, input[val:val+patternLen])
	}
}
