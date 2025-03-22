package myreader

import (
	"fmt"
	"io"
	"strings"
)

func ReadTest() {
	r := strings.NewReader("Hello world")
	b := make([]byte, 10)
	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%v b=%v\n", n, err, b)
		fmt.Printf("b[:n]=%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
