package main

import "fmt"

func main() {
	// Guess the number
	secretNumber := 7
	if secretNumber % 2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}

	// FizzBuzz
	for i := 1; i <= 15; i++ {
		switch {
			case i % 3 == 0 && i % 5 == 0:
				fmt.Println("FizzBuzz")
			case i % 3 == 0:
				fmt.Println("Fizz")
			case i % 5 == 0:
				fmt.Println("Buzz")
			default:
				fmt.Println(i)
		}

		// if i % 3 == 0 && i % 5 == 0 {
		// 	fmt.Println("FizzBuzz")
		// } else if i % 3 == 0 {
		// 	fmt.Println("Fizz")
		// } else if i % 5 == 0 {
		// 	fmt.Println("Buzz")
		// } else {
		// 	fmt.Println(i)
		// }
	}

	// Check the month name
	month := 3
	switch month {
		case 1:
			fmt.Println("January")
		case 2:
			fmt.Println("February")
		case 3:
			fmt.Println("March")
		case 4:
			fmt.Println("April")
		case 5:
			fmt.Println("April")
		case 6:
			fmt.Println("June")
		case 7:
			fmt.Println("July")
		case 8:
			fmt.Println("August")
		case 9:
			fmt.Println("September")
		case 10:
			fmt.Println("October")
		case 11:
			fmt.Println("November")
		case 12:
			fmt.Println("December")
		default:
            fmt.Println("Unknown")
	}

	// Challenge logic
	candidate := 2
	height := 5

	// #1
	for i := 1; i <= height; i++ {
		for j := 1; j <= i; j++ {
			for {
                isPrime := true
                for d := 2; d * d <= candidate; d++ {
                    if candidate % d == 0 {
                        isPrime = false
                        break
                    }
                }

                if isPrime {
                    fmt.Printf("%d ", candidate)
                    candidate++
                    break
                }

                candidate++
            }
		}

		fmt.Print("\n")
	}
	
	fmt.Print("\n") // Breakline

	// #2
	for i := 1; i <= height; i++ {
		for spacer := 1; spacer <= height - i; spacer++ {
			fmt.Print("   ")
		}

		for j := 1; j <= i; j++ {
			for {
                isPrime := true
                for d := 2; d * d <= candidate; d++ {
                    if candidate % d == 0 {
                        isPrime = false
                        break
                    }
                }

                if isPrime {
                    fmt.Printf("%3d   ", candidate)
                    candidate++
                    break
                }

                candidate++
            }

			// fmt.Print("* ")
		}

		fmt.Print("\n")
	}
}