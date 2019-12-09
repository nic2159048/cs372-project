package main

/*
Since me and my partner disbanded our team my project is now located at: https://github.com/nic2159048/cs372-project

TA: Tito Ferra
Creative-program 
mandelVis.go
due date: 12/9/19

This program assumes you have ffmpeg, from https://www.ffmpeg.org/,
a built mandelbrot.exe from https://github.com/marijnfs/gomandel made by running go build mandelbrot.go
and the go-wav library from https://github.com/youpy/go-wav

This program will create a music video based on the samples of a given wav file.
The video portion will be a progressive sequence of generated mandelbrots that when combined with ffmpeg,
will look like a moving image instead of a collection of static images.





*/
import (
	"fmt"
	wav "github.com/youpy/go-wav"
	"io"
	"os"
	"os/exec"
	"testing"
	"log"
	"strconv"
)

func main() {
	//different sample rates you can use
	//n := []uint32{1, 10, 100, 1000, 2000, 3000, 5000, 8000, 10000, 20000, 40000}
	n := []uint32{1}
	
	var t int
	// num is used to add progressive rotation to the generated mandelbrot
	num := .001
	// curr is used to determine which run we are currently on for file naming
	// i.e out1, out2, out3, ...
	curr := 1
    
	// based on how many samples you give it generates a video for all of them
	for _, numSamples := range n {
		result := testing.Benchmark(func(b *testing.B) {
			// wav file 
			file, _ := os.Open("a.wav")
			reader := wav.NewReader(file)

			for {
				samples, err := reader.ReadSamples(numSamples)
				if err == io.EOF {
					break
				}
				for _, sample := range samples {
					// using the built mandelbrot generator mandel.exe we call it to create a mandelbrot and save it to out + num ie out1, out2, ...
					// the x and y are based on the int values that we get from the samples
					cmnd := exec.Command("mandel.exe",  "-x", strconv.Itoa(reader.IntValue(sample, 0)), "-y",  strconv.Itoa(reader.IntValue(sample, 1)), "-r", strconv.FormatFloat(num, 'E', -1,64), "-aa","1", "-out", "out"+fmt.Sprintf("%010d", curr))
					err := cmnd.Run() // and wait
					// if we error we stop computation
					if err != nil{
						log.Fatal(err)
						//fmt.Printf("%s\n",err.Error());
						return
					}
					curr+=1
					num+=.001
					// adds up the number of samples
					t += reader.IntValue(sample, 0)
					t += reader.IntValue(sample, 1)
				}
			}
		})
		// prints the number of samples 
		fmt.Printf("ReadSamples(%d): \t%s\n", numSamples, result.String())
		
		// we call ffmpeg to combine the mandelbrots that we created together into a video called test.pm4
		cmnd2 := exec.Command("ffmpeg.exe" , "-r", "60", "-f", "image2", "-s", "1920x1080", "-i", "out%010d.png", "-vcodec", "libx264", "-crf", "25",  "-pix_fmt", "yuv420p", "test.mp4")
		err1 := cmnd2.Run() // and wait
		// if we error we stop computation
		if err1 != nil{
			log.Fatal(err1)
			//fmt.Printf("%s\n",err.Error());
			return
		}
		// finally we add the wav file to the mp4 we generated using ffmpeg
		cmnd1 := exec.Command("ffmpeg", "-i", "test.mp4", "-i", "a.wav", "-vcodec", "libx264", "-acodec", "libmp3lame", "end.mp4")
		err2 := cmnd1.Run() // and wait
		// if we error we stop computation
		if err2 != nil{
			log.Fatal(err2)
			//fmt.Printf("%s\n",err.Error());
			return
		}
	}
}
