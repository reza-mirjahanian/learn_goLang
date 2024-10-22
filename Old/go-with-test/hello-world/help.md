

    go run hello.go

    go test

What's the problem? In a word, [modules](https://blog.golang.org/go116-module-changes). Luckily, the problem is easy to fix. Enter `go mod init hello` in your terminal. That will create a new file with the following contents:


        go mod <command> [arguments]

The commands are:

        download    download modules to local cache
        edit        edit go.mod from tools or scripts
        graph       print module requirement graph
        init        initialize new module in current directory
        tidy        add missing and remove unused modules
        vendor      make vendored copy of dependencies
        verify      verify dependencies have expected content
        why         explain why packages or modules are needed


### 

Writing tests[](#writing-tests)

Writing a test is just like writing a function, with a few rules

-   It needs to be in a file with a name like `xxx_test.go`


-   The test function must start with the word `Test`


-   The test function takes one argument only `t *testing.T`


-   In order to use the `*testing.T` type, you need to `import "testing"`, like we did with `fmt` in the other file


For now, it's enough to know that your `t` of type `*testing.T` is your "hook" into the testing framework so you can do things like `t.Fail()` when you want to fail.

### 

Go doc[](#go-doc)

Another quality of life feature of Go is the documentation. You can launch the docs locally by running `godoc -http :8000`. If you go to [localhost:8000/pkg](http://localhost:8000/pkg) you will see all the packages installed on your system.

The vast majority of the standard library has excellent documentation with examples. Navigating to [http://localhost:8000/pkg/testing/](http://localhost:8000/pkg/testing/) would be worthwhile to see what's available to you.

If you don't have `godoc` command, then maybe you are using the newer version of Go (1.14 or later) which is [no longer including `godoc`](https://golang.org/doc/go1.14#godoc). You can manually install it with `go install golang.org/x/tools/cmd/godoc@latest`.

`t.Helper()` is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our _function call_ rather than inside our test helper. This will help other developers track down problems easier. If you still don't understand, comment it out, make a test fail and observe the test output.

A few new concepts:

-   In our function signature we have made a _named return value_  `(prefix string)`.


-   This will create a variable called `prefix` in your function.

    -   It will be assigned the "zero" value. This depends on the type, for example `int`s are 0 and for `string`s it is `""`.

        -   You can return whatever it's set to by just calling `return` rather than `return prefix`.



    -   This will display in the Go Doc for your function so it can make the intent of your code clearer.



-   `default` in the switch case will be branched to if none of the other `case` statements match.


-   The function name starts with a lowercase letter. In Go, public functions start with a capital letter and private ones start with a lowercase. We don't want the internals of our algorithm to be exposed to the world, so we made this function private.


-   Also, we can group constants in a block instead of declaring them each on their own line. It's a good idea to use a line between sets of related constants for readability.