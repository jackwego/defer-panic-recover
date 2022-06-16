package main

import "fmt"

func main() {
	//Go語言有一個延遲執行的設計 - defer
	//Defer一般是用於釋放資源或者收尾工作 類似JAVA的finally區塊
	//Defer function 是在return之前觸發
	//若有多個Defer 採用先進後出的方式執行

	WithoutDefer()
	//WithDefer()
	//WithMultipleDefer()
	//DeferAnonymousFunction()
}

func WithoutDefer() {
	fmt.Println("=====START=====")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("=====END=====")
}

func WithDefer() {
	fmt.Println("=====START=====")
	fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("=====END=====")
}

func WithMultipleDefer() {
	fmt.Println("=====START=====")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("=====END=====")
}

func DeferAnonymousFunction() {
	fmt.Println("=====START=====")
	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()
	fmt.Println("=====END=====")
}
