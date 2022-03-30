package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("hello", i)
}
func main() {
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go hello(i)
	//}
	//wg.Wait()

	ans := f3()
	fmt.Println(ans)
}

/*
多次执行上面的代码会发现每次终端上打印数字的顺序都不一致。这是因为10个 goroutine 是并发执行的，而 goroutine 的调度是随机的。
*/

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
