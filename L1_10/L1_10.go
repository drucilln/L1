package main

import "fmt"

func main() {
	celsius := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10.0
	groups := make(map[float64][]float64)

	for _, temp := range celsius {
		tmp := float64(int64(temp/step)) * step
		groups[tmp] = append(groups[tmp], temp)
	}

	for group, value := range groups {
		fmt.Println(group, value)
	}
}
