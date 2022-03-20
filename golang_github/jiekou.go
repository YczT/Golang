package main

import "fmt"

/*
通过如下方法判断接口是否实现
*/
type Inter interface {
	ping()
	pang()
}
type Anter interface {
	Inter
	String()
}
type St struct {
	Name string
}

func (St) ping() {
	print("ping\n")
}
func (*St) pang() {
	print("pang\n")
}

func (St) String() {
	print("AnterString\n")
}
func main() {
	st := &St{
		Name: "chenyan_handsome",
	}
	var i interface{} = st //将具体类型变量（对象）存到空接口变量中
	o := i.(Inter)         //断言赋值方法，通过接口变量实现断言
	o.ping()
	o.pang()
	var j interface{} = st
	p := j.(Anter)
	p.String()
	p.ping()
	p.pang()
	//判断i绑定的实例是否是具体类型St
	s := i.(*St)
	fmt.Printf("%v\n", s)
}
