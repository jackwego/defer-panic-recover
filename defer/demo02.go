package main

import "fmt"

func main() {
	fmt.Println(WithReturnValue())
	//fmt.Println(WithNamedReturnValue())
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
