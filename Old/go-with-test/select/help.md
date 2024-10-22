### 

`defer`

[](#defer)

By prefixing a function call with `defer` it will now call that function _at the end of the containing function_.

Sometimes you will need to clean up resources, such as closing a file or in our case closing a server so that it does not continue to listen to a port.

You want this to execute at the end of the function, but keep the instruction near where you created the server for the benefit of future readers of the code.

Our refactoring is an improvement and is a reasonable solution given the Go features covered so far, but we can make the solution simpler.

### 

Synchronising processes

[](#synchronising-processes)

-   Why are we testing the speeds of the websites one after another when Go is great at concurrency? We should be able to check both at the same time.
    

-   We don't really care about _the exact response times_ of the requests, we just want to know which one comes back first.
    

To do this, we're going to introduce a new construct called `select` which helps us synchronise processes really easily and clearly.

    func Racer(a, b string) (winner string) {
    	select {
    	case <-ping(a):
    		return a
    	case <-ping(b):
    		return b
    	}
    }
    
    func ping(url string) chan struct{} {
    	ch := make(chan struct{})
    	go func() {
    		http.Get(url)
    		close(ch)
    	}()
    	return ch
    }

We have defined a function `ping` which creates a `chan struct{}` and returns it.

In our case, we don't _care_ what type is sent to the channel, _we just want to signal we are done_ and closing the channel works perfectly!

**Why** `struct{}` and not another type like a `bool`? Well, a `chan struct{}` is the smallest data type available from a memory perspective so we get no allocation versus a `bool`. Since we are closing and not sending anything on the chan, why allocate anything?

### 

Timeouts

[](#timeouts)

Our final requirement was to return an error if `Racer` takes longer than 10 seconds

    func Racer(a, b string) (winner string, error error) {
    	select {
    	case <-ping(a):
    		return a, nil
    	case <-ping(b):
    		return b, nil
    	case <-time.After(10 * time.Second):
    		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
    	}
    }

--
## 

Wrapping up

[](#wrapping-up)

### 

`select`

[](#select-1)

-   Helps you wait on multiple channels.
    

-   Sometimes you'll want to include `time.After` in one of your `cases` to prevent your system blocking forever.
    

### 

`httptest`

[](#httptest)

-   A convenient way of creating test servers so you can have reliable and controllable tests.
    

-   Uses the same interfaces as the "real" `net/http` servers which is consistent and less for you to learn.
----

**Always** **`make`** **channels**

Notice how we have to use `make` when creating a channel; rather than say `var ch chan struct{}`. When you use `var` the variable will be initialised with the "zero" value of the type. So for `string` it is `""`, `int` it is 0, etc.

For channels the zero value is `nil` and if you try and send to it with `<-` it will block forever because you cannot send to `nil` channels


In Go, the `make` function is used to initialize channels, slices, and maps because these data types require runtime initialization. Unlike the `new` function, which only allocates memory for a type and returns a pointer to it, `make` not only allocates memory but also initializes the memory, returning an initialized (non-zeroed) value of the specified type. This distinction is crucial for slices, maps, and channels, where initialization involves setting up data structures and internal pointers, ensuring that the created instances are ready for use [1][4].

For example, when creating a channel, using `make` allows you to specify the buffer size for the channel, which can be particularly useful for controlling the flow of data and managing concurrency in your program. Here's how you can create a channel with a buffer size:

```go
c := make(chan int, 10) // Creates a channel of integers with a buffer size of 10
```

-----
Here's an example of how to use the `make` function to create a channel:


`c := make(chan int)` 

This creates a channel of type `int` that can be used to send and receive integers between goroutines.If you want to create a channel with a specific capacity, you can use the `make` function with a second argument that specifies the capacity:

    c := make(chan int, 10)

This creates a channel with a capacity of 10, meaning it can hold up to 10 values before blocking new sends or receives