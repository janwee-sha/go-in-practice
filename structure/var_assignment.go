package main

import (
	"fmt"
	"os"
)

func main() {
	z := gcd(100, 25)
	fmt.Println(z)

	f, _ := os.Open("foo.txt")
	if f != nil {
		fmt.Println(f.Name())
	}
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
