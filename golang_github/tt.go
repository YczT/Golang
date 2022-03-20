package main

import "fmt"

func main() {
	type User struct {
		name string
		age  int
	}
	admin := User{
		name: "goods",
		age:  21,
	}

	p := admin

	fmt.Println(p)
	tflag := make([]int, 1)
	flag := tflag[:]
	flag[0] = -1
	fmt.Println(flag)
}
