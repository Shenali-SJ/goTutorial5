package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func openFile() {
	f, err := os.Open("/test.txt")
	if err != nil {  //there is an error
		if pErr ,ok := err.(*os.PathError); ok {
			//path of the file that occurred the error
			fmt.Println("Failed to open file at path ", pErr.Path)
			return
		}
		fmt.Println("generic error ", err)
		return
	}
	fmt.Println("Successfully opened file ", f)
}

func lookupError() {
	addr, err := net.LookupHost("shenaliJayakody.com")
	if err != nil {
		if dnsErr, ok := err.(*net.DNSError); ok {
			if dnsErr.Timeout() {
				fmt.Println("Operation timed out")
				return
			}
			if dnsErr.Temporary() {
				fmt.Println("Temporary error")
				return
			}
			fmt.Println("Generic DNS error", err)
			return
		}
		fmt.Println("Generic error ", err)
		return
	}
	fmt.Println(addr)
}

func checkPattern() {
	files, err := filepath.Glob("[")
	if err != nil {
		if err == filepath.ErrBadPattern {  //ErrBadPattern is a global variable in filePath
			fmt.Println("Bad pattern error: ", err)
			return
		}
		fmt.Println("Generic Error: ", err)
		return
	}
	fmt.Println("Matched files: ", files)
}

func main() {
	openFile()
	lookupError()
	checkPattern()
}
