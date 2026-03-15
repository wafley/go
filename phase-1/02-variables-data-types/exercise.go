package main

import "fmt"

func main() {
	var school string = "SMK Negeri 2 Subang"
	fmt.Printf("School Name: %s\n", school)

	students := 30
	fmt.Printf("Total Students: %d\n", students)

	const planet = "Earth"

	var isRaining bool
	fmt.Printf("Is it raining: %t\n", isRaining)

	// The hardest part
	a := 10
	b := 20
	a, b = b, a
	fmt.Print(a, "\n")
	fmt.Print(b, "\n")
}