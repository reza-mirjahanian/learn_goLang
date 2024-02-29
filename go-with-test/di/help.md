_Motivated by our tests_ we refactored the code so we could control _where_ the data was written by **injecting a dependency** which allowed us to:

-   **Test our code** If you can't test a function _easily_, it's usually because of dependencies hard-wired into a function _or_ global state. If you have a global database connection pool for instance that is used by some kind of service layer, it is likely going to be difficult to test and they will be slow to run. DI will motivate you to inject in a database dependency (via an interface) which you can then mock out with something you can control in your tests.


-   **Separate our concerns**, decoupling _where the data goes_ from _how to generate it_. If you ever feel like a method/function has too many responsibilities (generating data _and_ writing to a db? handling HTTP requests _and_ doing domain level logic?) DI is probably going to be the tool you need.


-   **Allow our code to be re-used in different contexts** The first "new" context our code can be used in is inside tests. But further on if someone wants to try something new with your function they can inject their own dependencies.



Dependency Injection (DI) is a design pattern that helps to reduce the hard-coded dependencies among your classes or functions, making your code more modular, testable, and maintainable. In Go, you can achieve DI by passing dependencies as parameters to functions or struct methods.

Here's a simple example of Dependency Injection in Go:

```go
package main

import (
	"fmt"
	"io"
	"os"
)

// Greeter struct
type Greeter struct {
	Writer io.Writer
}

// Greet method
func (g *Greeter) Greet(name string) {
	fmt.Fprintf(g.Writer, "Hello, %s\n", name)
}

func main() {
	// Create a new Greeter with os.Stdout as the writer
	greeter := &Greeter{
		Writer: os.Stdout,
	}

	// Call the Greet method
	greeter.Greet("World")
}
```

In this example, `Greeter` is a struct that has a dependency on something that implements `io.Writer` (an interface). Instead of hard-coding this dependency (e.g., always writing to `os.Stdout`), we inject it through the struct's field. This way, we can easily change where `Greeter` writes its output, making it more flexible and testable.

In the `main` function, we create a new `Greeter` and inject `os.Stdout` as the writer. Then we call the `Greet` method to write a greeting to the standard output. If we wanted to write the greeting somewhere else (e.g., a file or a buffer), we could do so by injecting a different `io.Writer` when creating the `Greeter`.