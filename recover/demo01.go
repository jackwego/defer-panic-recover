package main

import "fmt"

//當我們不希望因為panic而導致整個程式crash的情況下可以使用recover
//recover只有在 defer function中有效 如果不在defer中使用 recover會直接return nil
//使用recover defer panic 可以達到其他語言的 try catch finally同樣的效果

func main() {
	NormalFunction()
	RecoverFunction()
	fmt.Println("====END====")
}

func NormalFunction() {
	fmt.Println("====START====")
	//err := recover()
	//fmt.Println(err)
	//panic("Exception happened")
}

func RecoverFunction() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recover: ", err)
		}
	}()

	panic("Exception happened")
}
