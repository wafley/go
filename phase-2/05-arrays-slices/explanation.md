`Explanation Code:`

```go

package main

import "fmt"

func explanation() {
    // 1. SLICE (Most popular)
    hobbies := []string{"Coding", "Debian"} // No number inside [ ] = Slice
    hobbies = append(hobbies, "Gaming")     // Automatically add data
    fmt.Println("Hobby Slice:", hobbies)

    // 2. MAP (Key-Value)
    scores := make(map[string]int)
    scores["Gopher"] = 90
    scores["Debian"] = 100
    fmt.Println("Debian Score:", scores["Debian"])

    // 3. RANGE (Elegant way to loop data)
    for index, value := range hobbies {
        fmt.Printf("Hobby #%d is %s\n", index, value)
    }
}

```
