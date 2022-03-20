package main

import "fmt"

func CheckPermutation(s1 string, s2 string) bool {
	match1 := make(map[rune]int)
	match2 := make(map[rune]int)
	for _, e1 := range s1 {
		match1[e1]++
	}
	for _, e2 := range s2 {
		match2[e2]++
	}
	var sq1 rune = 0
	var sq2 rune = 0
	for sq1 < 100 {
		if match1[sq1] != match2[sq2] {
			fmt.Printf("%v,%v\n", match1[sq1], match2[sq2])
			return false
		} else {
			sq1++
			sq2++
		}
	}
	fmt.Printf("%v,%v\n", match1[sq1], match2[sq2])
	return true
}
func main() {
	s1 := "abc"
	s2 := "bce"

	ans := CheckPermutation(s1, s2)

	fmt.Printf("%v\n", ans)

}
