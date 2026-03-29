## Goroutines

This is one of the reasons why Go is often called the language of the future. In other languages, running tasks concurrently can be heavy and consume a lot of memory. In Go, we have **Goroutines**.

### What is a Goroutine?

A Goroutine is a very lightweight "thread." You can run **thousands** or even **millions** of Goroutines simultaneously without crashing your Debian system.

Using it is very simple: just add the `go` keyword before a function call.

---

### Example Code Explanation

Please create a new folder:
`mkdir 13-goroutines && cd 13-goroutines`

```go
package main

import (
    "fmt"
    "time"
)

func sayHello(name string) {
    for i := 0; i < 5; i++ {
        fmt.Println("Hello", name)
        time.Sleep(100 * time.Millisecond) // Pause for a moment
    }
}

func main() {
    // Run the function asynchronously (in the background)
    go sayHello("Gopher")

    // Run the function synchronously (normal execution)
    sayHello("Debian")

    // If main() finishes, all goroutines will be terminated.
}
```

---
