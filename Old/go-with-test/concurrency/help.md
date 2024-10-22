
Go can help us to spot race conditions with its built in [_race detector_](https://blog.golang.org/race-detector). To enable this feature, run the tests with the `race` flag: `go test -race`.

### 

Channels[](#channels)

We can solve this data race by coordinating our goroutines using _channels_. Channels are a Go data structure that can both receive and send values. These operations, along with their details, allow communication between different processes.

In this case we want to think about the communication between the parent process and each of the goroutines that it makes to do the work of running the `WebsiteChecker` function with the url.

In making it faster we learned about

-   _goroutines_, the basic unit of concurrency in Go, which let us check more than one website at the same time.


-   _anonymous functions_, which we used to start each of the concurrent processes that check websites.


-   _channels_, to help organize and control the communication between the different processes, allowing us to avoid a _race condition_ bug.


-   _the race detector_ which helped us debug problems with concurrent code