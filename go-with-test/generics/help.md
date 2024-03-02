### 

Generics are simpler than using `interface{}` in most cases

[](#generics-are-simpler-than-using-interface-in-most-cases)

If you're inexperienced with statically-typed languages, the point of generics may not be immediately obvious, but I hope the examples in this chapter have illustrated where the Go language isn't as expressive as we'd like. In particular using `interface{}` makes your code:

-   Less safe (mix apples and oranges), requires more error handling
    

-   Less expressive, `interface{}` tells you nothing about the data
    

-   More likely to rely on [reflection](/learn-go-with-tests/go-fundamentals/reflection), type-assertions etc which makes your code more difficult to work with and more error prone as it pushes checks from compile-time to runtime
    

Using statically typed languages is an act of describing constraints. If you do it well, you create code that is not only safe and simple to use but also simpler to write because the possible solution space is smaller.

Generics gives us a new way to express constraints in our code, which as demonstrated will allow us to consolidate and simplify code that was not possible until Go 1.18.

### 

Will generics turn Go into Java?

[](#will-generics-turn-go-into-java)

-   No.
    

There's a lot of [FUD (fear, uncertainty and doubt)](https://en.wikipedia.org/wiki/Fear,_uncertainty,_and_doubt) in the Go community about generics leading to nightmare abstractions and baffling code bases. This is usually caveatted with "they must be used carefully".

Whilst this is true, it's not especially useful advice because this is true of any language feature.

Not many people complain about our ability to define interfaces which, like generics is a way of describing constraints within our code. When you describe an interface you are making a design choice that _could be poor_, generics are not unique in their ability to make confusing, annoying to use code.