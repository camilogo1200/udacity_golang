# Conditionals

```go
// DECISION TREE IN GO
if CONDITION {
    //runs if the CONDITION is true
} else if ANOTHER_CONDITION {
    //code runs if  ANOTHER_CONDITION is true
} else{
    //code runs in all other cases if there's not a condition to match
}

```

## example

```go

fun main(){
    var language = "Go"
    if(language == "java"){
        fmt.Println("Language is Java")
    }else if language == "C++"{
        fmt.Println("Language is C++")
    }else if language == "Rust"{
        fmt.Println("Language is Rust")
    }
    else{
        fmt.Println("Not Java, C++ or Rust.")
    }
}
```

## Comparison Operators

| Symbol | Meaning          |
|--------| ---------------- |
| ==     | equal            |
| !=     | not equal        |
| <      | less             |
| <=     | less or equal    |
| `>`    | greater          |
| `>=`    | greater or equal |

## Logical Operators

| Symbol | Meaning         |
| ------ | --------------- |
| &&     | conditional AND |
| \|\|   | conditional OR  |
| !      | NOT             |

# Switch

- break statement is not needed ( provided automatically by Go).
- Go switch case need not be constants, and the values involved need not be integers.

```go

func main(){
    fmt.Print("Go runs on ")

    switch os := runtime.GOOS; os {
        case "darwin"
        fmt.Println("OS X.")
        case "linux"
        fmt.Println("Linux.")
        default
        fmt.Printf("%s \n",os)
    }
}

```
