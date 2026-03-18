package main

import "fmt"

func main() {
	// 1. Slice
	studentList := []string{"Revan", "Hanny", "Denra"}
	
	// 2. Add Data
	studentList = append(studentList, "Rival", "Cantika")
	fmt.Println(studentList)

	// 3. Map
	testScores := make(map[string]int)
	for _, v := range studentList {
		switch v {
			case "Revan":
				testScores[v] = 80
			case "Hanny":
				testScores[v] = 100
			case "Denra":
				testScores[v] = 90
			case "Rival":
				testScores[v] = 65
			case "Cantika":
				testScores[v] = 70
			default:
				continue
		}
	}

	// 4. Looping
	for name, score := range testScores {
		fmt.Printf("Student %s got a score %d\n", name, score)
	}

	// Experiment
	score, ok := testScores["Hanny"]
	if ok {
		fmt.Println("Hanny's Score:", score)
	} else {
		fmt.Println("Hanny's data not found!")
	}
}