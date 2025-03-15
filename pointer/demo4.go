package main

import "fmt"

type Vertex2 struct {
	X, Y int
}

var (
	v1 = Vertex2{1, 2}
	v2 = Vertex2{X: 1}
	v3 = Vertex2{}
	p  = &Vertex2{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
