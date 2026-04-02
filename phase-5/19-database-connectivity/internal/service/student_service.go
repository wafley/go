package service

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"school/internal/repository"
)

var logChan chan string

func InitLogger() {
	wd, _ := os.Getwd()
	root := filepath.Join(wd, "..", "..")

	logPath := filepath.Join(root, "access.log")
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "", log.LstdFlags)

	logChan = make(chan string, 100)

	go func() {
		for msg := range logChan {
			logger.Println(msg)
		}
	}()
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	logChan <- "METHOD=" + r.Method + " PATH=" + r.URL.Path + " FROM=" + r.RemoteAddr

	db := repository.ConnectDB()

	students, err := repository.GetAllStudents(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		
		return
	}

    json.NewEncoder(w).Encode(students)
}