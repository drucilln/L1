package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v[:10]))
}
func main() {
	someFunc()
}
