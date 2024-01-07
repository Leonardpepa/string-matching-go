package rabinKarp

import (
	"string-matching/internal/shared"
	"string-matching/internal/util"
)

func MatchString(input string, pattern string) ([]int, error) {
	if pattern == "" {
		return nil, shared.EmptyPattern
	}

	if input == "" {
		return nil, shared.EmptyInputText
	}

	indexes := make([]int, 0)

	d := 256
	prime := 16777619
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

	// d ^ m-1 mod prime
	h, err := modulo.PowMod(d, m-1, prime)

	if err != nil {
		return nil, err
	}

	p := 0
	t := 0

	for i := 0; i < m; i++ {
		p = (d*p + int(pattern[i])) % prime
		t = (d*t + int(input[i])) % prime
	}

	for s := 0; s <= n-m; s++ {

		if p == t && pattern == input[s:s+m] {
			indexes = append(indexes, s)
		}

		if s < n-m {
			t = (d*(t-int(input[s])*h) + int(input[s+m])) % prime

			if t < 0 {
				t += prime
			}
		}

	}

	if len(indexes) == 0 {
		return nil, shared.NotFoundError
	}
	return indexes, nil
}
