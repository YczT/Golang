package main

import "fmt"

func replaceSpaces(S string, length int) string {
	var re = make([]string, len(S))
	str := ""
	for i, v := range S {
		fmt.Printf("%v\n", i)
		if i+1 > length {
			continue
		}
		re[i] = string(v)
	}
	for i, _ := range re {
		if re[i] == " " {
			re[i] = "%20"
		}
		str = str + re[i]
	}

	return str
}
func main() {
	s1 := "csn suj s     "
	l := len(s1)
	fmt.Printf("%v\n", l)
	s := replaceSpaces(s1, l)

	fmt.Printf("%v\n", s)
}
