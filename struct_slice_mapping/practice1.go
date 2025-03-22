package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	show := make([][]uint8, dy)
	for i := range show {
		show[i] = make([]uint8, dx)
		for j := range show[i] {
			show[i][j] = uint8(i * j)
		}
	}
	return show
}

func main() {
	pic.Show(Pic)
}
