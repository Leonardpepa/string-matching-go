package main

import (
	"fmt"
	"math"
)

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

	pattern := "recently"
	s := RabinKarp(input, pattern, 256, 13)
	fmt.Println(input[s : s+len(pattern)])
}

func RabinKarp(input string, pattern string, d int, q int) int {
	n := len(input)
	m := len(pattern)

	h := int(math.Pow(float64(d), float64(m-1))) % q

	p := 0
	t := 0

	for i := 0; i < m; i++ {
		p = (d*p + int(pattern[i])) % q
		t = (d*t + int(input[i])) % q
	}

	for s := 0; s < n-m; s++ {

		if p == t {
			if pattern == input[s:s+m] {
				return s
			}
		}

		if s < n-m {
			t = (d*(t-int(input[s])*h) + int(input[s+m])) % q

			for t < 0 {
				t = t + q
			}
		}

	}

	return -1
}
