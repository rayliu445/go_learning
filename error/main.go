package main

import (
	"fmt"
	"go_learning/error/error"
)

func main() {
	//if err := error.Run(); err != nil {
	//	fmt.Println(err)
	//}
	fmt.Println(error.Sqrt(2))
	fmt.Println(error.Sqrt(-2))
}
