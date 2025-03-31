package main

import (
	"fmt"
	"math"
)

func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint 8 range", n)
}

func main() {
	aInt := 1234567890123456789
	fmt.Println(Uint8FromInt(aInt))
}
