package main

import "fmt"

var global *int

func main() {
	globalPoint2X()

	fmt.Println(*global) // 1

	y := returnedVariablePointToY()
	fmt.Println(*y) // 1
}

func globalPoint2X() {
	var x int // 这里的x使用堆空间，因为它在函数返回以后还可以从global变量访问
	x = 1
	global = &x
}

func returnedVariablePointToY() *int {
	y := new(int)
	*y = 1
	return y
}
