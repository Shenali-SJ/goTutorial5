package main

import (
	"fmt"
	"runtime/debug"
)

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from ", r)
		debug.PrintStack()
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverFullName()   //recover

	//panic msg and stack trace will be shown after the defer call
	defer fmt.Println("Defer call from fullName() function")
	if firstName == nil {
		panic("runtime error: first name cannot be empty")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be empty")
	}

	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")  //this will be executed if recovered
}

func main() {
	defer fmt.Println("Defer call from main()")
	firstName := "Shenali"
	fullName(&firstName, nil)

	fmt.Println("returned normally from main")  //this will be executed if recovered
}
