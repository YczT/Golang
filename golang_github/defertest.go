package main

import "fmt"

func f() int {
	a := 0
	defer func(i int) {
		println("defer i=", i)
	}(a)
	a++
	fmt.Println(a)
	return a
}
func b() {
	for i := 0; i < 5; i++ {
		defer println(i)
	}
	return
}
func main() {
	b()
}
