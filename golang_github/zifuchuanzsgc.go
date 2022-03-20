package main

import (
	"fmt"
)

/*
func oneEditAway(first string, second string) bool {
	fmap := make(map[rune]int, 26)
	smap := make(map[rune]int, 26)
	fcount := 0
	scount := 0
	for _, ef := range first {
		fmap[ef] = fmap[ef] + 1
		fcount++
	}
	for _, es := range second {
		smap[es] = smap[es] + 1
		scount++
	}
	if math.abs(fcount-scount) >= 2 {
		return false
	} else {

	}

	return
}
*/
func oneEditAway(first string, second string) bool {
	if first == second {
		return true
	}
	if len(first)+len(second) <= 1 {
		return true
	}
	flag := false
	ef, es := 0, 0
	for ef < len(first) || es < len(second) {
		if ef < len(first) && es < len(second) && first[ef] == second[es] {
			ef++
			es++
			continue
		} else if !flag {
			if len(first)-ef > len(second)-es { //first insert
				ef++
			} else if len(first)-ef < len(second)-es { //second insert
				es++
			} else { //replace
				ef++
				es++
			}
			flag = true
		} else {
			return true
		}
	}
	if ef == len(first) && es == len(second) {
		return true
	}
	return false
}

func main() {
	f := "a"
	s := "ab"
	ans := oneEditAway(f, s)
	fmt.Printf("%v\n", ans)
}
