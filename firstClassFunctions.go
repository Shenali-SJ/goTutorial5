package main

import "fmt"

type word func(s string, name string) string

func handleNum(add func (num1 int, num2 int) int, a int, b int ) {
	fmt.Println(add(a, b))
}

func subtractNum() func(a, b int) int {
	f := func(a, b int) int {
		return a - b
	}
	return f
}

func increaseNum() func(int) int {
	initial := 10
	f := func(a int) int {
		initial = initial + a
		return initial
	}
	return f
}

func main() {
	//anonymous functions
	anonymous := func() {
		fmt.Println("Hello")
	}

	anonymous()

	func(name string) {
		fmt.Println("Hello ", name)
	}("Peter")

	//user defined functions
	var word = func(a string, b string) string{
		final := a + b
		fmt.Println(final)
		return final
	}
	word("hi", "john")

	//higher order functions
	hof := func(num1 int, num2 int) int {
		return num1 + num2
	}
	handleNum(hof, 100, 2)

	function := subtractNum()
	fmt.Println(function(30, 10))

	//closures
	a := 5
	func() {
		fmt.Println(a + 10)
	}()

	f1 := increaseNum()
	fmt.Println(f1(100))   //bound is 10
	fmt.Println(f1(13))    //bound is 110 (10 + 100)

}
