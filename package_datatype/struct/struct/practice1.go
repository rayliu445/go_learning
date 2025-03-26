package _struct

import "fmt"

type student struct {
	name string
	age  int
}

func PracStruct1() {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 20},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}

	for key, value := range m {
		fmt.Printf("%s=>%v\n", key, value.name)
	}
}
