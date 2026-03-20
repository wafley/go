# Module 07 Notes: Pointers

-   **Regular Variable**: Stores a Value.
-   **Pointer**: Stores a Memory Address.
-   **Operator `&`**: Retrieves the address of a variable (example: `&name`).
-   **Operator `*`**: Accesses or modifies the value at a pointer’s address (_dereferencing_).

## Why Use Pointers?

1. **Efficiency**: Avoids copying large data (such as structs or large arrays) multiple times.
2. **Mutation**: Allows functions to modify the original value of variables passed as arguments.

## Important Notes

-   The default value (_zero value_) of a pointer is `nil`.
-   Accessing a `nil` pointer will cause a program _panic_ (crash).
