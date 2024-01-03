package bruteForce

func MatchPattern(input string, pattern string) int {
	n := len(input)
	m := len(pattern)

	if n < m {
		return -1
	}
	for s := 0; s < n-m; s++ {
		if pattern == input[s:s+m] {
			return s
		}
	}

	return -1
}

func MatchPatternV2(input string, pattern string) int {
	n := len(input)
	m := len(pattern)

	if n < m {
		return -1
	}
	for s := 0; s < n-m; s++ {
		count := 0
		for i := s; i < s+m; i++ {
			if pattern[count] == input[i] {
				count++
			} else {
				break
			}
			if count == m {
				return s
			}
		}
	}

	return -1
}
