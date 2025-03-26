package _struct

import "fmt"

type newstudent struct {
	id   int
	name string
	age  int
}

func PracSlice1() {
	var ce []newstudent
	ce = []newstudent{
		newstudent{1, "xiaoming", 22},
		newstudent{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	demo1(ce)
	fmt.Println(ce)
}

func PracSlice2() {
	ce := make(map[int]newstudent)
	ce[1] = newstudent{1, "xiaoming", 22}
	ce[2] = newstudent{2, "xiaozhang", 33}
	fmt.Println(ce)
	delete(ce, 2)
	fmt.Println(ce)
}

func demo1(ce []newstudent) {
	ce[1].age = 999
	ce = append(ce, newstudent{1, "xiaoming", 44})
	fmt.Println(ce)
}
