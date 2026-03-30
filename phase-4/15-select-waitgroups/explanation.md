## Select & WaitGroups

### 1. WaitGroups: Waiting for a Group of Goroutines

Imagine you have 100 students taking an exam. You don’t want to wait for them one by one using 100 channels. You need a **WaitGroup** from the `sync` package.

**The Three WaitGroup Rules:**

1. `wg.Add(n)`: "There are `n` people entering the exam room."
2. `wg.Done()`: "I’m done, sir!" (Called inside a Goroutine).
3. `wg.Wait()`: "The teacher stands at the door and cannot leave until everyone says `Done`."

---

### 2. Select: Multi-tasking with Channels

`select` is similar to `switch`, but specifically for **Channels**. It executes the code from whichever channel receives data first. This is very useful when you are waiting for data from multiple sources (e.g., Google API and Facebook API at the same time).

```go
select {
case msg1 := <-channel1:
    fmt.Println("Received message from 1:", msg1)
case msg2 := <-channel2:
    fmt.Println("Received message from 2:", msg2)
}
```
