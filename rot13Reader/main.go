// https://tour.golang.org/methods/23
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		if b[i] < 'A' || b[i] > 'z' || b[i] > 'Z' && b[i] < 'a' {
			continue
		}

		b[i] += 13
		if b[i] > 'z' || b[i] < 'a' && b[i] > 'Z' {
			b[i] -= 26
		}
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
