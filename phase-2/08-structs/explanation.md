## 🚀 Module 08: Structs

### 1. Struct Declaration

For example, if we want to define a `Student`:

```go
type Student struct {
    ID       int
    Name     string
    Major    string
    IsActive bool
}
```

### 2. Code Example Explanation (`explanation.go`)

```go
package main

import "fmt"

type Student struct {
    Name  string
    Grade int
}

func main() {
    // 1. First initialization method
    s1 := Student{"Revan", 90}

    // 2. Second initialization method (Safer/more explicit)
    s2 := Student{
        Name:  "Hanny",
        Grade: 100,
    }

    fmt.Println("Student 1:", s1.Name)
    fmt.Println("Student 2:", s2.Name)
}
```
