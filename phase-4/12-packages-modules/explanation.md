### Packages & Modules

So far, we’ve always written `package main`. In real-world applications, projects consist of dozens or even hundreds of packages (such as `auth`, `database`, `models`).

Please create a new folder:
`mkdir 12-packages-modules && cd 12-packages-modules`

#### What is `go mod`?

`go mod init` is a command used to start a Go project. It will create a `go.mod` file that records all the libraries (dependencies) you use.

#### Golden Rules of Packages in Go:

- **Capital Letter (Exported)**: If a function/struct starts with a capital letter (e.g., `Calculate()`), it can be accessed from other packages.
- **Lowercase Letter (Unexported)**: If it starts with a lowercase letter (e.g., `calculate()`), it is private and can only be used within the same package.
