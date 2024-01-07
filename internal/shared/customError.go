package shared

import "errors"

var NotFoundError = errors.New("match not found")
var BiggerPatternThanText = errors.New("pattern needs to be smaller or equal to input text")
var EmptyPattern = errors.New("pattern cannot be empty")
var EmptyInputText = errors.New("input text cannot be empty")
