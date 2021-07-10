package main

import (
	"fmt"
	"os"
)

func writeString() {
	//create a file
	//if the file already exit, it will be truncate
	//returns a file descriptor
	file, err := os.Create("test2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//l - number of bytes written
	l, err := file.WriteString("I am the guardian of lost souls")
	if err != nil {
		fmt.Println(err)
		file.Close()  //closing the file if an error occurred
		return
	}

	fmt.Println("Wrote ", l, " bytes.")
	//close returns an error if the file is already closed
	err = file.Close()  //closing the file if writing was successful
	if err != nil {
		fmt.Println(err)
		return
	}
}

func writeBytes() {
	file, err := os.Create("bytes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//word "hello"
	bytes := []byte{104, 101, 108, 108, 111}
	noOfBytes, err := file.Write(bytes)
	if err != nil {
		fmt.Println(err)
		file.Close()  //close when error occurred
		return
	}

	fmt.Println("wrote ", noOfBytes, " bytes")
	err = file.Close()  //close when no error
	if err != nil {
		fmt.Println(err)
		return
	}
}

func writeLines() {
	file, err := os.Create("lines.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := []string{"This is a small world.", "May be too small.", "May be we should find another planet ASAP."}

	for _, v := range lines {
		//takes writer ass a parameter
		//appends new line
		fmt.Fprintln(file, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Wrote successfully")
}

func appendLines() {
	file, err := os.OpenFile("lines.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	newLine := "Aliens are invading us."
	_, err = fmt.Fprintln(file, newLine)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Appended successfully")
}

func main() {
	writeString()
	writeBytes()
	writeLines()
	appendLines()
}
