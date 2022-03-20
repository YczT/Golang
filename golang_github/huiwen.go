package main

import "fmt"

func canPermutePalindrome(s string) bool {
	var strmap = make(map[rune]int, 26)
	for _, e := range s {
		strmap[e] = strmap[e] + 1
	}
	signal := 0
	for _, e := range strmap {
		if e%2 != 0 {
			signal++
		}
	}
	fmt.Printf("%v\n", signal)
	return signal == 0 || signal == 1
}

func main() {
	s := "code"
	ans := canPermutePalindrome(s)
	fmt.Printf("%v\n", ans)
}
