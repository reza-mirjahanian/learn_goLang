## `fmt` Package in Go: Tips and Tricks - A Complete Reference

### Overview

The `fmt` package in Go implements formatted I/O (input/output) with functions analogous to C's `printf` and `scanf`. It is essential for displaying data to users, debugging, and interacting with input streams.

### Printing Functions

The `fmt` package offers several functions for printing output. They can be categorized by their destination (standard output, `io.Writer`, string) and formatting style (plain, line-ending, formatted).

#### Standard Output Printing

*   **`Print(a ...interface{}) (n int, err error)`**:
    *   Prints to standard output.
    *   Formats using default formats for its operands.
    *   Spaces are added between operands when neither is a string.
    *   Returns the number of bytes written and any error encountered.

    ```go
    package main

    import "fmt"

    func main() {
        fmt.Print("Hello", " ", "World!") // Output: Hello World!
        fmt.Print(10, 20)             // Output: 1020
        fmt.Print("Value:", 10)         // Output: Value:10
    }
    ```

*   **`Println(a ...interface{}) (n int, err error)`**:
    *   Prints to standard output.
    *   Formats using default formats for its operands.
    *   Spaces are **always** added between operands.
    *   **Always** appends a newline.
    *   Returns the number of bytes written and any error encountered.

    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello", "World!") // Output: Hello World!
        fmt.Println(10, 20)          // Output: 10 20
        fmt.Println("Value:", 10)      // Output: Value: 10
    }
    ```

*   **`Printf(format string, a ...interface{}) (n int, err error)`**:
    *   Prints to standard output.
    *   Formats according to a format string.
    *   The first argument is the format string containing verbs (placeholders).
    *   Subsequent arguments are inserted into the format string according to the verbs.
    *   Returns the number of bytes written and any error encountered.

    ```go
    package main

    import "fmt"

    func main() {
        name := "Go"
        age := 14
        fmt.Printf("Language: %s, Age: %d\n", name, age) // Output: Language: Go, Age: 14
    }
    ```

#### Printing to `io.Writer`

These functions are similar to `Print`, `Println`, and `Printf`, but they write to an `io.Writer` interface instead of standard output. This allows you to print to files, network connections, buffers, etc.

*   **`Fprint(w io.Writer, a ...interface{}) (n int, err error)`**: Writes to `io.Writer` using default formats.
*   **`Fprintln(w io.Writer, a ...interface{}) (n int, err error)`**: Writes to `io.Writer` using default formats and appends a newline.
*   **`Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)`**: Writes to `io.Writer` using a format string.

    ```go
    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        file, err := os.Create("output.txt")
        if err != nil {
            panic(err)
        }
        defer file.Close()

        fmt.Fprintln(file, "Writing to a file using Fprintln!")
        fmt.Fprintf(file, "Formatted output to file: Value = %d\n", 100)
    }
    ```

#### String Formatting

These functions format strings without printing them to any output. They return the formatted string.

*   **`Sprint(a ...interface{}) string`**: Returns a string formatted using default formats.
*   **`Sprintln(a ...interface{}) string`**: Returns a string formatted using default formats and appends a newline.
*   **`Sprintf(format string, a ...interface{}) string`**: Returns a string formatted according to a format string.

    ```go
    package main

    import "fmt"

    func main() {
        message := fmt.Sprintf("The answer is %d", 42)
        fmt.Println(message) // Output: The answer is 42

        greeting := fmt.Sprintln("Hello", "from", "Sprintf!")
        fmt.Print(greeting)    // Output: Hello from Sprintf!\n
    }
    ```

#### Error Formatting

*   **`Errorf(format string, a ...interface{}) error`**:
    *   Formats according to a format string and returns an `error` value.
    *   Useful for creating custom error messages.

    ```go    package main

    import (
        "fmt"
        "errors"
    )

    func divide(a, b int) (int, error) {
        if b == 0 {
            return 0, fmt.Errorf("division by zero: a=%d, b=%d", a, b)
        }
        return a / b, nil
    }

    func main() {
        result, err := divide(10, 0)
        if err != nil {
            fmt.Println("Error:", err) // Output: Error: division by zero: a=10, b=0
            if errors.Is(err, errors.New("division by zero")) { // errors.New("division by zero") won't match fmt.Errorf
               fmt.Println("This error is related to division by zero")
            }
        } else {
            fmt.Println("Result:", result)
        }
    }
    ```

### Scanning Functions

The `fmt` package also provides functions for scanning formatted input, similar to C's `scanf`.

#### Standard Input Scanning

*   **`Scan(a ...interface{}) (n int, err error)`**:
    *   Scans text read from standard input.
    *   Reads space-separated values into successive arguments.
    *   Newlines count as space.
    *   Returns the number of items successfully scanned and any error.

    ```go
    package main

    import "fmt"

    func main() {
        var name string
        var age int

        fmt.Print("Enter your name and age: ")
        n, err := fmt.Scan(&name, &age)
        if err != nil {
            fmt.Println("Error during scan:", err)
            return
        }
        fmt.Printf("Scanned %d items: Name = %s, Age = %d\n", n, name, age)
    }
    ```

*   **`Scanln(a ...interface{}) (n int, err error)`**:
    *   Similar to `Scan`, but stops scanning at a newline.
    *   The last item must be followed by a newline or EOF.

    ```go
    package main

    import "fmt"

    func main() {
        var firstName, lastName string

        fmt.Print("Enter your first and last name (on one line): ")
        n, err := fmt.Scanln(&firstName, &lastName)
        if err != nil {
            fmt.Println("Error during scan:", err)
            return
        }
        fmt.Printf("Scanned %d items: First Name = %s, Last Name = %s\n", n, firstName, lastName)
    }
    ```

*   **`Scanf(format string, a ...interface{}) (n int, err error)`**:
    *   Scans text read from standard input according to a format string.
    *   Uses format verbs to parse the input.

    ```go
    package main

    import "fmt"

    func main() {
        var name string
        var age int

        fmt.Print("Enter your name and age in format 'Name Age': ")
        n, err := fmt.Scanf("%s %d", &name, &age)
        if err != nil {
            fmt.Println("Error during scan:", err)
            return
        }
        fmt.Printf("Scanned %d items: Name = %s, Age = %d\n", n, name, age)
    }
    ```

#### Scanning from `io.Reader`

These functions are similar to `Scan`, `Scanln`, and `Scanf`, but they read from an `io.Reader` interface.

*   **`Fscan(r io.Reader, a ...interface{}) (n int, err error)`**: Reads from `io.Reader` using space-separated values.
*   **`Fscanln(r io.Reader, a ...interface{}) (n int, err error)`**: Reads from `io.Reader` stopping at a newline.
*   **`Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)`**: Reads from `io.Reader` using a format string.

    ```go
    package main

    import (
        "fmt"
        "strings"
    )

    func main() {
        reader := strings.NewReader("John 30\nJane 25")
        var name1 string
        var age1 int
        var name2 string
        var age2 int

        n, err := fmt.Fscanln(reader, &name1, &age1)
        if err != nil {
            fmt.Println("Error during Fscanln:", err)
        } else {
            fmt.Printf("Fscanln 1: Scanned %d items: Name = %s, Age = %d\n", n, name1, age1)
        }

        n, err = fmt.Fscan(reader, &name2, &age2) // Notice no ln here - reads from where the previous scan left off
        if err != nil {
            fmt.Println("Error during Fscan:", err)
        } else {
            fmt.Printf("Fscan 2: Scanned %d items: Name = %s, Age = %d\n", n, name2, age2)
        }
    }
    ```

#### Scanning from Strings

These functions scan directly from a string.

*   **`Sscan(str string, a ...interface{}) (n int, err error)`**: Scans from a string using space-separated values.
*   **`Sscanln(str string, a ...interface{}) (n int, err error)`**: Scans from a string stopping at a newline.
*   **`Sscanf(str string, format string, a ...interface{}) (n int, err error)`**: Scans from a string using a format string.

    ```go
    package main

    import "fmt"

    func main() {
        inputString := "Apple 5 2.5"
        var fruit string
        var count int
        var price float64

        n, err := fmt.Sscan(inputString, &fruit, &count, &price)
        if err != nil {
            fmt.Println("Error during Sscan:", err)
            return
        }
        fmt.Printf("Sscan: Scanned %d items: Fruit = %s, Count = %d, Price = %f\n", n, fruit, count, price)
    }
    ```

### Formatting Verbs in `Printf` and `Sprintf`

Formatting verbs are placeholders in the format string that dictate how arguments are formatted.

| Verb | Description                                                                | Example                                      |
| :---- | :------------------------------------------------------------------------- | :------------------------------------------- |
| **General** |                                                                            |                                              |
| `%v`  | Default format.                                                            | `fmt.Printf("%v", 10)` Output: `10`           |
| `%+v` | Verbose format. Adds field names for structs.                             | `fmt.Printf("%+v", struct{X int}{X:10})` Output: `{X:10}` |
| `%#v` | Go syntax representation of the value.                                     | `fmt.Printf("%#v", "hello")` Output: `"hello"`   |
| `%T`  | Type of the value.                                                         | `fmt.Printf("%T", 10)` Output: `int`          |
| `%%`  | Literal percent sign.                                                      | `fmt.Printf("%%")` Output: `%`               |
| **Boolean** |                                                                        |                                              |
| `%t`  | `true` or `false`.                                                         | `fmt.Printf("%t", true)` Output: `true`        |
| **Integer** |                                                                          |                                              |
| `%b`  | Binary representation.                                                     | `fmt.Printf("%b", 10)` Output: `1010`         |
| `%d`  | Decimal representation.                                                    | `fmt.Printf("%d", 10)` Output: `10`           |
| `%+d` | Synonym for `%d`.                                                          | `fmt.Printf("%+d", 10)` Output: `10`          |
| `%o`  | Octal representation (base 8).                                             | `fmt.Printf("%o", 10)` Output: `12`           |
| `%O`  | Octal representation with `0o` prefix. (Go 1.17+)                          | `fmt.Printf("%O", 10)` Output: `0o12`         |
| `%x`  | Hexadecimal representation (lowercase, base 16).                           | `fmt.Printf("%x", 10)` Output: `a`            |
| `%X`  | Hexadecimal representation (uppercase, base 16).                           | `fmt.Printf("%X", 10)` Output: `A`            |
| `%c`  | Character represented by the corresponding Unicode code point.              | `fmt.Printf("%c", 65)` Output: `A`            |
| `%q`  | Single-quoted character literal safely escaped.                              | `fmt.Printf("%q", '"')` Output: `'"'`         |
| `%U`  | Unicode format: `U+1234`; same as `U+%04X`.                               | `fmt.Printf("%U", 'A')` Output: `U+0041`        |
| **Floating-point** |                                                                  |                                              |
| `%f`  | Decimal notation, no exponent.                                             | `fmt.Printf("%f", 3.14159)` Output: `3.141590` |
| `%F`  | Synonym for `%f`.                                                          | `fmt.Printf("%F", 3.14159)` Output: `3.141590` |
| `%e`  | Scientific notation (lowercase 'e').                                        | `fmt.Printf("%e", 1234.56)` Output: `1.234560e+03` |
| `%E`  | Scientific notation (uppercase 'E').                                        | `fmt.Printf("%E", 1234.56)` Output: `1.234560E+03` |
| `%g`  | `%e` for large exponents, `%f` otherwise. Precision-dependent.             | `fmt.Printf("%g", 0.00001)` Output: `1e-05`  `fmt.Printf("%g", 123.45)` Output: `123.45`|
| `%G`  | `%E` for large exponents, `%F` otherwise.                                   | `fmt.Printf("%G", 0.00001)` Output: `1E-05`  `fmt.Printf("%G", 123.45)` Output: `123.45`|
| `%x`  | Hexadecimal floating-point representation.                                 | `fmt.Printf("%x", 12.34)` Output: `0x1.8ba7ae147ae14p+3`|
| `%X`  | Uppercase hexadecimal floating-point representation.                         | `fmt.Printf("%X", 12.34)` Output: `0X1.8BA7AE147AE14P+3`|
| **String and Byte Slice** |                                                          |                                              |
| `%s`  | String or slice of bytes as is.                                            | `fmt.Printf("%s", "hello")` Output: `hello`   |
| `%q`  | Double-quoted string safely escaped.                                        | `fmt.Printf("%q", "hello\nworld")` Output: `"hello\nworld"` |
| `%x`  | Hexadecimal representation (lowercase) for strings and byte slices.        | `fmt.Printf("%x", "hello")` Output: `68656c6c6f` |
| `%X`  | Hexadecimal representation (uppercase) for strings and byte slices.        | `fmt.Printf("%X", "hello")` Output: `68656C6C6F` |
| **Pointer** |                                                                          |                                              |
| `%p`  | Pointer address in hexadecimal, prefixed with `0x`.                       | `var i int; fmt.Printf("%p", &i)` Output: `0xc0000b2008` (address will vary) |

### Width and Precision

You can control the width and precision of the formatted output using modifiers placed between the `%` and the verb.

*   **Width**: Specified by a decimal number immediately preceding the verb.
    *   Minimum number of characters to output.
    *   Pads with spaces if the output is shorter.
    *   If width is `*`, the width is taken from the next operand (must be an `int`).

    ```go
    fmt.Printf("|%10d|\n", 123)     // Output: |       123| (right-aligned, padded with spaces to width 10)
    fmt.Printf("|%-10d|\n", 123)    // Output: |123       | (left-aligned)
    width := 10
    fmt.Printf("|%*d|\n", width, 123) // Output: |       123| (width taken from variable)
    ```

*   **Precision**: Specified after the width by a period `.` followed by a decimal number.
    *   For floating-point verbs (`%f`, `%e`, `%E`, `%g`, `%G`): number of digits after the decimal point.
    *   For `%g` and `%G`: maximum number of significant digits.
    *   For strings (`%s`): maximum number of characters to print.
    *   If precision is `.*`, the precision is taken from the next operand (must be an `int`).

    ```go
    fmt.Printf("%.2f\n", 3.14159)     // Output: 3.14 (2 digits after decimal point)
    fmt.Printf("%.*f\n", 2, 3.14159)  // Output: 3.14 (precision taken from variable)
    fmt.Printf("%.5s\n", "Hello World") // Output: Hello (max 5 characters from string)
    ```

*   **Width and Precision Combined**:

    ```go
    fmt.Printf("|%10.2f|\n", 3.14159)  // Output: |      3.14| (width 10, precision 2)
    ```

### Flags

Flags modify the format of the output. They are placed between the `%` and the verb (and before width/precision).

| Flag | Description                                                                   | Applicable Verbs | Example                                      |
| :---- | :---------------------------------------------------------------------------- | :---------------- | :------------------------------------------- |
| `-`   | Left-justify within the field width. (Default is right-justified).             | All               | `fmt.Printf("|%-10d|", 123)` Output: `|123       |` |
| `+`   | Always print sign for numeric types.                                          | Numeric (%d, %f, %e, etc.) | `fmt.Printf("%+d\n", 10)` Output: `+10`     `fmt.Printf("%+d\n", -10)` Output: `-10`|
| ` ` (space) | Leave a space for positive sign in numeric types, print negative sign.       | Numeric (%d, %f, %e, etc.) | `fmt.Printf("% d\n", 10)` Output: ` 10`     `fmt.Printf("% d\n", -10)` Output: `-10`|
| `#`   | Alternate format, depends on verb:                                         |                   |                                              |
|       | - `%#o`: octal with `0o` prefix                                              | Integer (%o, %O, %x, %X) | `fmt.Printf("%#o\n", 10)` Output: `0o12`      |
|       | - `%#x`, `%#X`: hexadecimal with `0x` or `0X` prefix                             | Integer (%o, %O, %x, %X) | `fmt.Printf("%#x\n", 10)` Output: `0xa`       |
|       | - `%#U`: Unicode format (U+xxxx)                                            | Integer (%U)        | `fmt.Printf("%#U\n", 'A')` Output: `U+0041`      |
|       | - `%#q`: Backquoted string (for strings and byte slices)                    | String, Byte Slice (%q) | `fmt.Printf("%#q\n", "hello")` Output: `` `hello` `` |
|       | - For other verbs, behaviour might vary or have no effect.                    |                   |                                              |
| `0`   | Pad with leading zeros instead of spaces (for numeric types).                | Integer, Floating-point (%d, %f, etc.) | `fmt.Printf("%010d\n", 123)` Output: `0000000123` |

### Error Handling

All `fmt` printing and scanning functions return an `error` as the last return value. It is important to check this error to handle potential issues during I/O operations.

*   Printing functions generally return `nil` error unless there's an issue writing to the output destination (e.g., disk full, network error for `io.Writer`).
*   Scanning functions return `io.EOF` if input ends before scanning all arguments. Other errors can indicate formatting issues or I/O errors.

Always check the error return to ensure operations were successful, especially when dealing with external resources like files or network connections.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	n, err := fmt.Fprintln(file, "This is some text.")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("Written %d bytes to file.\n", n)
}
```

### Custom Formatting with `Stringer` and `Formatter` Interfaces

Go allows you to define custom formatting for your types by implementing the `Stringer` and `Formatter` interfaces from the `fmt` package.

#### `Stringer` Interface

If a type implements the `Stringer` interface, the `%s` and `%v` verbs will use the `String() string` method to produce the string representation of the value.

```go
package main

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func main() {
	point := Point{X: 10, Y: 20}
	fmt.Printf("Point: %s\n", point) // Output: Point: (10, 20)
	fmt.Printf("Point: %v\n", point) // Output: Point: (10, 20)
	fmt.Printf("Point: %#v\n", point) // Output: Point: main.Point{X:10, Y:20} (still default %#v)
}
```

#### `Formatter` Interface

The `Formatter` interface provides more fine-grained control over formatting. Implementing `Format(f fmt.State, verb rune)` allows you to handle all formatting verbs for your type.

```go
package main

import (
	"fmt"
	"strconv"
)

type Duration int // Duration in seconds

func (d Duration) Format(f fmt.State, verb rune) {
	switch verb {
	case 's': // %s verb, format as seconds
		fmt.Fprintf(f, "%d seconds", d)
	case 'm': // %m verb (custom verb), format as minutes
		minutes := d / 60
		seconds := d % 60
		fmt.Fprintf(f, "%d minutes and %d seconds", minutes, seconds)
	default: // Default verb handling (e.g., %v, %d, etc.) - fallback to string representation
		if f.Flag('#') { // %#v
			fmt.Fprintf(f, "Duration(%d)", d)
			return
		}
		if f.Flag('+') { // %+v
			fmt.Fprintf(f, "Duration in seconds: %d", d)
			return
		}
		fmt.Fprintf(f, "%d", d) // Default %v, %d
	}
}

func main() {
	duration := Duration(150)
	fmt.Printf("Duration (seconds): %s\n", duration)   // Output: Duration (seconds): 150 seconds
	fmt.Printf("Duration (minutes): %m\n", duration)   // Output: Duration (minutes): 2 minutes and 30 seconds
	fmt.Printf("Duration (default): %v\n", duration)   // Output: Duration (default): 150
	fmt.Printf("Duration (verbose): %+v\n", duration)  // Output: Duration (verbose): Duration in seconds: 150
	fmt.Printf("Duration (Go syntax): %#v\n", duration) // Output: Duration (Go syntax): Duration(150)
	fmt.Printf("Duration (decimal): %d\n", duration)    // Output: Duration (decimal): 150 (default)
}
```

### Performance Considerations

While `fmt` is convenient, excessive string formatting, especially within loops, can be less performant compared to direct string concatenation or using a `strings.Builder`.

*   For simple string concatenation or building strings in loops, consider using `strings.Builder` for better performance.
*   Avoid complex formatting in performance-critical sections of your code if possible.
*   `fmt.Sprintf` allocates a new string each time, which can lead to memory allocation overhead if used frequently.

However, for most common use cases like logging, displaying output to users, and debugging, the performance overhead of `fmt` is usually acceptable and the readability and ease of use it provides are valuable.