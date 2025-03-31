package main

import "fmt"

const c1 = -0.2879998977377777878777777787
const c2 = len("abc")

const beef, two, veg = "beef", 2, "veg"
const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	l1 = iota // 0
	l2        // 1
	l3        // 2
	l4 = 5    // 5
	l5        // 5
)

const (
	Apple, Banana = iota + 1, iota + 2 // 1,2
	Chris, Dune                        // 2,3
)

const (
	_        = iota
	KB int64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Println(l1, l2, l3, l4, l5)
	fmt.Println(KB, MB, GB, TB, PB, EB)
}
