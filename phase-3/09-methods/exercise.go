package main

import "fmt"

type Student struct {
	name string
	score int
}

func (s Student) Display() string {
	return fmt.Sprintf("%s's score is %d", s.name, s.score)
}

func (s *Student) BoostScore(bonus int) {
	s.score += bonus
}

func main() {
	myStudent := Student{
		name: "Revan",
		score: 85,
	}

	fmt.Println(myStudent.Display() + " (before)")
	
	myStudent.BoostScore(10)
	fmt.Println(myStudent.Display() + " (after)")
}