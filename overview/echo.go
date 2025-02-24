package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Echo with iteration: ")
	echoWithIteration()

	fmt.Print("Echo with range: ")
	echoWithRange()

	fmt.Print("Echo with join: ")
	echoWithStringsJoin()

	fmt.Print("Echo with println: ")
	echoWithPrintln()
}

func echoWithIteration() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echoWithRange() {
	s, sep := "", ""
	// 每次迭代range都会产生一对值：索引和索引处的值
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echoWithStringsJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echoWithPrintln() {
	fmt.Println(os.Args[1:])
}
