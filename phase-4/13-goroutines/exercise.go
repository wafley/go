package main

import (
	"fmt"
	"time"
)

func printNumbers(taskName string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s: %d\n", taskName, i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go printNumbers("Task A")

	printNumbers("Task B")
}