package math

import (
	"slices"
)

func PowMod(base int, exponent int, mod int) (int, error) {

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
