Yes, the `select` statement in Go is blocking by default. It waits for one of its cases to be ready to proceed. If multiple cases are ready, it chooses one at random to execute. If none of the cases are ready, `select` will block until at least one case can proceed.

Here's a simple example to illustrate how `select` works:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
```

### Explanation:

1. **Channels `ch1` and `ch2`**: Two channels are created to communicate between goroutines and the main function.

2. **Goroutines**: Two goroutines are started. Each sends a message to its respective channel after a delay. The first goroutine sleeps for 2 seconds before sending a message to `ch1`, and the second sleeps for 1 second before sending a message to `ch2`.

3. **Select Statement**: The `select` statement waits for one of the channels to receive a message. Since `ch2` receives a message first (after 1 second), the second case in the `select` statement executes first, printing "Message from channel 2". Then, it waits for the next message, which comes from `ch1` (after 2 seconds), and prints "Message from channel 1".

### Blocking Behavior:

- **Initial Block**: The `select` statement blocks until either `ch1` or `ch2` has a message ready to be received.
- **Random Selection**: If both channels were ready simultaneously, `select` would randomly choose one to execute.
- **Non-blocking with Default**: If you want a non-blocking behavior, you can add a `default` case, which executes if no other case is ready:

```go
select {
case msg1 := <-ch1:
    fmt.Println(msg1)
case msg2 := <-ch2:
    fmt.Println(msg2)
default:
    fmt.Println("No messages received, continuing...")
}
```

In this case, if neither channel is ready, the `default` case will execute, allowing the program to continue without blocking.