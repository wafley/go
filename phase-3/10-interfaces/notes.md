# Module 10 Notes: Interfaces

An interface in Go defines a **set of behaviors** without specifying how those behaviors are implemented.

## Core Concepts

- **Implicit Implementation**: We don’t need to write the `implements` keyword. Simply define methods that match the interface contract, and the struct will automatically be considered as implementing that interface.
- **Polymorphism**: A single function (such as `CheckGraduation`) can accept multiple data types as long as they follow the same contract.
- **Duck Typing**: "If it looks like a duck and quacks like a duck, it's a duck." (If a type has the required methods, it is considered that type.)

## Tips

- Use interfaces to make your code more flexible and easier to test (_testable_).
- Interface names usually end with "-er" (such as `Reader`, `Writer`, `Evaluator`).
