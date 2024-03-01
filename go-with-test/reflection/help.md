golang challenge: write a function `walk(x interface{}, fn func(string))` which takes a struct `x` and calls `fn` for all strings fields found inside. difficulty level: recursively.

## 

What is `interface{}`?

[](#what-is-interface)

We have enjoyed the type-safety that Go has offered us in terms of functions that work with known types, such as `string`, `int` and our own types like `BankAccount`.

This means that we get some documentation for free and the compiler will complain if you try and pass the wrong type to a function.

You may come across scenarios though where you want to write a function where you don't know the type at compile time.

Go lets us get around this with the type `interface{}` which you can think of as just _any_ type (in fact, in Go `any` is an [alias](https://cs.opensource.google/go/go/+/master:src/builtin/builtin.go;drc=master;l=95) for `interface{}`).

So `walk(x interface{}, fn func(string))` will accept any value for `x`.

---
## 

Wrapping up

[](#wrapping-up)

-   Introduced some concepts from the `reflect` package.
    

-   Used recursion to traverse arbitrary data structures.
    

-   This only covered a small aspect of reflection. [The Go blog has an excellent post covering more details](https://blog.golang.org/laws-of-reflection).