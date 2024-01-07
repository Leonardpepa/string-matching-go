package bruteForce

import (
	"string-matching/internal/shared"
)

func MatchString(input string, pattern string) ([]int, error) {
	if pattern == "" {
		return nil, shared.EmptyPattern
	}

	if input == "" {
		return nil, shared.EmptyInputText
	}

	indexes := make([]int, 0)
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

	for s := 0; s <= n-m; s++ {
		if pattern == input[s:s+m] {
			indexes = append(indexes, s)
		}
	}

	if len(indexes) == 0 {
		return nil, shared.NotFoundError
	}

	return indexes, nil
}

func MatchStringV2(input string, pattern string) ([]int, error) {
	if pattern == "" {
		return nil, shared.EmptyPattern
	}

	if input == "" {
		return nil, shared.EmptyInputText
	}

	indexes := make([]int, 0)

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

	for s := 0; s <= n-m; s++ {
		count := 0
		for i := s; i < s+m; i++ {
			if pattern[count] == input[i] {
				count++
			} else {
				break
			}
			if count == m {
				indexes = append(indexes, s)
			}
		}
	}

	if len(indexes) == 0 {
		return nil, shared.NotFoundError
	}

	return indexes, nil
}
