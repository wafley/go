package service

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"school/internal/model"
	"time"
)

var logChan chan string
var sleep = time.Sleep

func InitLogger() {
	file, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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
	logChan <- "METHOD=" + r.Method + " PATH=" + r.URL.Path + " FROM=" + r.RemoteAddr

    sleep(200 * time.Millisecond)

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(model.Students)
}

func CalculateGrade(score int) string {
	switch {
		case score >= 90:
			return "A"
		case score >= 80:
			return "B"
		case score >= 70:
			return "C"
		default:
			return "D"
	}
}