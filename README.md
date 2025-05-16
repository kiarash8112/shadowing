# Variable Shadowing Linter

**Variable Shadowing Linter** is a static analysis tool for Go that detects cases of variable shadowing in your codebase. Shadowing can lead to subtle bugs and reduced code clarity, and this linter helps you catch those issues early.

## 🔍 What is Variable Shadowing?

Variable shadowing occurs when a variable declared within a certain scope (e.g., inside a function or block) has the same name as a variable in an outer scope. This can lead to confusion or unintentional logic errors.

### Example:
```go
package main

func example() {
    x := 10
    if true {
        x := 20 // shadows outer 'x'
        fmt.Println(x) // prints 20
    }
    fmt.Println(x) // prints 10
}
