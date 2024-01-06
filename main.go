package main

import (
	"fmt"
	"string-matching/internal/DFA"
)

func main() {

	input := `The Project Gutenberg eBook of The Art of War`

	pattern := "eBook"

	indexes := DFA.MatchString(input, pattern)

	printMatches(indexes, input, len(pattern))
}

func printMatches(indexes []int, input string, patternLen int) {
	for _, val := range indexes {
		fmt.Println(input[val : val+patternLen])
	}
}
