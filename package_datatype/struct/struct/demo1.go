package _struct

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

type NewInt int

type MyInt = int

func TestStruct1() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)
}

func TestStruct2() {
	var p1 person
	p1.name = "pprof.cn"
	p1.city = "shanghai"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1.name=%v\n", p1.name)
}

func TestStruct3() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "pprof.cn"
	user.Age = 18
	fmt.Printf("user=%v\n", user)
}

func TestStruct4() {
	var p2 = new(person)
	fmt.Printf("p2=%T\n", p2)
	fmt.Printf("p2=%#v\n", p2)
}

func TestStruct5() {
	p3 := &person{}
	fmt.Printf("p3=%T\n", p3)
	fmt.Printf("p3=%#v\n", p3)
	p3.name = "pprof.cn"
	p3.city = "shanghai"
	p3.age = 18
	fmt.Printf("p3=%#v\n", p3)
}

func TestStruct6() {
	var p4 person
	fmt.Printf("p4=%#v\n", p4)
}

func TestStruct7() {
	p5 := person{
		city: "北京",
	}
	fmt.Printf("p5=%#v\n", p5)
	p7 := &person{
		name: "pprof.cn",
	}
	fmt.Printf("p7=%#v\n", p7)
}

// constructor func demo
func NewPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
