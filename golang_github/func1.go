package main

import "fmt"

func person(name string, age int) (string, int) {
	nAme := name
	aGe := age
	return nAme, aGe
}

func main() {
	n := "chen"
	a := 24
	n1, a1 := person(n, a)
	fmt.Printf("%v,%v\n", n1, a1)
}
