### What is an Interface?

If a **Struct** is about **what an object is** (it has a name, it has a score), then an **Interface** is about **what an object can do** (it can speak, it can calculate).

Interfaces in Go are unique because they are **implicit**. You don’t need to write `implements` like in Java. If a Struct has all the methods required by an Interface, then that Struct is automatically considered to implement that Interface.

#### Analogy:

Imagine a contract called `Smart`. The contract has only one rule: _"Must have a method `Learn()`"_.

- If `Student` has a `Learn()` method, it is `Smart`.
- If `Robot` has a `Learn()` method, it is also `Smart`.

---

### Example Code Explanation

```go
package main

import "fmt"

// 1. Interface definition (Contract)
type Speaker interface {
    Speak() string
}

// 2. Person struct
type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}

// 3. Cat struct
type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

// 4. Function that accepts an Interface
func Introduce(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    p := Person{Name: "Revan"}
    c := Cat{}

    Introduce(p) // Person is treated as Speaker
    Introduce(c) // Cat is also treated as Speaker
}
```
