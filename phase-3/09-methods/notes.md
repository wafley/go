# Module Notes 09: Methods

A method is a function that is attached to a specific data type (usually a Struct).

## Syntax

`func (receiver Type) MethodName() { ... }`

## Types of Receivers

1. **Value Receiver**: `(s Student)`
   Used when we do not want to modify the original data. It is safer but can consume more memory if the struct is very large (because it is copied).

2. **Pointer Receiver**: `(s *Student)`
   Used when we want to modify the original data or for efficiency (avoiding data duplication).

## Advantages

-   Code is more organized (Encapsulation).
-   Makes it easier to implement Interfaces (covered in the next material).
-   Cleaner syntax (`object.Method()` vs `Function(object)`).
