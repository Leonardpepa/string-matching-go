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

func MatchString(T string, P string) ([]int, error) {
	indexes := make([]int, 0)

	p := computePrefix(P)
	q := 0

	n := len(T)
	m := len(P)

	for i := 0; i < n; i++ {
		for q > 0 && P[q] != T[i] {
			q = p[q-1]
		}

		if P[q] == T[i] {
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
