package main

import "fmt"

func main() {
	arr := []string{"A", "B", "C", "D", "E", "F", "G"}
	i := 2

	arr1 := deleteElem1(append([]string(nil), arr...), i)
	arr2 := deleteElem2(append([]string(nil), arr...), i)
	arr3 := deleteElem3(append([]string(nil), arr...), i)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

}

func deleteElem1(arr []string, i int) []string {
	arr[i] = arr[len(arr)-1]
	//fmt.Println(arr)
	arr[len(arr)-1] = ""
	//fmt.Println(arr)
	arr = arr[:len(arr)-1]
	return arr
}

func deleteElem2(arr []string, i int) []string {
	copy(arr[i:], arr[i+1:])
	//fmt.Println(arr)
	arr[len(arr)-1] = ""
	//fmt.Println(arr)
	arr = arr[:len(arr)-1]
	return arr
}

func deleteElem3(arr []string, i int) []string {
	arr = append(arr[:i], arr[i+1:]...)
	return arr
}
