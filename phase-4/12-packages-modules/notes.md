# Module 12 Notes: Packages & Modules

- **Go Modules**: Initialized with `go mod init [module-name]`. Used to manage dependencies.
- **Main Package**: The file where the program starts must always use `package main`.
- **Exported (Public)**: Function/struct names start with a capital letter. Accessible from other packages.
- **Unexported (Private)**: Function/struct names start with a lowercase letter. Only accessible within the same package.
- **Import Path**: Always begins with the module name from `go.mod`, followed by the folder path.
