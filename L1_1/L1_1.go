package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
}

func (a Action) speak() {
	fmt.Printf("Hi, I`m %s, I`m %d years old", a.Name, a.Age)
}

func main() {
	person := Action{Human{"James", 20}}
	person.speak()
}
