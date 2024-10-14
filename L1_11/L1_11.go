package main

import "fmt"

func setToMap[T comparable](arr []T) map[T]struct{} {
	unique := make(map[T]struct{})
	for _, v := range arr {
		unique[v] = struct{}{}
	}
	return unique
}

func intersect[T comparable](arr []T, unique map[T]struct{}) {
	for _, v := range arr {
		if _, ok := unique[v]; ok {
			fmt.Println(v)
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}

	unique := setToMap(arr)

	intersect(arr2, unique)
}
