# String matching algorithms written in go

## Purpose

* Educational

## Description

This project is an implementation of various string matching algorithms described
in [Introduction to algorithms](https://dl.ebooksworld.ir/books/Introduction.to.Algorithms.4th.Leiserson.Stein.Rivest.Cormen.MIT.Press.9780262046305.EBooksWorld.ir.pdf)<br>

## Algorithms Implemented

* Brute force
* Rabin Karp
* DFA (Deterministic finite automaton)
* Knuth Morris Pratt

## Time Complexity
 * n = length of the input text
 * m = length of the pattern
 * This table shows the worst case time
<table>
  <tr>
    <th>Algorithms</th>
    <th>Preprocessing time</th>
    <th>Matching time</th>
  </tr>
  <tr>
    <td>Brute Force</td>
    <td>0</td>
    <td>Ο((n - m + 1)m)</td>
  </tr>
<tr>
    <td>Rabin Karp</td>
    <td>Θ(m)</td>
    <td>Ο((n - m + 1)m)</td>
  </tr>
<tr>
    <td>DFA</td>
    <td>Ο(m |Σ|)</td>
    <td>Θ(n)</td>
  </tr>
<tr>
    <td>Knuth Morris Pratt</td>
    <td>Θ(m)</td>
    <td>Θ(n)</td>
  </tr>
</table>

### String Matching Definition
The problem of finding occurrence(s) of a pattern string within another string or body of text.
The algorithm returns an array of the first index for every occurrence in the text

## Example

```go
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

	pattern := "Gutenberg"

	indexes, err := KMP.MatchString(input, pattern)

	if err != nil {
		log.Fatal(err)
	}

	printMatches(indexes, input, len(pattern))
}

func printMatches(indexes []int, input string, patternLen int) {
	for _, index := range indexes {
		fmt.Printf("found at %d, %s\n", index, input[index:index+patternLen])
	}
}
```

### output

```terminal
found at 8, Gutenberg
found at 244, Gutenberg
```

# How to run
1. Clone the repo ```git clone https://github.com/Leonardpepa/string-matching-go.git```
2. Build ```go build```
3. run on windows```string-matching.exe```
4. run on linux ```./string-matching```
5. run tests ```go test```
