package main

import (
	"errors"
	"fmt"
)

func main() {
	CheckFileReadable()
	fmt.Println("Continue process.....")
}

func ReadFile(fileName string) error {

	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("read file failed")
	}

	//f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	//defer func(){
	//	if f != nil {
	//		f.Close()
	//	}
	//}()
	//if err != nil {
	// fmt.Println(err)
	//}
	//defer f.Close()
}

func CheckFileReadable() {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println("Recover:", e)
			fmt.Println("Send Notification Mail")
		}
	}()

	err := ReadFile("test.go")
	if err != nil {
		panic(err)
	}
}
