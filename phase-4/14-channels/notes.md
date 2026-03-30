# Module 14 Notes: Channels

Channels are a way for Goroutines to communicate with each other.

## Main Syntax

- Create: `ch := make(chan DataType)`
- Send: `ch <- data`
- Receive: `variable := <- ch` or `<- ch` (wait only)

## Important Concepts

- **Blocking**: The sender waits until there is a receiver, and the receiver waits until there is a sender.
- **Synchronization**: Used to ensure Goroutines finish before the main program exits (replacing the hacky `time.Sleep` approach).
- **Direction**: Channels can be directional (send-only or receive-only), but by default they are bidirectional.
