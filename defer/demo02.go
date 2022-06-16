package main

import "fmt"

func main() {
	fmt.Println(WithReturnValue())
	//fmt.Println(WithNamedReturnValue())
	//DeferExecute()
}

func WithReturnValue() int {
	var a int
	defer func() {
		a++
	}()
	return a
}

func WithNamedReturnValue() (a int) {
	defer func() {
		a++
	}()
	return a
}

func DeferExecute() int {
	defer func() {
		fmt.Println("Defer Function")
	}()

	return SimpleFunction()
}

func SimpleFunction() int {
	fmt.Println("Simple Function")
	return 0
}
