# Go Syntax Notes: Print & Variables

## 1. Differences in Print Functions (`fmt`)

-   **`fmt.Print`**: Prints text exactly as it is, without adding a newline (`\n`) at the end.
-   **`fmt.Println`**: Prints text and automatically appends a newline at the end.
-   **`fmt.Printf`**: Prints formatted text using specific **formatting verbs**.

## 2. Formatting Verbs Table (Placeholders)

| Verb | Usage                  | Example Result             |
| ---- | ---------------------- | -------------------------- |
| `%s` | String                 | `"Gopher"`                 |
| `%d` | Integer (Whole Number) | `10`                       |
| `%f` | Float (Decimal)        | `3.14`                     |
| `%t` | Boolean                | `true/false`               |
| `%v` | Universal (Any value)  | (Matches variable content) |
| `%T` | Displays Data Type     | `int`, `string`, etc.      |

## 3. When to Use `var` vs. `:=`?

### **`:=` (Short Variable Declaration)**

-   Can **only** be used inside a function.
-   Preferred for more concise and readable code.

### **`var`**

-   Used for variables declared **outside a function** (global/package level).
-   Used when declaring a variable without an initial value (**zero value**).
-   Used when you want to explicitly define the data type.
