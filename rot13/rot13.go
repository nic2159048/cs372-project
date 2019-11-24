/*
Package rot13 is a simple ceaser cipher for making questionable content opt-in.
*/
package rot13

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Reader ciphers and deciphers text in rot13.
type Reader struct {
	reader io.Reader
	rotn   int
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
			ltr := byte(int(out[i]) + (r.rotn % 26))

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

// splits filename and rotation & posts updates for the user
func parseDecrypt(inFilename string) (outFilename string, rot int) {
	fmt.Printf("Decrypting %v\t", inFilename)

	end := strings.LastIndex(inFilename, ".rot")
	filename = inFilename[:end]

	n, err := strconv.ParseInt(inFilename[end:], 0, 0)
	if err != nil {
		panic("Rotation couldn't be parsed!")
	}
	rot = int(n)
	return
}

// parses rotation & prints stuff for the user
func parseEncrypt(outFilename string, intstr string) (outFilename string, rot int) {
	fmt.Printf("Ecrypting %v with rot=%s\t", outFilename, intstr)

	outFilename = outFilename + ".rot" + intstr

	n, err := strconv.ParseInt(intstr, 0, 0)
	if err != nil {
		panic("Rotation couldn't be parsed!")
	}
	rot = int(n)
	return
}

// rotates infile's content and writes to new file
func rotateFiles(inputFilename string, outputFilename string, rot int) {
	// attempt to read file
	fileIn, err := os.Open(inputFilename)
	if err != nil {
		panic(inputFilename + " is not readable!")
	}
	defer fileIn.Close()

	// open for writing
	fileOut, err := os.Create(outputFilename)
	if err != nil {
		panic("Couldn't make " + outputFilename)
	}
	defer fileOut.Close()

	// pass content over
	rotReader := Reader{fileIn, rot}
	for {
		var buff []byte
		_, err := rotReader.Read(buff)
		fileOut.Write(buff)

		if err == io.EOF {
			break
		}
	}
}

func main() {
	var inFilename, outFilename string
	var rot int

	switch {
	case len(os.Args) == 1: // ask for input
		fmt.Print("File? ")
		_, err := fmt.Scanf("%s", &inFilename)

		if strings.Contains(inFilename, ".rot") {
			outFilename, rot := parseDecrypt(inFilename)
		} else {
			fmt.Print("Rot? ")
			_, err = fmt.Scanf("%d", &rot)
			outFilename = outFilename + ".rot" + strconv.Itoa(rot)
			fmt.Printf("Ecrypting %v with rot=%v\t", outFilename, rot)
		}
	case len(os.Args) == 2 && strings.Contains(os.Args[1], ".rot"):
		inFilename = os.Args[1]
		outFilename, rot := parseDecrypt(inFilename)

	case len(os.Args) == 3:
		inFilename = os.Args[1]
		outFilename, rot := parseEncrypt(inFilename, os.Args[2])
	default:
		fmt.Printf("Unexpected Arguments: Expected 1 or 2 got %d", len(os.Args))
	}

	rotateFiles(inFilename, outFilename, rot)
	return
}

/*
s := strings.NewReader("Lbh penpxrq gur pbqr!")
r := Reader{s, 13}
io.Copy(os.Stdout, &r)
*/
/*
	Test Cases

	https://rot13.com/
	If this file is correctly encoded and decoded; then the ROT13/rot13 cipher is probably working.
	Vs guvf svyr vf pbeerpgyl rapbqrq naq qrpbqrq; gura gur EBG13/ebg13 pvcure vf cebonoyl jbexvat.
*/
