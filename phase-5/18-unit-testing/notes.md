# Module 18 Notes: Unit Testing

## Basic Rules

- File: `test_name.go`
- Function: `func TestXxx(t *testing.T)`

## Table-Driven Tests

A method to test multiple scenarios using a slice of structs. Very effective for checking edge cases (critical values).

## Coverage

An indicator of how many lines of code are executed by tests.  
Command: `go test -cover ./...`
