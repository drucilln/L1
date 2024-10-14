package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})

	for _, word := range words {
		if _, ok := set[word]; !ok {
			set[word] = struct{}{}
		}
	}

	for word := range set {
		fmt.Println(word)
	}

}
