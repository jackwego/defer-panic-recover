package main

import "fmt"

//Go語言中沒有錯誤處理的機制
//可以使用panic和recover模式來處理錯誤
//panic可以在任何地方觸發
//主要是在無法恢復的錯誤或是返回error也沒有意義的情況下使用
//panic發生的時候 會將已經註冊好的defer function執行完 在退出程序
func main() {
	fmt.Println("====START====")
	NormalFunction()
	//PanicFunction()
	PanicFunctionWithDefer()
	fmt.Println("===END====")
}

func NormalFunction() {
	fmt.Println("Normal Function")
}

func PanicFunction() {
	panic("Exception happened")
}

func PanicFunctionWithDefer() {
	defer func() {
		fmt.Println("Panic with Defer Function")
	}()
	panic("Exception happened")
}
