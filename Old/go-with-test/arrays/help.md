
Go's built-in testing toolkit features a [coverage tool](https://blog.golang.org/cover). Whilst striving for 100% coverage should not be your end goal, the coverage tool can help identify areas of your code not covered by tests. If you have been strict with TDD, it's quite likely you'll have close to 100% coverage anyway.

Try running

`go test -cover`

You should see

    PASS
    
    coverage: 100.0% of statements

-------

`reflect.DeepEqual`  is a function in Go (from the  `reflect`  package), which is used to compare two values for equality. It can handle complex types like arrays, slices, structs, maps, and so on. The comparison performed by  `DeepEqual`  is "deep" in the sense that it goes beyond just comparing types and values. For compound types, it will recursively check the value of each element.



Here are some key things to know about  `reflect.DeepEqual`:



-   It is not type-safe, meaning that you can compare values of different types without compilation errors, but the result will be  `false`.
-   For slices and maps,  `nil`  is not equal to an empty slice or map (`nil != []int{}`  or  `nil != make(map[string]int)`).
-   When comparing floating-point numbers, it treats  `NaN`  values as equal to each other, which is not the case when using the  `==`  operator.
-   Pointers are deeply equal if they are equal using the  `==`  operator or if they point to deeply equal values.



Here's an example of how you would use  `reflect.DeepEqual`:

```
import (
    "fmt"
    "reflect"
)

type MyStruct struct {
    Value int
}

func main() {
    a := []int{1, 2, 3}
    b := []int{1, 2, 3}
    c := []int{1, 2, 4}
    fmt.Println(reflect.DeepEqual(a, b)) // true
    fmt.Println(reflect.DeepEqual(a, c)) // false

    ms1 := MyStruct{Value: 5}
    ms2 := MyStruct{Value: 5}
    fmt.Println(reflect.DeepEqual(ms1, ms2)) // true

    var s1, s2 *MyStruct
    s1 = &ms1
    fmt.Println(reflect.DeepEqual(s1, s2)) // false

    s2 = &ms2
    fmt.Println(reflect.DeepEqual(s1, s2)) // true
}
```
-----


There are several ways to initialize an array in Go lang:

**1. Declaration & Initialization:**

Go

```
var myArray [5]int // Declares an array of 5 integers, initialized with zeros by default

myArray := [3]string{"apple", "banana", "cherry"} // Declares and initializes with values


```

**2. Using composite literal syntax:**

Go

```
arr := []int{1, 2, 3, 4, 5} // Initializes an array with any length

fruits := [...]string{"apple", "banana", "orange"} // Initializes an array with specific size


```

**3. Using the make function:**

Go

```
numbers := make([]int, 10) // Creates an array of 10 integers with zero values

fruits := make([]string, 0, 5) // Creates an array with initial capacity of 0 and maximum capacity of 5


```

**4. Using multi-dimensional arrays:**

Go

```
matrix := [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}


```

**Key points to remember:**

-   Arrays in Go are fixed-size data structures.
-   The `var` keyword is used to declare an array with a specific size.
-   Composite literal syntax allows for flexible declaration and initialization.
-   `make` function provides more control over initial capacity and values.
-   Multi-dimensional arrays are also possible.


-----

    func  Sum(numbers [5]int)  int  {
    
    sum :=  0
    
    for  _, number :=  range numbers {
    
    sum += number
    
    }
    
    return sum
    
    }

`range` lets you iterate over an array. On each iteration, `range` returns two values - the index and the value. We are choosing to ignore the index value by using `_` [blank identifier](https://golang.org/doc/effective_go.html#blank).



---
| Feature | Array | Slice |
| --- | --- | --- |
| **Size** | Fixed | Variable |
| **Declaration** | `[size]type` or `:= [size]value` | `[]type` or `make([]type, capacity)` |
| **Memory allocation** | Own memory space | References underlying array |
| **Function passing** | By value (copy) | By reference (pointer) |
| **Multidimensional** | Yes | No |
| **Common Use Cases** | Fixed-size data known at compile time, avoid unnecessary copying | Dynamically changing data, efficient memory usage |