package main

import (
	"flag"
	"fmt"
	"strings"
)

// flag.Bool([variable name], [default value], [output message])
var ol = flag.Bool("ol", false, "omit trailing line")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*ol {
		fmt.Println()
	}
}
