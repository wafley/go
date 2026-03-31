# Module 17 Notes: JSON & Web APIs

## Core Concepts

- **Marshaling**: The process of converting Go structs into JSON (`json.NewEncoder`).
- **ServeMux**: Go’s built-in routing system used to direct URLs to handler functions.
- **Header**: It is crucial to set `Content-Type: application/json` so the client knows how to read the data.

## Advanced Architecture

- Using **Channels** for logging ensures the API remains fast even under heavy logging workloads.
- Separating data logic, handlers, and the main function makes the code easier to maintain.
