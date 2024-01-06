package DFA

import (
	"log"
	"strings"
)

type DFA struct {
	states  []map[byte]State
	input   string
	pattern string
}

func InitDFA(input string, pattern string) *DFA {
	return &DFA{
		input:   input,
		pattern: pattern,
		states:  make([]map[byte]State, 0, len(pattern)),
	}
}

type State struct {
	next int
}

func (dfa *DFA) addTransition(index int, a byte, k int) {
	dfa.states[index][a] = State{next: k}
}

func (dfa *DFA) CalcTransitionTable() {
	m := len(dfa.pattern)

	for q := 0; q < m; q++ {
		dfa.states = append(dfa.states, make(map[byte]State))
		for i := 0; i < 256; i++ {
			k := min(m, q+1)

			for isSuffix := strings.HasSuffix(dfa.pattern[:q], dfa.pattern[:k-1]); !isSuffix; {
				k--
				if k == 0 {
					break
				}
			}

			dfa.addTransition(q, byte(i), k)
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
		q := dfa.states[q][input[i]].next
		log.Println(q, m)
		if q == m {
			indexes = append(indexes, i-m)
		}
	}

	return indexes
}
