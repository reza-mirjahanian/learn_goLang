
In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private _outside the package it's defined in_.
In Go, **when you call a function or a method the arguments are**  **_copied_**.

When calling `func (w Wallet) Deposit(amount int)` the `w` is a copy of whatever we called the method from.

Without getting too computer-sciency, when you create a value - like a wallet, it is stored somewhere in memory. You can find out what the _address_ of that bit of memory with `&myVal`.

We can fix this with _pointers_. [Pointers](https://gobyexample.com/pointers) let us _point_ to some values and then let us change them. So rather than taking a copy of the whole Wallet, we instead take a pointer to that wallet so that we can change the original values within it.

    func  (w *Wallet)  Deposit(amount int)  {
    
    w.balance += amount
    
    }
    
    func  (w *Wallet)  Balance()  int  {
    
    return w.balance
    
    }

The difference is the receiver type is `*Wallet` rather than `Wallet` which you can read as "a pointer to a wallet".

Now you might wonder, why did they pass? We didn't dereference the pointer in the function, like so:

    func  (w *Wallet)  Balance()  int  {
    
    return  (*w).balance
    
    }

and seemingly addressed the object directly. In fact, the code above using `(*w)` is absolutely valid. However, the makers of Go deemed this notation cumbersome, so the language permits us to write `w.balance`, without an explicit dereference. These pointers to structs even have their own name: _struct pointers_ and they are [automatically dereferenced](https://golang.org/ref/spec#Method_values).

Technically you do not need to change `Balance` to use a pointer receiver as taking a copy of the balance is fine. However, by convention you should keep your method receiver types the same for consistency.

### 

Pointers[](#pointers)

-   Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.


-   The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).


### 

nil[](#nil)

-   Pointers can be nil


-   When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.


-   Useful for when you want to describe a value that could be missing


### 

Errors[](#errors)

-   Errors are the way to signify failure when calling a function/method.


-   By listening to our tests we concluded that checking for a string in an error would result in a flaky test. So we refactored our implementation to use a meaningful value instead and this resulted in easier to test code and concluded this would be easier for users of our API too.


-   This is not the end of the story with error handling, you can do more sophisticated things but this is just an intro. Later sections will cover more strategies.


-   [Don’t just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)