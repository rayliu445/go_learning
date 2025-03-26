package _struct

import "fmt"

type Person struct {
	name string
	city string
	age  int8
}

type Person1 struct {
	string
	int
}

type Address struct {
	Province   string
	City       string
	CreateTime string
}

type Email struct {
	Account    string
	CreateTime string
}

type User struct {
	Name   string
	Gender string
	Address
	Email
}

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动!\n", a.name)
}

type Dog struct {
	Feet int8
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪～\n", d.name)
}

func (p *Person) setAge(age int8) {
	p.age = age
}

func TestAnonymousStruct() {
	p1 := Person1{
		"pprof.cn",
		18,
	}
	fmt.Println(p1.string, p1.int)
}

func TestNestAnonymousStruct() {
	var user3 User
	user3.Name = "pprof.cn"
	user3.Gender = "male"
	user3.Address.CreateTime = "2000"
	user3.Email.CreateTime = "2000"
}

func TestStructExtends() {
	d1 := &Dog{
		Feet: 10,
		Animal: &Animal{
			name: "乐乐",
		},
	}
	d1.wang()
	d1.move()
}
