`Explanation Code:`

```go

package main

import "fmt"

func explanation() {
    // A. VARIABLE DECLARATION
    var name string = "Gopher" // Explicit (Data type is clearly stated)
    var age = 25               // Type Inference (Go infers the data type)
    work := "Developer"        // Short Declaration (Only available inside functions)

    // B. ZERO VALUES (Default values if not initialized)
    var a int     // Default: 0
    var b float64 // Default: 0.0
    var c string  // Default: ""
    var d bool    // Default: false

    // C. MULTIPLE DECLARATION
    var x, y, z int = 1, 2, 3

    // D. CONSTANTS (Fixed values, cannot be changed)
    const Pi = 3.14

    fmt.Printf("Data: %s, %d, %s\n", name, age, work)
    fmt.Printf("Zero Values: %d, %f, %q, %t\n", a, b, c, d)
    fmt.Printf("Multiple Declaration: %d, %d, %d\n", x, y, z)
}

```
