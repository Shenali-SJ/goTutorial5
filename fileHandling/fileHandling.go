package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//1. should be run from the folder with test.txt file. Else error
func justRead() {
	data, err := ioutil.ReadFile("test.txt")

	if err != nil {
		fmt.Println("Error 1 :  ", err)
		return
	}
	fmt.Println("Content 1 : ", string(data))
}

func getAbsPath() {
	absPath, err2 := filepath.Abs("test.txt")
	if err2 != nil {
		fmt.Println("Error : ", err2)
		return
	}
	fmt.Println("Absolute path ", absPath)
}

//2. using absolute path
func readFromAbs() {
	data, err := ioutil.ReadFile("C:\\Users\\shena\\go\\src\\goTutorial5\\fileHandling\\test.txt")

	if err != nil {
		fmt.Println("Error 2 : ", err)
		return
	}
	fmt.Println("Content 2 : ", string(data))
}

//3. command line argument
func useFlag() {
	//this return an address
	filePath := flag.String("fPath", "test.txt", "file path to read from")
	flag.Parse()

	//since filePath is a pointer(address), de-reference should be passed
	data, err := ioutil.ReadFile(*filePath)
	if err != nil {
		fmt.Println("Error 3 : ", err)
		return
	}
	fmt.Println("Contents 3 : ", string(data))
}

//4. bundling text file with binary
func bundleText() {
	//box - a folder whose content will be embedded into binary
	//fileHandling folder's content will be embedded into binary
	box := packr.New("fileBox", "../fileHandling")
	data, err := box.FindString("test.txt")
	if err != nil {
		fmt.Println("Error 4 : ", err)
		return
	}
	fmt.Println("Content 4 : ", data)
}

//5. read by chunks
func readChunks() {
	//command line flag - returns the address of the file
	fPath := flag.String("path", "test.txt", "file path to read from")
	flag.Parse()

	//open the file
	file, err := os.Open(*fPath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	//create a buffered reader
	r := bufio.NewReader(file)
	//create a byte slice of capacity 3
	//bytes of the file will be read
	b := make([]byte, 3)

	for {
		//bytes returned are stored in variable n
		n, err2 := r.Read(b)
		if err2 != nil {
			fmt.Println("Error 5 : ", err2)
			break
		}
		fmt.Println(string(b[0:n]))
	}
}

//6. reading line by line
func readLines() {
	fPath := flag.String("path", "test.txt", "file path to read")
	flag.Parse()

	file, err := os.Open(*fPath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(file)
	for s.Scan() {
		fmt.Println(s.Text())
	}

	err2 := s.Err()
	if err2 != nil {
		log.Fatal(err2)
	}
}

func main() {

	justRead()
	getAbsPath()
	readFromAbs()
	useFlag()
	bundleText()

	fmt.Println()

	//readChunks()
	readLines()

}
