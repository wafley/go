### Anatomy of an HTTP Server in Go

In Go, building a web server is not about installing extra software, but about managing **Requests** (what the user asks for) and **Responses** (what we send back).

#### A. JSON Tagging (The Communication Key)

Structs in Go use **PascalCase** (capitalized first letters) so they can be exported. However, JSON standards usually use **snake_case**. We use "tags" to bridge the two.

```go
type User struct {
    FullName string `json:"full_name"` // In Go: FullName, in JSON: full_name
}
```

#### B. Handler (The Request Receiver)

A handler is a function that accepts two essential parameters:

1. `http.ResponseWriter`: The "paper" where we write the response to send back to the user.
2. `*http.Request`: The "letter" from the user containing information (URL, headers, body data).

#### C. ServeMux (The Route Guard)

`http.HandleFunc` registers routes. If a user accesses `/api`, this router directs the request to the correct function.

---

### Code Explanation

Take a look at this code. It uses simple **encapsulation** to keep the data organized.

```go
package main

import (
	"encoding/json"
	"net/http"
)

// Info: Data structure for a single entity
type Teacher struct {
	ID      int    `json:"id"`
	Subject string `json:"subject"`
}

// Info: Handler function responsible for processing requests
func teacherHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Set response format (always put this at the top)
	w.Header().Set("Content-Type", "application/json")

	// 2. Logic: Reject any method other than GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 3. Data to be sent
	data := Teacher{ID: 101, Subject: "Computer Science"}

	// 4. Encoding process: Convert struct directly into a data stream (Writer)
	// NewEncoder is more efficient than Marshal when writing directly to the response
	json.NewEncoder(w).Encode(data)
}
```
