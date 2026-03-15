package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 { 
		return false
	}

	for d := 2; d * d <= n; d++ {
		if n % d == 0 {
			return false
		}
	}

	return true
}

func divider(a, b float64) (result float64, msg string) {
	if b == 0 {
		return 0, "Error: divisor cannot be zero"
	}

	result = a / b
    msg = "success"
	
	return
}

func main() {
	// Use the prime check function
	for i := 1; i <= 20; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}

	fmt.Print("\n") // Breakline
	
	// Use divider function
	result, msg := divider(10, 3)
	fmt.Printf("%g, %q\n", result, msg)
}