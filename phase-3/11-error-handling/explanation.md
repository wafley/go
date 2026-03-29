### Error Philosophy in Go

In other languages, when an error occurs, the program will “throw” an exception that can stop the application if it is not caught. In Go, a function that may fail returns **two values**: the desired result and an `error`.

If the `error` is `nil`, it means everything went smoothly. If it is not `nil`, it means there is a problem that must be handled.

---

### Code Explanation Example

```go
package main

import (
    "errors"
    "fmt"
)

// This function returns two values: float64 (result) and error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        // We create a new error using the errors package
        return 0, errors.New("cannot divide by zero")
    }
    // If successful, error is set to nil
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)

    // The "if err != nil" pattern is the essence of Go programming
    if err != nil {
        fmt.Println("An error occurred:", err)
        return // Stop execution if there is an error
    }

    fmt.Println("The result is:", result)
}
```
