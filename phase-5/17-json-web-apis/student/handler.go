package student

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

var logChan chan string

func InitLogger() {
	file, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "", log.LstdFlags)

	logChan = make(chan string, 100)

	// worker goroutine
	go func() {
		for msg := range logChan {
			logger.Println(msg)
		}
	}()
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	logChan <- "METHOD=" + r.Method + " PATH=" + r.URL.Path + " FROM=" + r.RemoteAddr

    time.Sleep(200 * time.Millisecond)

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Students)
}