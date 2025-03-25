package myreader

import (
	"io"
	"os"
	"strings"
)

type Rot13Reader struct {
	r io.Reader
}

func (rt Rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rt.r.Read(b)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}

	return n, nil
}

func rot13(c byte) byte {
	switch {
	case 'A' <= c && c <= 'Z':
		return (c-'A'+13)%26 + 'A'
	case 'a' <= c && c <= 'z':
		return (c-'a'+13)%26 + 'a'
	default:
		return c
	}
}

func ReadDeractor() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{s}
	io.Copy(os.Stdout, r)
}
