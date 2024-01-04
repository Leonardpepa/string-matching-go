package math

import (
	"errors"
	"slices"
)

func PowMod(base int, exponent int, mod int) (int, error) {
	if base < 0 || exponent < 0 {
		return 0, errors.New("base and exponent need to be greater than 0")
	}

	if mod <= 0 {
		return 0, errors.New("modulo need to be positive")
	}

	d := 1
	bin := toBinaryArray(exponent)

	for _, val := range bin {
		d = (d * d) % mod

		if val == 1 {
			d = (d * base) % mod
		}
	}

	return d, nil
}

func toBinaryArray(num int) []int8 {
	bin := make([]int8, 0)

	for num > 0 {
		bin = append(bin, int8(num%2))
		num = num / 2
	}

	slices.Reverse(bin)
	return bin
}
