package shared

import "errors"

var NotFoundError = errors.New("match not found")
var BiggerPatternThanText = errors.New("pattern needs to be smaller or equal to input text")
