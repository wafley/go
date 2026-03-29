# Module 11 Notes: Error Handling (Peak of Phase 3)

In Go, errors are not "exceptions" that are thrown, but rather **values** that are returned.

## Standard Pattern

`result, err := Function()`
`if err != nil { // handle error }`

## Key Points

- **Multiple Returns**: A function can return both data and its error status at the same time.
- **Nil**: If `error == nil`, it means the operation was successful.
- **Explicit**: We are required to check errors immediately after calling a function, making applications more stable.
- **Custom Error**: Use `errors.New("message")` or `fmt.Errorf("format %s", var)`.
