package main

import "fmt"

func main() {
	var n int16 = 34
	var m int32
	//m == n // 编译错误
	m = int32(n)

	fmt.Printf("n is %d\n", n)
	fmt.Printf("m is %d\n", m)

	var a32BitInt int32
	var a32Float float32 = 3.141592693
	a32BitInt = int32(a32Float)
	fmt.Println(a32BitInt)
}
