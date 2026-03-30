### Standard Library (The Gopher's Swiss Army Knife)

Go is known for its _"Batteries Included"_ philosophy. This means Go’s **standard library** is already very complete. You don’t need many external libraries to build powerful applications.

Create a new folder:

```
mkdir 16-standard-library && cd 16-standard-library
```

#### String Manipulation (`strings`)

We often need to change letter cases, trim spaces, or check whether a sentence contains a certain word.

#### Time & Duration (`time`)

Important for recording when a transaction happens or scheduling tasks.

#### File I/O (`os` & `bufio`)

Used for reading and writing files directly to your Debian system.

---

### Example Code Explanation

```go
package main

import (
    "fmt"
    "strings"
    "time"
)

func main() {
    // 1. Strings: Clean up user input
    input := "   learning golang at SMKN 2 SUBANG   "
    cleanInput := strings.TrimSpace(input)
    upperInput := strings.ToUpper(cleanInput)

    fmt.Println("Result:", upperInput)
    fmt.Println("Contains 'GOLANG'?", strings.Contains(upperInput, "GOLANG"))

    // 2. Time: Format current time
    now := time.Now()
    // Go's unique layout: 02-01-2006 (Day-Month-Year)
    fmt.Println("Current time:", now.Format("02-01-2006 15:04:05"))
}
```
