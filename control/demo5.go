package main

import (
	"fmt"
	"math"
)

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", x, n)
	}
	return lim
}

func main() {
	fmt.Println(pow1(2, 3, 10), pow1(2, 3, 10))
}
