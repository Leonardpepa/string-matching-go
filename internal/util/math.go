package math

import (
	"errors"
)

func PowMod(base int, exponent int, mod int) (int, error) {

	if mod == 1 {
		return 0, nil
	}

	if base < 0 || exponent < 0 {
		return 0, errors.New("base and exponent need to be greater than 0")
	}

	if mod <= 0 {
		return 0, errors.New("modulo need to be positive")
	}

	base = base % mod
	result := 1
	for exponent > 0 {
		if exponent%2 == 1 {
			result = (result * base) % mod
		}
		// exponent/2
		exponent = exponent >> 1
		base = (base * base) % mod
	}

	return result, nil
}
