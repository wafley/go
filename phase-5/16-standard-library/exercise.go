package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func writeLog(id int, wg *sync.WaitGroup, file io.Writer) {
	defer wg.Done()

	now := time.Now().Format("02-01-2006 15:04:05")
	
	fmt.Fprintf(file, "[%s] - Sending emails to students #%d...\n", now, id)

	time.Sleep(100 * time.Millisecond)
}

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go writeLog(i, &wg, file)
	}
	
	wg.Wait()
	
	fmt.Println("All emails were sent successfully!")
}