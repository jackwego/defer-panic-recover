package main

import "fmt"

func main() {
	fmt.Println(Example1())
	fmt.Println(Example2())
	fmt.Println(Example3())
	fmt.Println(Example4())
}

func Example1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func Example2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func Example3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func Example4() (x int) {
	defer func(x int) {
		//fmt.Println(x)
		x++
	}(x)
	return 5
}
