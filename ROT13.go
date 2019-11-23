package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type ROT13Reader struct {
	r io.Reader
}

/*	The implementation of this method makes the ROT13Reader into a Reader by
	implicitly satisfying the Reader interface!
*/
func (reader ROT13Reader) Read(out []byte) (int, error) {
	n, err := reader.r.Read(out)

	if err != nil {
		return n, err
	}

	for i := range out {
		ch := rune(out[i])

		if unicode.IsLetter(ch) {
			ltr := out[i] - 13

			// wrap alphabet
			if unicode.IsUpper(ch) && ltr < 'A' {
				ltr = 'Z' - ('A' - ltr) + 1
			} else if ltr < 'a' {
				ltr = 'z' - ('a' - ltr) + 1
			}
			out[i] = ltr
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := ROT13Reader{s}
	io.Copy(os.Stdout, &r)
}
