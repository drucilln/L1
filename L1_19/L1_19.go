package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Главрыба"
	fmt.Println(reverse(str))
	fmt.Println(reverse2(str))
	fmt.Println(reverse3(str))
}

func reverse(str string) string {
	length := len(str) - 1
	tmp := make([]rune, length+1)
	for i, v := range str {
		tmp[length-i] = v
	}
	return string(tmp)
}

func reverse2(str string) string {
	builder := strings.Builder{}
	tmp := strings.Split(str, "")
	for i := len(tmp) - 1; i >= 0; i-- {
		builder.WriteString(tmp[i])
	}
	return builder.String()
}

func reverse3(str string) string {
	builder := strings.Builder{}
	tmp := []rune(str)
	for i := len(tmp) - 1; i >= 0; i-- {
		builder.WriteRune(tmp[i])
	}
	return builder.String()
}
