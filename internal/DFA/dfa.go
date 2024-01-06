package DFA

import (
	"strings"
)

type DFA struct {
	delta   []map[byte]int
	input   string
	pattern string
}

func InitDFA(input string, pattern string) *DFA {
	return &DFA{
		input:   input,
		pattern: pattern,
		delta:   make([]map[byte]int, 0, len(pattern)),
	}
}

func (dfa *DFA) CalcTransitionTable() {
	m := len(dfa.pattern)

	for q := 0; q <= m; q++ {
		dfa.delta = append(dfa.delta, make(map[byte]int))

		for a := 0; a < 256; a++ {
			k := min(m, q+1)

			//P[:q]a
			temp := dfa.pattern[:q] + string(byte(a))

			for {
				isSuffix := strings.HasSuffix(temp, dfa.pattern[:k])

				if isSuffix {
					break
				}

				k -= 1
			}

			dfa.delta[q][byte(a)] = k
		}
	}
}

func MatchString(input string, pattern string) []int {
	indexes := make([]int, 0)
	dfa := InitDFA(input, pattern)

	dfa.CalcTransitionTable()

	n := len(input)
	m := len(pattern)
	q := 0

	for i := 0; i < n; i++ {
		q = dfa.delta[q][input[i]]
		if q == m {
			indexes = append(indexes, i-m+1)
		}
	}

	return indexes
}
