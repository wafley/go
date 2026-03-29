package main

import "fmt"

type StudentEvaluator interface {
	getName() string
	IsGraduated() bool
}

type RegularStudent struct {
	name 	string
	score 	int
}

func (s RegularStudent) IsGraduated() bool {
    return s.score > 75
}

func (s RegularStudent) getName() string {
	return s.name
}

type VocationalStudent struct {
	name 		string
	skillCheck	bool
}

func (s VocationalStudent) IsGraduated() bool {
    return s.skillCheck
}

func (s VocationalStudent) getName() string {
	return s.name
}

func CheckGraduation(s StudentEvaluator) {
	if s.IsGraduated() {
		fmt.Println(s.getName() + " was declared to have passed")
	} else {
		fmt.Println(s.getName() + " was declared to have failed")
	}
}

func main() {
	regularStudents := []RegularStudent{
		{name: "Andi", score: 85},
		{name: "Budi", score: 78},
		{name: "Citra", score: 92},
		{name: "Dewi", score: 88},
		{name: "Eko", score: 74},
	}

	vocationalStudents := []VocationalStudent{
		{name: "Fajar", skillCheck: true},
		{name: "Gina", skillCheck: false},
		{name: "Hadi", skillCheck: true},
		{name: "Intan", skillCheck: true},
		{name: "Joko", skillCheck: false},
	}

	for _, student := range regularStudents {
		CheckGraduation(student)
	}

	for _, student := range vocationalStudents {
		CheckGraduation(student)
	}
}