package repository

import (
	"database/sql"
	"school/internal/model"
)

func GetAllStudents(db *sql.DB) ([]model.Student, error) {
    rows, err := db.Query("SELECT id, name, score FROM students")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var students []model.Student
    for rows.Next() {
        var s model.Student
        if err := rows.Scan(&s.ID, &s.Name, &s.Score); err != nil {
            return nil, err
        }
        students = append(students, s)
    }
    return students, nil
}