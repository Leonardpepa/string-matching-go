package main

import (
	"string-matching/internal/DFA"
	"string-matching/internal/KMP"
	"string-matching/internal/bruteForce"
	"string-matching/internal/rabinKarp"
	"testing"
)

var input = `Project Gutenberg's International Short Stories: French, by Various

This eBook is for the use of anyone anywhere at no cost and with
almost no restrictions whatsoever.  You may copy it, give it away or
re-use it under the terms of the Project Gutenberg License included
with this eBook or online at www.gutenberg.net
`

var cases = []string{
	"Gutenberg",
	"Project",
	"online at www.gutenberg.net",
	"restrictions whats",
	"by Various\n\nThis eBook is for the use of anyone anywhere at no cost and with\nalmost no restrictions whatsoever.  You may copy it, give it away or",
	": French,",
	" ",
}

func shouldSucceed(t *testing.T, input string, pattern string, matcher func(input string, pattern string) ([]int, error)) {
	matches, err := matcher(input, pattern)

	if err != nil {
		t.Error(err)
	}

	for _, index := range matches {
		if input[index:index+len(pattern)] != pattern {
			t.Errorf("expected %s, got %s", pattern, input[index:index+len(pattern)])
		}
	}
}

func TestBruteForce(t *testing.T) {
	t.Run("brute force approach", func(t *testing.T) {
		for _, val := range cases {
			shouldSucceed(t, input, val, bruteForce.MatchString)
		}

		for _, val := range cases {
			shouldSucceed(t, input, val, bruteForce.MatchStringV2)
		}
	})
}

func TestRabinKarp(t *testing.T) {
	t.Run("rabin karp approach", func(t *testing.T) {
		for _, val := range cases {
			shouldSucceed(t, input, val, rabinKarp.MatchString)
		}
	})
}

func TestDFA(t *testing.T) {
	t.Run("DFA approach", func(t *testing.T) {
		for _, val := range cases {
			shouldSucceed(t, input, val, DFA.MatchString)
		}
	})
}

func TestKMP(t *testing.T) {
	t.Run("KMP approach", func(t *testing.T) {
		for _, val := range cases {
			shouldSucceed(t, input, val, KMP.MatchString)
		}
	})
}
