# Module 03 Notes: Control Flow

## 1. Conditionals (If-Else)

-   In Go, the `if` statement does not require parentheses `()` for the condition.
-   Code blocks **must** always use curly braces `{ }`.
-   **Example:**

    ```go
    if x > 0 {
        // do something
    }
    ```

## 2. Switch Case

-   **Implicit Break**: By default, Go automatically "breaks" after each case. You do not need the `break` keyword as you do in C-style languages.
-   **Order Matters**: Cases are evaluated from top to bottom. Place the most specific conditions (e.g., checking `i % 15 == 0` in FizzBuzz) at the top.
-   **Expressionless Switch**: `switch` can be used without a variable to replace long, cluttered `if-else` chains.

## 3. For Loops (Iteration)

-   `for` is the **only** keyword available for loops in Go.
-   **Three Common Patterns**:

1. **Standard**: `for i := 0; i < 10; i++ { ... }`
2. **While-style**: `for condition { ... }` (similar to `while` in other languages).
3. **Infinite Loop**: `for { ... }` (exited using a `break` statement).

## 4. Nested Logic (Nested Loops)

-   Used to handle multidimensional data structures or visual patterns (like pyramids).
-   **Caution**: Be careful with counter variables and flag management to avoid accidental infinite loops.

## 5. Prime Numbers & Spacing (Key Insights)

-   **Prime Logic**: To optimize prime number checks, verify divisors from `2` up to the `square root of n` (using the condition $d \times d \leq n$).
-   **Visual Patterns**: To create a symmetrical pyramid, use a spacing formula such as `(height - current_row)`.
