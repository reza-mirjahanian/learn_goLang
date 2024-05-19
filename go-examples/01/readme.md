


### update golang from command prompt on Windows?

admin terminal:

```go
choco upgrade golang -y
```
install
```go
choco install golang -y
```
---------------
The `go env` command in Go (Golang) is used to print Go environment information. It provides details about the environment variables used by the Go toolchain, which can be useful for understanding how your Go environment is configured and for troubleshooting issues related to your Go setup.

#### Common Go Environment Variables

-   `GOROOT`: The root directory where Go is installed.
-   `GOPATH`: The directory where Go projects and packages are stored.
-   `GOOS`: The operating system for which Go code is being compiled (e.g., `linux`, `darwin`).
-   `GOARCH`: The architecture for which Go code is being compiled (e.g., `amd64`, `386`).
-   `GOMODCACHE`: The directory where Go modules are cached.
-   `GOPROXY`: The proxy server for downloading Go modules.
-   `GOCACHE`: The directory where Go build cache is stored.


###  What is the difference between  `go build`  and  `go run`?

`go run` both compiles and runs a program; whereas `go build` just compiles it