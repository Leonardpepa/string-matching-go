package rabinKarp

import "math"

func MatchString(input string, pattern string, d int, q int) int {
	n := len(input)
	m := len(pattern)

	h := int(math.Pow(float64(d), float64(m-1))) % q

	p := 0
	t := 0

	for i := 0; i < m; i++ {
		p = (d*p + int(pattern[i])) % q
		t = (d*t + int(input[i])) % q
	}

	for s := 0; s <= n-m; s++ {

		if p == t {
			if pattern == input[s:s+m] {
				return s
			}
		}

		if s < n-m {
			t = d*(t-int(input[s])*h) + int(input[s+m])
			// mod q
			t = t % q

			if t < 0 {
				t = t + q
			}
		}

	}

	return -1
}
