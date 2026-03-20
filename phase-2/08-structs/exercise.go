package main

import "fmt"

type Student struct {
	name string
	score int
}

func upgradeScore(s *Student, bonus int) {
	s.score += bonus
}

func main() {
	myStudent := Student{
		name: "Denra",
		score: 80,
	}

	fmt.Printf("%s's initial value is %d\n", myStudent.name, myStudent.score)

	upgradeScore(&myStudent, 20)

	fmt.Printf("%s's value after adding 20 is %d\n", myStudent.name, myStudent.score)
}