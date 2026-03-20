package main

import "fmt"

func changeWithoutPointer(val int) int {
	val = 100
	return val
}

func changeWithPointer(val *int) *int {
	*val = 100
	return val
}

func main() {
	number := 5

	fmt.Println("1. Initial value:", number)

    changeWithoutPointer(number)
    fmt.Println("2. After changeWithoutPointer:", number)

    changeWithPointer(&number)
    fmt.Println("3. After changeWithPointer:", number)
}