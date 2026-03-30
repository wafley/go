# Module 16 Notes: Standard Library & File I/O

## Time Formatting

Time formatting in Go does not use `YYYY-MM-DD`. Instead, it uses a reference time format:

- `02` (day), `01` (month), `2006` (year)
- `15` (hour), `04` (minutes), `05` (second)

## File I/O

- `os.Create("name.txt")`: Creates a new file (overwrites the existing one if it exists).
- `os.OpenFile(...)`: Opens a file with specific options (append, read, etc.).
- `fmt.Fprintf(file, ...)`: Writes formatted text directly to a file or writer.
- `defer file.Close()`: MUST be used to prevent memory leaks and file resource issues on the Debian system.
