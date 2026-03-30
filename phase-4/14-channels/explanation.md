### Communication Pipes (Channels)

Think of a **Channel** as a pipe.

- On one end, a Goroutine **sends** data (`ch <- data`).
- On the other end, another Goroutine **receives** data (`variable := <-ch`).

**Most Important Property: Blocking**
If you try to receive data from an empty pipe, the program will **pause and wait** (block) at that line until someone sends data. That’s why we no longer need `time.Sleep` to wait for a Goroutine to finish.

---

### Example Code Explanation (`explanation.go`)

```go
package main

import "fmt"

func main() {
    // 1. Create a channel for string data
    // We use 'make' just like when creating a Slice/Map
    message := make(chan string)

    // 2. Run a Goroutine to send data
    go func() {
        message <- "Hello from the pipe!" // Data is sent to the channel
    }()

    // 3. Receive data
    // The program will WAIT here until data is sent from above
    incomingData := <-message

    fmt.Println("Data received:", incomingData)
}
```
