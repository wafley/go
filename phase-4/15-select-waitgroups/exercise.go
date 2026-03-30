package main

import (
	"fmt"
	"sync"
	"time"
)

func sendEmail(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Sending emails to students #%d...\n", id)
	time.Sleep(100 * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go sendEmail(i, &wg)
	}
	
	wg.Wait()
	
	fmt.Println("All emails were sent successfully!")
}