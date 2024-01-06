package shared

type PatternMatcher interface {
	MatchString() ([]int, error)
}
