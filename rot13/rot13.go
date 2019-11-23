/*
Package rot13 is a simple ceaser cipher for making questionable content opt-in.
*/
package rot13

import (
	"io"
	"os"
	"strings"
	"unicode"
)

// Reader ciphers and deciphers text in rot13.
type Reader struct {
	reader io.Reader
}

/*	Read decodes or encodes text read in from a wrapped Reader.
 */
func (r Reader) Read(out []byte) (int, error) {
	n, err := r.reader.Read(out)

	if err != nil {
		return n, err
	}

	for i := range out {
		if ch := rune(out[i]); unicode.IsLetter(ch) {
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
	r := Reader{s}
	io.Copy(os.Stdout, &r)
}

/*
	Test Cases

	https://rot13.com/
	If this file is correctly encoded and decoded; then the ROT13/rot13 cipher is probably working.
	Vs guvf svyr vf pbeerpgyl rapbqrq naq qrpbqrq; gura gur EBG13/ebg13 pvcure vf cebonoyl jbexvat.
*/
