#### Important Operations on Maps:

-   **`make`**: The safest way to create a map to avoid a _nil panic_.
-   **`delete(map, key)`**: Built-in function to remove an element.
-   **`len(map)`**: Returns the number of key-value pairs.

#### Example Code (`explanation.go`)

```go
package main

import "fmt"

func main() {
    // 1. Initialization with initial values (Map Literal)
    inventory := map[string]int{
        "Pencil": 10,
        "Book":   5,
    }

    // 2. Add/Update data
    inventory["Eraser"] = 15
    inventory["Book"] = 7 // Update Book value

    // 3. Delete data
    delete(inventory, "Pencil")

    // 4. Check map length
    fmt.Println("Total item types:", len(inventory))

    // 5. Iteration (Remember: Map order is always random!)
    for item, qty := range inventory {
        fmt.Printf("Stock of %s: %d\n", item, qty)
    }
}
```
