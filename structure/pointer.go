package main

import (
	"fmt"
)

func main() {
	x := 1
	var px *int = &x
	*px = 2
	fmt.Println(px)  // address of x
	fmt.Println(*px) // x
	fmt.Println(x)   //x

	var y, z int
	fmt.Println(&y == &y, &y == &z, &y == nil) // true, false, false

	var pv = f()
	fmt.Println(*pv) // 1

	a := 100
	increase(&a)
	fmt.Println(a) //101
}

func f() *int {
	v := 1
	return &v // local variable v will still be existed after returning
}

func increase(p *int) int {
	*p++
	return *p
}
