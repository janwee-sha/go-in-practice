package main

var g2 string

func main() {
	g2 = "G"
	print(g2)
	f1()
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(g2)
}
