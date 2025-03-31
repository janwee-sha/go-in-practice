package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	goos := runtime.GOOS
	fmt.Println(goos)
	path := os.Getenv("PATH")
	fmt.Println(path)
}
