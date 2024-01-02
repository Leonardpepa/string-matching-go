package main

import "fmt"

func main() {

	input := `The Project Gutenberg eBook of The Art of War
    
This ebook is for the use of anyone anywhere in the United States and
most other parts of the world at no cost and with almost no restrictions
whatsoever. You may copy it, give it away or re-use it under the terms
of the Project Gutenberg License included with this ebook or online
at www.gutenberg.org. If you are not located in the United States,
you will have to check the laws of the country where you are located
before using this eBook.

Title: The Art of War


Author: active 6th century B.C. Sunzi

Translator: Lionel Giles

Release date: May 1, 1994 [eBook #132]
                Most recently updated: October 16, 2021

Language: English
`

	pattern := "located"

	s := simpleStringMatching(input, pattern)
	fmt.Println(input[s : s+len(pattern)])
}

func simpleStringMatching(input string, pattern string) int {
	n := len(input)
	m := len(pattern)

	if n < m {
		return -1
	}
	for s := 0; s < n-m; s++ {
		if pattern == input[s:s+m] {
			return s
		}
	}

	return -1
}
