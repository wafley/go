Berikut versi yang sudah diubah ke **full bahasa Inggris**:

---

### What are Methods?

A **method** is essentially a function, but it is "attached" to a struct.

-   If a **regular function** is called like this: `upgradeScore(&myStudent, 20)`
-   Then a **method** is called like this: `myStudent.UpgradeScore(20)`

This makes your code feel more natural, as if the object `myStudent` has its own behavior.

#### Two Types of Receivers:

1. **Value Receiver**: Only reads data (cannot modify the original value).
2. **Pointer Receiver**: Can modify the original struct (uses `*`).

---

### Code Example (`explanation.go`)

```go
package main

import "fmt"

type Rect struct {
    Width, Height int
}

// 1. Value Receiver (Only calculates, does not modify)
func (r Rect) Area() int {
    return r.Width * r.Height
}

// 2. Pointer Receiver (Modifies the original size)
func (r *Rect) Resize(newW, newH int) {
    r.Width = newW
    r.Height = newH
}

func main() {
    myRect := Rect{Width: 10, Height: 5}

    fmt.Println("Initial Area:", myRect.Area())

    myRect.Resize(20, 10)
    fmt.Println("Area After Resize:", myRect.Area())
}
```
