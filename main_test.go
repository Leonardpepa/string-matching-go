package main

import (
	"string-matching/internal/bruteForce"
	"testing"
)

func TestBruteForce(t *testing.T) {
	t.Run("brute force approach", func(t *testing.T) {
		input := `Project Gutenberg's International Short Stories: French, by Various

This eBook is for the use of anyone anywhere at no cost and with
almost no restrictions whatsoever.  You may copy it, give it away or
re-use it under the terms of the Project Gutenberg License included
with this eBook or online at www.gutenberg.net
`

		cases := []string{
			"Gutenberg",
			"Project",
			"online at www.gutenberg.net",
			"restrictions whats",
			"by Various\n\nThis eBook is for the use of anyone anywhere at no cost and with\nalmost no restrictions whatsoever.  You may copy it, give it away or",
			": French,",
			" ",
		}

		for _, val := range cases {
			shouldSucceed(t, input, val, bruteForce.MatchString)
		}

		for _, val := range cases {
			shouldSucceed(t, input, val, bruteForce.MatchStringV2)
		}
	})
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
