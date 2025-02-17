### **Control Structures in Golang**

---

### **1. `if/else` Statements**

#### **Basic Syntax**
```go
if condition {
    // Code to execute if condition is true
} else {
    // Code to execute if condition is false
}
```

#### **Key Features**
- **No Parentheses Required**: Unlike other languages, Go does not require parentheses around the condition.
  ```go
  if x > 10 {
      fmt.Println("x is greater than 10")
  }
  ```

- **Optional `else`**:
  ```go
  if x > 10 {
      fmt.Println("x is greater than 10")
  } else {
      fmt.Println("x is less than or equal to 10")
  }
  ```

- **Initialization Statement**: You can declare and initialize variables within the `if` statement.
  ```go
  if x := 5; x > 3 {
      fmt.Println("x is greater than 3")
  }
  ```

#### **Best Practices**
- **Avoid Deep Nesting**: Use early returns or guard clauses to reduce nesting.
  ```go
  if err != nil {
      return err
  }
  ```

- **Use Meaningful Conditions**: Keep conditions simple and readable.

---

### **2. Loops**

#### **`for` Loop**
Go only has one looping construct: `for`. It can be used in multiple ways.

##### **Basic `for` Loop**
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

##### **Single Condition (While-like Loop)**
```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

##### **Infinite Loop**
```go
for {
    fmt.Println("This will run forever")
}
```

##### **Range-based Loop**
Used to iterate over arrays, slices, maps, strings, and channels.
```go
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

- **Ignore Index or Value**: Use `_` to ignore unwanted values.
  ```go
  for _, value := range numbers {
      fmt.Println(value)
  }
  ```

#### **Best Practices**
- **Prefer `range` for Iteration**: It’s safer and more idiomatic.
- **Break Early**: Use `break` to exit loops when a condition is met.
  ```go
  for i := 0; i < 10; i++ {
      if i == 5 {
          break
      }
      fmt.Println(i)
  }
  ```
- **Skip Iterations with `continue`**:
  ```go
  for i := 0; i < 5; i++ {
      if i%2 == 0 {
          continue
      }
      fmt.Println(i)
  }
  ```

---

### **3. `switch` Statements**

#### **Basic Syntax**
```go
switch variable {
case value1:
    // Code to execute if variable == value1
case value2:
    // Code to execute if variable == value2
default:
    // Code to execute if no case matches
}
```

#### **Key Features**
- **No `break` Needed**: Go automatically breaks after each case unless `fallthrough` is used.
  ```go
  switch day {
  case "Monday":
      fmt.Println("Start of the work week")
  case "Friday":
      fmt.Println("End of the work week")
  default:
      fmt.Println("Midweek days")
  }
  ```

- **Multiple Values per Case**:
  ```go
  switch num {
  case 1, 2, 3:
      fmt.Println("Small number")
  case 4, 5, 6:
      fmt.Println("Medium number")
  default:
      fmt.Println("Large number")
  }
  ```

- **Condition-Based Cases**:
  ```go
  switch {
  case x < 0:
      fmt.Println("Negative")
  case x == 0:
      fmt.Println("Zero")
  default:
      fmt.Println("Positive")
  }
  ```

- **Type Switches**: Used to determine the type of an interface.
  ```go
  var i interface{} = "hello"
  switch v := i.(type) {
  case int:
      fmt.Println("Integer:", v)
  case string:
      fmt.Println("String:", v)
  default:
      fmt.Println("Unknown type")
  }
  ```

#### **Best Practices**
- **Use `switch` for Multiple Conditions**: It improves readability compared to nested `if/else`.
- **Avoid Overusing `fallthrough`**: It can make code harder to follow.
  ```go
  switch num {
  case 1:
      fmt.Println("One")
      fallthrough
  case 2:
      fmt.Println("Two")
  }
  ```

---

### **Comparison Table**

| Feature                  | `if/else`                          | `for`                              | `switch`                           |
|--------------------------|-------------------------------------|-------------------------------------|-------------------------------------|
| **Purpose**              | Conditional branching             | Iteration                          | Multi-case branching               |
| **Parentheses Required?**| No                                 | No                                 | No                                 |
| **Initialization**       | Supported in `if` statement        | Supported in `for` loop            | Not supported                      |
| **Early Exit**           | Use `return` or `break`            | Use `break`                        | Use `break`                        |
| **Default Case**         | Not applicable                    | Not applicable                     | Supported                          |
| **Fallthrough Behavior** | Not applicable                    | Not applicable                     | Requires explicit `fallthrough`    |

---

### **Code Examples**

#### **`if/else` Example**
```go
score := 85
if score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 75 {
    fmt.Println("Grade: B")
} else {
    fmt.Println("Grade: C")
}
```

#### **`for` Loop Example**
```go
// Sum of first 10 numbers
sum := 0
for i := 1; i <= 10; i++ {
    sum += i
}
fmt.Println("Sum:", sum)
```

#### **`switch` Example**
```go
day := "Wednesday"
switch day {
case "Monday", "Tuesday":
    fmt.Println("Start of the week")
case "Wednesday", "Thursday":
    fmt.Println("Midweek")
case "Friday":
    fmt.Println("End of the week")
default:
    fmt.Println("Weekend")
}
```

#### **Type Switch Example**
```go
var value interface{} = 42
switch v := value.(type) {
case int:
    fmt.Println("It's an integer:", v)
case string:
    fmt.Println("It's a string:", v)
default:
    fmt.Println("Unknown type")
}
```

---
