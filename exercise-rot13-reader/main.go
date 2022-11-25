package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int,  error) {
	n, err := rot13.r.Read(b)

	for i:= range b {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = 'A' + (b[i] - 'A' + 13) % 26
		} else if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = 'a' + (b[i] - 'a' + 13) % 26
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
