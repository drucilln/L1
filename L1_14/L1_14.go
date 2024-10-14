package main

import (
	"fmt"
)

func main() {
	var a interface{} = 42
	var b interface{} = "Привет"
	var c interface{} = true
	var d interface{} = make(chan int)

	identifyType(a)
	identifyType(b)
	identifyType(c)
	identifyType(d)
}

func identifyType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Тип переменной: int")
	case string:
		fmt.Println("Тип переменной: string")
	case bool:
		fmt.Println("Тип переменной: bool")
	case chan int, chan string, chan bool:
		fmt.Println("Тип переменной: chan")
	default:
		fmt.Println("Неизвестный тип:")
	}
}
