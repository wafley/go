### Material: What is a Pointer?

Imagine a regular variable as a **Box** that holds items.
Imagine a Pointer as a **Note** that contains the **Address** of where that box is located.

In Go, there are two important symbols you must remember:

1. **Operator `&` (Address-of)**: Used to get the **address** of a variable.
2. **Operator `*` (Star)**: Used to access or modify the **value** at that address (Dereferencing).

#### Why do we need pointers?

Without pointers, when you pass a variable into a function, Go creates a **copy** of it. If the variable is large (for example: a slice with 1 million elements), copying it repeatedly will slow down your program. With pointers, you only pass the "address" instead.

---

### Code Example Explanation (`explanation.go`)

```go
package main

import "fmt"

func main() {
    name := "Gopher"

    // 1. Get memory address (&)
    nameAddress := &name

    fmt.Println("Value of variable name:", name)         // Output: Gopher
    fmt.Println("Address of variable name:", nameAddress) // Output: 0xc000... (Hexadecimal)

    // 2. Modify value through the address (*)
    *nameAddress = "Debian 13"

    fmt.Println("Value of name after being modified via pointer:", name) // Output: Debian 13
}
```
