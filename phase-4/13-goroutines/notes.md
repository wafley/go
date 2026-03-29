# Module 13 Notes: Goroutines

A Goroutine is a lightweight thread managed by the Go runtime.

## How to Use

Simply add the `go` keyword before calling a function.

## Characteristics

- **Non-blocking**: The `main` function continues to the next line without waiting for the goroutine to finish.
- **Shared Memory**: Goroutines share the same memory space (be careful of data races!).
- **Main Lifecycle**: If `main()` finishes, all goroutines will be forcibly terminated.

## Lesson Learned

Without synchronization mechanisms, we cannot guarantee that goroutines will complete on time before the program exits.
