# Functions

- A function refers to a reusable block of code.
- Functions also benefits the readability of the code.
- Reduces complexity

## Signature of a function

- `func` keyword, followed by a name of the function, then any comma-separated parameters, then the data type of the function's return value (if any.)
- Insde the body of the function can be any custom logic that we want.
- can take zero or more arguments
```go
func <name> (<parameters>) <return type>{
    //Body of the function
}
```

### Example

```go
func substract(x, y int) int{
    return x-y
}
```
