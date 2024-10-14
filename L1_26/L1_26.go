package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	s4 := "арпвблк"
	s5 := "ААпрувВ"

	fmt.Printf("%s - %t\n", s1, uniqueLetters(s1))
	fmt.Printf("%s - %t\n", s2, uniqueLetters(s2))
	fmt.Printf("%s - %t\n", s3, uniqueLetters(s3))
	fmt.Printf("%s - %t\n", s4, uniqueLetters(s4))
	fmt.Printf("%s - %t\n", s5, uniqueLetters(s5))

}

func uniqueLetters(str string) bool {
	str = strings.ToLower(str)
	var res bool
	unique := make(map[int32]struct{})
	for _, v := range str {
		if _, ok := unique[v]; ok {
			res = false
			break
		} else {
			unique[v] = struct{}{}
			res = true
		}
	}
	return res
}
