# Module 15 Notes: Select & WaitGroups

## WaitGroups (`sync.WaitGroup`)

Used to wait for a group of Goroutines to finish.

- `Add(n)`: Adds the number of goroutines to wait for.
- `Done()`: Decreases the counter (called when a task is completed).
- `Wait()`: Blocks execution until the counter reaches 0.

## Select

Used to control communication between multiple channels simultaneously (similar to switch-case, but for channels).

## Phase 4 Conclusion

Go makes concurrency easy and efficient through Goroutines, Channels, and WaitGroups.
