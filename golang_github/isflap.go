package main

/*
字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。
*/
import "fmt"

func isFlipedString(s1 string, s2 string) bool {
	s1map := make(map[rune]int)
	s2map := make(map[rune]int)
	for _, e1 := range s1 {
		s1map[e1]++
	}
	for _, e2 := range s2 {
		s2map[e2]++
	}
	var i rune
	i = 0
	for i < 200 {
		if s1map[i] != s2map[i] {
			return false
		} else {
			i++
		}
	}
	return true
}
func main() {

	test_s1 := "chenyn"
	test_s2 := "yanchen"

	ans := isFlipedString(test_s1, test_s2)
	fmt.Println(ans)
}
