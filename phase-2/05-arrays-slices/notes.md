# Module 05 Notes: Arrays, Slices, & Maps

In this module, I learned how to manage collections of data using three main data structures in Go.

## 1. Arrays vs Slices

-   **Array**: Has a fixed length. It is rarely used directly in web/backend application development.
-   **Slice**: An abstraction over arrays that is much more flexible.

    -   Dynamic size (can grow).
    -   Add data using the built-in function `append(slice, element)`.
    -   Declaration: `sliceName := []dataType{val1, val2}`.

## 2. Maps (Key-Value Pair)

Maps are used to store data in key-value format (like a dictionary or JSON).

-   **Declaration**: `make(map[keyType]valueType)`.
-   **Property**: Maps are unordered. When iterating, the order of data can change.
-   **Comma OK Idiom**: A safe way to check whether a key exists in a map.

    ```go
    value, ok := myMap["key"]
    if ok {
        // data found
    }
    ```

## 3. Iteration with `for range`

The most elegant way to read data from a slice or map.

-   **For Slice**: Returns `index` and `value`.
-   **For Map**: Returns `key` and `value`.
-   Use `_` (blank identifier) if one of them is not needed.

## 4. Learning Insights

-   Using `switch` inside a `range` loop is an effective way to conditionally process slice data before inserting it into a map.
-   Using `continue` in the `default` block is safer than `return` because it only skips one iteration without stopping the entire program.
