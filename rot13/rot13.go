/*
Package rot13 is a simple ceaser cipher for making questionable content opt-in.
*/
package rot13

import (
	"io"
	"unicode"
)

// Reader ciphers and deciphers text in rot13.
type Reader struct {
	reader io.Reader
	rotn   int
}

// NewReader is initialized and returned
func NewReader(r io.Reader, n int) Reader {
	return Reader{r, n}
}

/*	Read decodes or encodes text read in from a wrapped Reader.
 */
func (r Reader) Read(out []byte) (int, error) {
	n, err := r.reader.Read(out)

	if err != nil {
		return n, err
	}

	for i := range out[:n] {
		if ch := rune(out[i]); unicode.IsLetter(ch) {
			ltr := byte(int(out[i]) + (r.rotn % 26))

			// wrap alphabet
			if unicode.IsUpper(ch) {
				if ltr > 'Z' {
					ltr = 'A' - 1 + (ltr - 'Z')
				} else if ltr < 'A' {
					ltr = 'Z' - ('A' - ltr) + 1
				}
			} else {
				if ltr > 'z' {
					ltr = 'a' - 1 + (ltr - 'z')
				} else if ltr < 'a' {
					ltr = 'z' - ('a' - ltr) + 1
				}
			}
			out[i] = ltr
		}
	}
	return n, err
}
