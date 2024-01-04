package rabinKarp

import (
	"errors"
)

func MatchString(input string, pattern string, d int, q int) ([]int, error) {
	indexes := make([]int, 0)

	n := len(input)
	m := len(pattern)

	if n < m {
		return nil, errors.New("pattern needs to be smaller or equal to input text")
	}

	h := 1
	// h ^ m-1 mod q
	for i := 0; i < m-1; i++ {
		h = (h * d) % q
	}

	p := 0
	t := 0

	for i := 0; i < m; i++ {
		p = (d*p + int(pattern[i])) % q
		t = (d*t + int(input[i])) % q
	}

	for s := 0; s <= n-m; s++ {

		if p == t && pattern == input[s:s+m] {
			indexes = append(indexes, s)
		}

		if s < n-m {
			t = (d*(t-int(input[s])*h) + int(input[s+m])) % q

			if t < 0 {
				t += q
			}
		}

	}

	if len(indexes) == 0 {
		return nil, errors.New("match not found")
	}
	return indexes, nil
}
