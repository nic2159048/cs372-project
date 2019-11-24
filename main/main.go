/*
Links to other files

sample.txt.rot13

If this file is correctly encoded and decoded; then the ROT13/rot13 cipher is probably working.
Vs guvf svyr vf pbeerpgyl rapbqrq naq qrpbqrq; gura gur EBG13/ebg13 pvcure vf cebonoyl jbexvat.
Lbh penpxrq gur pbqr!
You cracked the code!
*/
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ivanthewebber/csc372-program3/rot13"
)

// splits filename and rotation & posts updates for the user
func parseDecrypt(inFilename string) (outFilename string, rot int) {
	fmt.Printf("Decrypting %v\t", inFilename)

	end := strings.LastIndex(inFilename, ".rot")
	outFilename = inFilename[:end]

	n, err := strconv.ParseInt(inFilename[end+4:], 0, 0)
	if err != nil {
		panic("Rotation couldn't be parsed!")
	}
	rot = -int(n)
	return
}

// parses rotation & prints stuff for the user
func parseEncrypt(inFilename string, intstr string) (outFilename string, rot int) {
	fmt.Printf("Ecrypting %v with rot=%s\t", inFilename, intstr)

	outFilename = inFilename + ".rot" + intstr

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

	rotReader := rot13.NewReader(fileIn, rot)

	var buff []byte
	buff = make([]byte, 32)

	// pass content over
	for {

		n, err := rotReader.Read(buff)
		fileOut.Write(buff[:n])

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

		if err != nil {
			panic(err)
		}

		if strings.Contains(inFilename, ".rot") {
			outFilename, rot = parseDecrypt(inFilename)
		} else {
			fmt.Print("Rot? ")
			_, err = fmt.Scanf("%d", &rot)
			outFilename = outFilename + ".rot" + strconv.Itoa(rot)
			fmt.Printf("Ecrypting %v with rot=%v\t", outFilename, rot)
		}
	case len(os.Args) == 2 && strings.Contains(os.Args[1], ".rot"):
		inFilename = os.Args[1]
		outFilename, rot = parseDecrypt(inFilename)

	case len(os.Args) == 3:
		inFilename = os.Args[1]
		outFilename, rot = parseEncrypt(inFilename, os.Args[2])
	default:
		fmt.Printf("Unexpected Arguments: Expected 1 or 2 got %d", len(os.Args))
	}

	rotateFiles(inFilename, outFilename, rot)
	return
}
