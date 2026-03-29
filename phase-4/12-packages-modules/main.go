package main

import (
	"fmt"
	"packages-modules/calculator"
)

func main() {
	fmt.Println(calculator.Add(10, 5))
	
	// fmt.Println(calculator.subtract(10, 5))
	//! ^ This will error: "cannot refer to unexported name calculator.subtract"
}