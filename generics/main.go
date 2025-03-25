package main

import (
	"fmt"
	"go_learning/generics/generics"
)

func main() {
	si := []int{10, 20, 15, -10}
	fmt.Println(generics.Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(generics.Index(ss, "hello"))
}
