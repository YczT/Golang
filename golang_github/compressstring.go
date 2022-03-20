package main

import (
	"fmt"
	"strconv"
)

/*
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。
比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返
回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。
*/
func compressString(S string) string {
	if len(S) <= 0 || S == "" {
		return S
	}
	cstring := ""
	scount := 1
	i, j := 0, 1
	for ; j < len(S); j++ {
		if S[i:i+1] == S[j:j+1] {
			scount++
		} else {
			cstring = cstring + S[i:i+1] + strconv.Itoa(scount)
			i = j
			scount = 1
		}
	}
	cstring = cstring + S[i:i+1] + strconv.Itoa(scount)
	if len(cstring) >= len(S) {
		return S
	}
	return cstring
}

func main() {
	test_s := "aa"
	ans := compressString(test_s)
	fmt.Printf("%v\n", ans)
}
