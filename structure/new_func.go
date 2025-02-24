package main

import "fmt"

func main() {
	p := new(int)
	i := *p

	fmt.Println(i) // 0

	*p = 2
	fmt.Println(i)  // 0
	fmt.Println(*p) // 2
}

func delta(old, new int) int { // 在delta函数内，内置的new函数是不可用的
	//p := new(int)
	//fmt.Println(*p) // 编译错误
	return new - old
}
