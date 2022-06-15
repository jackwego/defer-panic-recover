package main

import "fmt"

//使用defer的function參數都必須確定
func main() {
	x := 1
	y := 2
	defer Calc("AA", x, Calc("A", x, y))
	x = 10
	defer Calc("BB", x, Calc("B", x, y))
	y = 20
}

func Calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

/*
註冊順序:
	defer Calc("AA", x, Calc("A", x, y))
	defer Calc("BB", x, Calc("B", x, y))
執行順序
	defer Calc("BB", x, Calc("B", x, y))
	defer Calc("AA", x, Calc("A", x, y))
*/

/*
實際執行順序
	Calc("A", 1, 2)  return 3
	Calc("B", 10, 2)  return 12
	defer Calc("BB", 10, 12)
	defer Calc("AA", 1, 3)
*/
