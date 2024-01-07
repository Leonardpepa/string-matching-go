package main

import (
	"errors"
	"string-matching/internal/DFA"
	"string-matching/internal/KMP"
	"string-matching/internal/bruteForce"
	"string-matching/internal/rabinKarp"
	"string-matching/internal/shared"
	"testing"
)

var input = `Project Gutenberg's International Short Stories: French, by Various

This eBook is for the use of anyone anywhere at no cost and with
almost no restrictions whatsoever.  You may copy it, give it away or
re-use it under the terms of the Project Gutenberg License included
with this eBook or online at www.gutenberg.net
`

var successCases = []string{
	"Gutenberg",
	"Project",
	"online at www.gutenberg.net",
	"restrictions whats",
	"by Various\n\nThis eBook is for the use of anyone anywhere at no cost and with\nalmost no restrictions whatsoever.  You may copy it, give it away or",
	": French,",
	input,
}

var failCases = []struct {
	pattern string
	input   string
	err     error
}{
	{pattern: "", input: input, err: shared.EmptyPattern},
	{pattern: "WORD", input: input, err: shared.NotFoundError},
	{pattern: "it doesnt exist", input: input, err: shared.NotFoundError},
	{pattern: "Restrictions", input: input, err: shared.NotFoundError},
	{pattern: "yes", input: "no", err: shared.BiggerPatternThanText},
	{pattern: "yes", input: "", err: shared.EmptyInputText},
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

func shouldFail(t *testing.T, input string, pattern string, matcher func(input string, pattern string) ([]int, error), expectedErr error) {
	_, err := matcher(input, pattern)

	if err == nil {
		t.Error("found wrong match")
	}

	if !errors.Is(err, expectedErr) {
		t.Errorf("pattern: %s expected %v, got %v", pattern, expectedErr, err)
	}
}

func TestBruteForce(t *testing.T) {
	t.Run("brute force approach", func(t *testing.T) {
		for _, val := range successCases {
			shouldSucceed(t, input, val, bruteForce.MatchString)
		}

		for _, val := range successCases {
			shouldSucceed(t, input, val, bruteForce.MatchStringV2)
		}

		for _, val := range failCases {
			shouldFail(t, val.input, val.pattern, bruteForce.MatchString, val.err)
		}
		for _, val := range failCases {
			shouldFail(t, val.input, val.pattern, bruteForce.MatchStringV2, val.err)
		}
	})
}

func TestRabinKarp(t *testing.T) {
	t.Run("rabin karp approach", func(t *testing.T) {
		for _, val := range successCases {
			shouldSucceed(t, input, val, rabinKarp.MatchString)
		}

		for _, val := range failCases {
			shouldFail(t, val.input, val.pattern, rabinKarp.MatchString, val.err)
		}
	})
}

func TestDFA(t *testing.T) {
	t.Run("DFA approach", func(t *testing.T) {
		for _, val := range successCases {
			shouldSucceed(t, input, val, DFA.MatchString)
		}
		for _, val := range failCases {
			shouldFail(t, val.input, val.pattern, DFA.MatchString, val.err)
		}
	})
}

func TestKMP(t *testing.T) {
	t.Run("KMP approach", func(t *testing.T) {
		for _, val := range successCases {
			shouldSucceed(t, input, val, KMP.MatchString)
		}

		for _, val := range failCases {
			shouldFail(t, val.input, val.pattern, KMP.MatchString, val.err)
		}
	})
}
