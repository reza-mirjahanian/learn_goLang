### **Anonymous Functions in Golang**

#### **Definition**
- Anonymous functions are functions without a name.
- They are often used for short, one-off tasks or as closures.

---

### **Key Characteristics**
1. **No Name**:  
   - Unlike regular functions, anonymous functions do not have an identifier.
   - Example:
     ```go
     func(x int) int {
         return x * x
     }
     ```

2. **Can Be Declared Inline**:  
   - Anonymous functions can be defined and invoked immediately.
   - Example:
     ```go
     result := func(x int) int {
         return x + 10
     }(5) // Invoked with argument 5
     fmt.Println(result) // Output: 15
     ```

3. **Closures**:  
   - Anonymous functions can capture variables from their surrounding scope.
   - Example:
     ```go
     func main() {
         x := 10
         increment := func() int {
             x++
             return x
         }
         fmt.Println(increment()) // Output: 11
         fmt.Println(increment()) // Output: 12
     }
     ```

4. **Assigned to Variables**:  
   - Anonymous functions can be assigned to variables for reuse.
   - Example:
     ```go
     add := func(a, b int) int {
         return a + b
     }
     fmt.Println(add(3, 4)) // Output: 7
     ```

5. **Passed as Arguments**:  
   - Anonymous functions can be passed as arguments to other functions.
   - Example:
     ```go
     func operate(a, b int, operation func(int, int) int) int {
         return operation(a, b)
     }

     result := operate(10, 5, func(x, y int) int {
         return x * y
     })
     fmt.Println(result) // Output: 50
     ```

6. **Returned from Functions**:  
   - Anonymous functions can be returned from other functions.
   - Example:
     ```go
     func createMultiplier(factor int) func(int) int {
         return func(x int) int {
             return x * factor
         }
     }

     double := createMultiplier(2)
     fmt.Println(double(5)) // Output: 10
     ```

---

### **Use Cases**
- **Short-Lived Tasks**:  
  - Ideal for small operations that don’t need a named function.
  - Example:
    ```go
    numbers := []int{1, 2, 3, 4, 5}
    sum := 0
    for _, num := range numbers {
        func(n int) {
            sum += n
        }(num)
    }
    fmt.Println(sum) // Output: 15
    ```

- **Callbacks**:  
  - Used as callbacks in goroutines or higher-order functions.
  - Example:
    ```go
    go func(msg string) {
        fmt.Println(msg)
    }("Hello from goroutine")
    ```

- **Encapsulation**:  
  - Encapsulate logic within a specific scope.
  - Example:
    ```go
    func main() {
        var greet func()
        greet = func() {
            fmt.Println("Hello, World!")
        }
        greet()
    }
    ```

---

### **Best Practices**
1. **Avoid Overuse**:  
   - Use anonymous functions sparingly to maintain readability.

2. **Limit Scope**:  
   - Keep anonymous functions short and focused on a single task.

3. **Avoid Capturing Large Variables**:  
   - Capturing large variables in closures can lead to memory overhead.

4. **Error Handling**:  
   - Ensure proper error handling when using anonymous functions in goroutines.
   - Example:
     ```go
     go func() {
         defer func() {
             if r := recover(); r != nil {
                 fmt.Println("Recovered from panic:", r)
             }
         }()
         // Code that may panic
     }()
     ```

---

### **Common Pitfalls**
1. **Variable Capture Issues**:  
   - Captured variables can lead to unexpected behavior in loops.
   - Example of a common mistake:
     ```go
     funcs := []func(){}
     for i := 0; i < 3; i++ {
         funcs = append(funcs, func() {
             fmt.Println(i)
         })
     }
     for _, f := range funcs {
         f() // Output: 3, 3, 3 (not 0, 1, 2)
     }
     ```
   - Fix by creating a new variable inside the loop:
     ```go
     funcs := []func(){}
     for i := 0; i < 3; i++ {
         val := i // Create a new variable
         funcs = append(funcs, func() {
             fmt.Println(val)
         })
     }
     for _, f := range funcs {
         f() // Output: 0, 1, 2
     }
     ```

2. **Goroutine Leaks**:  
   - Anonymous functions running in goroutines may cause resource leaks if not properly managed.

---

### **Comparison Table**

| Feature                     | Named Function       | Anonymous Function       |
|-----------------------------|----------------------|--------------------------|
| **Name**                    | Has a name           | No name                  |
| **Reusability**             | High                 | Limited (unless assigned)|
| **Scope**                   | Global or package    | Local or inline          |
| **Closure Support**         | No                   | Yes                      |
| **Passing as Argument**     | Requires reference   | Directly usable          |
| **Returning from Function** | Requires reference   | Directly usable          |

---

### **Code Examples**

#### **Immediate Invocation**
```go
result := func(x int) int {
    return x * x
}(5)
fmt.Println(result) // Output: 25
```

#### **Closure with State**
```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c()) // Output: 1
fmt.Println(c()) // Output: 2
```

#### **Anonymous Function in Goroutine**
```go
go func(msg string) {
    fmt.Println(msg)
}("Hello from goroutine")
```

#### **Higher-Order Function**
```go
func apply(nums []int, transform func(int) int) []int {
    result := []int{}
    for _, num := range nums {
        result = append(result, transform(num))
    }
    return result
}

nums := []int{1, 2, 3, 4}
squared := apply(nums, func(x int) int {
    return x * x
})
fmt.Println(squared) // Output: [1, 4, 9, 16]
```

---
