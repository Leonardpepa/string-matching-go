package KMP

import (
	"string-matching/internal/shared"
)

func computePrefix(pattern string) []int {
	m := len(pattern)

	p := make([]int, m)

	p[0] = 0

	k := 0

	for q := 1; q < m; q++ {
		for k > 0 && pattern[k] != pattern[q] {
			k = p[k-1]
		}

		if pattern[k] == pattern[q] {
			k += 1
		}

		p[q] = k
	}
	return p
}

func MatchString(input string, pattern string) ([]int, error) {
	if pattern == "" {
		return nil, shared.EmptyPattern
	}

	if input == "" {
		return nil, shared.EmptyInputText
	}

	indexes := make([]int, 0)

	p := computePrefix(pattern)
	q := 0

	n := len(input)
	m := len(pattern)

	if n < m {
		return nil, shared.BiggerPatternThanText
	}

	if n == m {
		if input == pattern {
			return []int{0}, nil
		}
		return nil, shared.NotFoundError
	}

	for i := 0; i < n; i++ {
		for q > 0 && pattern[q] != input[i] {
			q = p[q-1]
		}

		if pattern[q] == input[i] {
			q += 1
		}

		if q == m {
			indexes = append(indexes, i-m+1)
			q = p[q-1]
		}
	}

	if len(indexes) == 0 {
		return nil, shared.NotFoundError
	}
	return indexes, nil
}
