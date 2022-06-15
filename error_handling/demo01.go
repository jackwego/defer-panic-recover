package main

import "fmt"

func main() {
	fmt.Println(CalculateWithoutHandling(5, 0))
	fmt.Println(CalculateWithoutHandling(5, 5))
	//fmt.Println(CalculateWithHandling(5, 0))
	//fmt.Println(CalculateWithHandling(5, 5))
	fmt.Println("====END====")
}

func CalculateWithoutHandling(a int, b int) int {
	return a / b
}

func CalculateWithHandling(a int, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recover:", err)
		}
	}()
	return a / b
}
