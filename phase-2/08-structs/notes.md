# Module 08 Notes: Structs

-   **Definition**: A struct is a collection of fields with different data types grouped under a single name.
-   **Initialization**:
    -   `s := Student{"Name", 80}` (Order must match)
    -   `s := Student{Name: "Name", Score: 80}` (Safer & recommended)
-   **Struct + Pointer**:
    -   Used so functions can modify the original data inside the struct.
    -   Go supports _Automatic Dereferencing_ (`s.Field` works automatically even if `s` is a pointer).

## Why Are Structs Important?

Structs are the foundation for building well-organized applications. Compared to maps, structs are more rigid (fixed data types) but much safer from typing errors (_type-safe_).
