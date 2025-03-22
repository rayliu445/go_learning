package main

import "fmt"

type Vertex3 struct {
	Lat, Long float64
}

var m map[string]Vertex3

func main() {
	m = make(map[string]Vertex3)
	m["Bell Labs"] = Vertex3{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
