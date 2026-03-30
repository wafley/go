package main

import "fmt"

func sendReport(done chan bool) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Sedang memproses laporan ke-%d...\n", i)
	}

	done <- true
	close(done)
}

func main() {
	done := make(chan bool)

	go sendReport(done)

	<-done
}