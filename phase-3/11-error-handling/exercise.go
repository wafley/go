package main

import (
	"errors"
	"fmt"
)

type Student struct {
	name 	string
	age 	int
}

func RegisterStudent(name string, age int) (*Student, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if age < 15 {
		return nil, fmt.Errorf("student %s is too young: %d years (min 15)", name, age)
	}

	student := &Student{
		name: name,
		age:  age,
	}

	return student, nil
}

func main() {
	students := []Student{
		{ name: "", age: 17 },
		{ name: "Budi", age: 12 },
		{ name: "Denra", age: 18 },
	}

	for _, student := range students {
		s, err := RegisterStudent(student.name, student.age)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Student:", s.name, s.age)
	}
}