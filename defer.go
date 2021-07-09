package main

import (
	"fmt"
)

func greet() {
	fmt.Println("Welcome")
}

func printNum() {
	defer greet() //this is invoked at the end
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("Printed numbers in order")
}

func increaseByTen(num int) (sum int) {
	sum = num + 10
	fmt.Println("In function ", num)
	return
}

func printReverse(word string) {
	for _, v := range []rune(word) {
		defer fmt.Printf("%c \n", v)
	}
}

func main() {
	printNum()

	a := 12
	defer increaseByTen(a)
	a = 56
	fmt.Println("Before executing ", a)

	fmt.Println()
	printReverse("Shenali")
}
