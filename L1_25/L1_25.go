package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	format := "15:04:05"
	fmt.Println(time.Now().Format(format))
	Sleep(5 * time.Second)
	fmt.Println(time.Now().Format(format))
}
