package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "sun goes down"
	fmt.Println(reverseWords(str))
	fmt.Println(reverseWords2(str))
}

func reverseWords(str string) string {
	tmp := strings.Split(str, " ")
	builder := strings.Builder{}
	for i := len(tmp) - 1; i >= 0; i-- {
		builder.WriteString(tmp[i] + " ")
	}
	return builder.String()
}

func reverseWords2(str string) string {
	tmp := strings.Fields(str)
	l, r := 0, len(tmp)-1
	for l < r {
		tmp[l], tmp[r] = tmp[r], tmp[l]
		l++
		r--
	}
	return strings.Join(tmp, " ")
}
