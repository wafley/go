# Module 04 Notes: Basic Functions (Comprehensive)

Functions are organized, reusable blocks of code designed to perform specific tasks. In Go, functions are highly flexible and come with several unique features.

## 1. Anatomy of a Basic Function

The basic syntax for a function in Go is as follows:

```go
func functionName(parameter dataType) returnType {
    // function body
    return value
}

```

## 2. Key Features in Go

### A. Multiple Return Values

Go allows functions to return more than one value. This is a standard pattern used to return both a result and an error message simultaneously.

-   **Example**: `func divide(a, b int) (int, error)`
-   **How to call**: `result, err := divide(10, 2)`

### B. Blank Identifier (`_`)

If a function returns multiple values but you only need one of them, use the `_` (blank identifier) to ignore the unwanted values.

-   **Example**: `result, _ := divide(10, 2)` (This ignores the error).

### C. Named Return Values

You can name the variables that will be returned directly within the function signature. This makes the code more self-documenting.

```go
func getRectArea(w, h int) (area int) {
    area = w * h
    return // Implicitly returns the 'area' variable
}

```

### D. Parameters with Identical Data Types

If consecutive parameters share the same data type, you only need to declare the type once at the end.

-   **Example**: `func add(a, b, c int)` instead of `func add(a int, b int, c int)`.

## 3. Best Practices

1. **Guard Clauses**: Check for negative or error conditions at the beginning of the function and `return` immediately. This prevents "deep nesting" of `if-else` blocks.
2. **Small & Specific**: A function should ideally perform only one task (Single Responsibility Principle).
3. **Naming Conventions**: Use **camelCase** for function names. If a function starts with an uppercase letter (e.g., `IsPrime`), it is **exported** and accessible from other packages.

## 4. Logical Analysis

-   **Primality Check**: Checking divisors up to the square root ($d \times d \leq n$) is the most efficient method for small to medium-sized numbers.
-   **Separation of Concerns**: Moving logic out of the `main` function and into dedicated functions makes the code significantly easier to **unit test** in the future.
