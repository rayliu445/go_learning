package main

import (
	"fmt"
	"go_learning/io/myreader"
	"golang.org/x/tour/reader"
)

func main() {
	//myreader.ReadTest()
	reader.Validate(myreader.MyReader{})

	r := myreader.MyReader{}
	b := make([]byte, 10)
	n, err := r.Read(b)
	fmt.Println(n, err, string(b))
	//fmt.Println("test...")
}
