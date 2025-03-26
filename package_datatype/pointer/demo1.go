package pointer

import (
	"fmt"
)

func TestPointer1() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)
}

func TestPointer2() {
	a := 10
	b := &a
	fmt.Printf("type of b:%T\n", b)
	c := *b
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}

func TestPointer3() {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%s\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
}

func TestPointer4() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["测试"] = 100
	fmt.Println(b)
}

// new function demo
func TestPointer5() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*a)
	fmt.Println(*b)
}

func TestPointer6() {
	//var b map[string]int
	//b = make(map[string]int)
	//b["测试"] = 100
	//fmt.Println(b)
}
