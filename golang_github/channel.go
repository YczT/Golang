package main

import (
	"fmt"
	"time"
)

//func recv(c chan int) {
//	re := <-c
//	fmt.Println("接收成功：", re)
//}
//func main() {
//	ch := make(chan int)
//	go recv(ch)
//	ch <- 10
//	fmt.Println("发送成功！")
//}

func demo() string {
	var result string
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "job"
	}()
	select {
	case result = <-ch:
		fmt.Println(result)
	case <-time.After(time.Second):
		return result
	}
	return result
}
func main() {
	ans := demo()
	fmt.Println(ans)
}
