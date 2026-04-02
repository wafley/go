package model

type Student struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
	Score	int		`json:"score"`
}

var Students = []Student{
	{
		ID: 1,
		Name: "Revan",
		Score: 88,
	},
	{
		ID: 2,
		Name: "Denra",
		Score: 70,
	},
	{
		ID: 3,
		Name: "Hanny",
		Score: 90,
	},
	{
		ID: 4,
		Name: "Rivaldi",
		Score: 68,
	},
	{
		ID: 5,
		Name: "Deden",
		Score: 72,
	},
}