Two videos I would recommend you watch are:

-   Dave Farley - [How to write acceptance tests](https://www.youtube.com/watch?v=JDD5EEJgpHU)
    

-   Nat Pryce - [E2E functional tests that can run in milliseconds](https://www.youtube.com/watch?v=Fk4rCn4YLLU)
    

"Growing Object Oriented Software" (GOOS) is such an important book for many software engineers, including myself. The approach it prescribes is the one I coach engineers I work with to follow.

-   [GOOS](http://www.growing-object-oriented-software.com) - Nat Pryce & Steve Freeman
    

Finally, [Riya Dattani](https://twitter.com/dattaniriya) and I spoke about this topic in the context of BDD in our talk, [Acceptance tests, BDD and Go](https://www.youtube.com/watch?v=ZMWJCk_0WrY).

## Recap

[](#recap)

We're talking about "black-box" tests that verify your system behaves as expected from the outside, from a "**business perspective**". The tests do not have access to the innards of the system it tests; they're only concerned with **what** your system does rather than **how**.

##  Anatomy of good acceptance tests

[](#anatomy-of-good-acceptance-tests)

If we want acceptance tests that only change when we change behaviour and not implementation detail, it stands to reason that we need to separate those concerns.

###  On types of complexity

[](#on-types-of-complexity)

As software engineers, we have to deal with two kinds of complexity.

-   **Accidental complexity** is the complexity we have to deal with because we're working with computers, stuff like networks, disks, APIs, etc.
    

-   **Essential complexity** is sometimes referred to as "domain logic". It's the particular rules and truths within your domain.
    
    -   For example, "if an account owner withdraws more money than is available, they are overdrawn". This statement says nothing about computers; this statement was true before computers were even used in banks!
        
    

Essential complexity should be expressible to a non-technical person, and it's valuable to have modelled it in our "domain" code, and in our acceptance tests.

### Separation of concerns

[](#separation-of-concerns)

What Dave Farley proposed in the video earlier, and what Riya and I also discussed, is we should have the idea of **specifications**. Specifications describe the behaviour of the system we want without being coupled with accidental complexity or implementation detail.

This idea should feel reasonable to you. In production code, we frequently strive to separate concerns and decouple units of work. Would you not hesitate to introduce an `interface` to allow your `HTTP` handler to decouple it from non-HTTP concerns? Let's take this same line of thinking for our acceptance tests.

### 

Testing on steroids

[](#testing-on-steroids)

Decoupling how the specification is executed allows us to reuse it in different scenarios. We can:

#### 

Make our drivers configurable

[](#make-our-drivers-configurable)

This means you can run your ATs locally, in your staging and (ideally) production environments.

-   Too many teams engineer their systems such that acceptance tests are impossible to run locally. This introduces an intolerably slow feedback loop. Wouldn't you rather be confident your ATs will pass _before_ integrating your code? If the tests start breaking, is it acceptable that you'd be unable to reproduce the failure locally and instead, have to commit changes and cross your fingers that it'll pass 20 minutes later in a different environment?
    

-   Remember, just because your tests pass in staging doesn't mean your system will work. Dev/Prod parity is, at best, a white lie. [I test in prod](https://increment.com/testing/i-test-in-production/).
    

-   There are always differences between the environments that can affect the _behaviour_ of your system. A CDN could have some cache headers incorrectly set; a downstream service you depend on may behave differently; a configuration value may be incorrect. But wouldn't it be nice if you could run your specifications in prod to catch these problems quickly?
    

#### 

Plug in _different_ drivers to test other parts of your system

[](#plug-in-different-drivers-to-test-other-parts-of-your-system)

This flexibility allows us to test behaviours at different abstraction and architectural layers, which allows us to have more focused tests beyond black-box tests.

-   For instance, you may have a web page with an API behind it. Why not use the same specification to test both? You can use a headless web browser for the web page, and HTTP calls for the API.
    

-   Taking this idea further, ideally, we want the **code to model essential complexity** (as "domain" code) so we should also be able to use our specifications for unit tests. This will give swift feedback that the essential complexity in our system is modelled and behaves correctly.
    

### 

Acceptance tests changing for the right reasons

[](#acceptance-tests-changing-for-the-right-reasons)

With this approach, the only reason for your specifications to change is if the behaviour of the system changes, which is reasonable.

-   If your HTTP API has to change, you have one obvious place to update it, the driver.
    

-   If your markup changes, again, update the specific driver.
    

As your system grows, you'll find yourself reusing drivers for multiple tests, which again means if implementation detail changes, you only have to update one, usually obvious place.

When done right, this approach gives us flexibility in our implementation detail and stability in our specifications. Importantly, it provides a simple and obvious structure for managing change, which becomes essential as a system and its team grows.

### 

Acceptance tests as a method for software development

[](#acceptance-tests-as-a-method-for-software-development)

In our talk, Riya and I discussed acceptance tests and their relation to BDD. We talked about how starting your work by trying to _understand the problem you're trying to solve_ and expressing it as a specification helps focus your intent and is a great way to start your work.

I was first introduced to this way of working in GOOS. A while ago, I summarised the ideas on my blog. Here is an extract from my post [Why TDD](https://quii.dev/The_Why_of_TDD)

----------

TDD is focused on letting you design for the behaviour you precisely need, iteratively. When starting a new area, you must identify a key, necessary behaviour and aggressively cut scope.

Follow a "top-down" approach, starting with an acceptance test (AT) that exercises the behaviour from the outside. This will act as a north-star for your efforts. All you should be focused on is making that test pass. This test will likely be failing for a while whilst you develop enough code to make it pass.

![](https://i.imgur.com/pxTaYu4.png)

Once your AT is set up, you can break into the TDD process to drive out enough units to make the AT pass. The trick is to not worry too much about design at this point; get enough code to make the AT pass because you're still learning and exploring the problem.

Taking this first step is often more extensive than you think, setting up web servers, routing, configuration, etc., which is why keeping the scope of the work small is essential. We want to make that first positive step on our blank canvas and have it backed by a passing AT so we can continue to iterate quickly and safely.

![](https://i.imgur.com/t5y5opw.png)

As you develop, listen to your tests, and they should give you signals to help you push your design in a better direction but, again, anchored to the behaviour rather than our imagination.

Typically, your first "unit" that does the hard work to make the AT pass will grow too big to be comfortable, even for this small amount of behaviour. This is when you can start thinking about how to break the problem down and introduce new collaborators.

![](https://i.imgur.com/UYqd7Cq.png)

This is where test doubles (e.g. fakes, mocks) are handy because most of the complexity that lives internally within software doesn't usually reside in implementation detail but "between" the units and how they interact.

#### The perils of bottom-up

[](#the-perils-of-bottom-up)

This is a "top-down" approach rather than a "bottom-up". Bottom-up has its uses, but it carries an element of risk. By building "services" and code without it being integrated into your application quickly and without verifying a high-level test, **you risk wasting lots of effort on unvalidated ideas**.

This is a crucial property of the acceptance-test-driven approach, using tests to get real validation of our code.

Too many times, I've encountered engineers who have made a chunk of code, in isolation, bottom-up, they think is going to solve a job, but it:

-   Doesn't work how we want to
    

-   Does stuff we don't need
    

-   Doesn't integrate easily
    

-   Requires a ton of re-writing anyway
    

This is waste.


## This is what the **adapter** pattern caters for.

> In [software engineering](https://en.wikipedia.org/wiki/Software_engineering), the **adapter pattern** is a [software design pattern](https://en.wikipedia.org/wiki/Software_design_pattern) (also known as [wrapper](https://en.wikipedia.org/wiki/Wrapper_function), an alternative naming shared with the [decorator pattern](https://en.wikipedia.org/wiki/Decorator_pattern)) that allows the [interface](https://en.wikipedia.org/wiki/Interface_(computer_science)) of an existing [class](https://en.wikipedia.org/wiki/Class_(computer_science)) to be used as another interface.[[1]](https://en.wikipedia.org/wiki/Adapter_pattern#cite_note-HeadFirst-1) It is often used to make existing classes work with others without modifying their [source code](https://en.wikipedia.org/wiki/Source_code).

A lot of fancy words for something relatively simple, which is often the case with design patterns, which is why people tend to roll their eyes at them. The value of design patterns is not specific implementations but a language to describe specific solutions to common problems engineers face. If you have a team that has a shared vocabulary, it reduces the friction in communication.


This is an example of **accidental complexity**. Remember, accidental complexity is the complexity we have to deal with because we're working with computers, stuff like networks, disks, APIs, etc. **The essential complexity has not changed**, so we shouldn't have to change our specifications.

Many repository structures and design patterns are mainly dealing with separating types of complexity. For instance, "ports and adapters" ask that you separate your domain code from anything to do with accidental complexity; that code lives in an "adapters" folder.

### 

Making the change easy

[](#making-the-change-easy)

Sometimes, it makes sense to do some refactoring _before_ making a change.

> First make the change easy, then make the easy change

### 

gRPC

[](#grpc)

If you're unfamiliar with gRPC, I'd start by looking at the [gRPC website](https://grpc.io). Still, for this chapter, it's just another kind of adapter into our system, a way for other systems to call (**r**emote **p**rocedure **c**all) our excellent domain code.

The twist is you define a "service definition" using Protocol Buffers. You then generate server and client code from the definition. This not only works for Go but for most mainstream languages too. This means you can share a definition with other teams in your company who may not even write Go and can still do service-to-service communication smoothly.

If you haven't used gRPC before, you'll need to install a **Protocol buffer compiler** and some **Go plugins**. [The gRPC website has clear instructions on how to do this](https://grpc.io/docs/languages/go/quickstart/).

### 

Further material

[](#further-material)

-   In this example, our "DSL" is not much of a DSL; we just used interfaces to decouple our specification from the real world and allow us to express domain logic cleanly. As your system grows, this level of abstraction might become clumsy and unclear. [Read into the "Screenplay Pattern"](https://cucumber.io/blog/bdd/understanding-screenplay-(part-1)/) if you want to find more ideas as to how to structure your specifications.
    

-   For emphasis, [Growing Object-Oriented Software, Guided by Tests,](http://www.growing-object-oriented-software.com) is a classic. It demonstrates applying this "London style", "top-down" approach to writing software. Anyone who has enjoyed Learn Go with Tests should get much value from reading GOOS.
    

-   [In the example code repository](https://github.com/quii/go-specs-greet), there's more code and ideas I haven't written about here, such as multi-stage docker build, you may wish to check this out.
    
    -   In particular, _for fun_, I made a **third program**, a website with some HTML forms to `Greet` and `Curse`. The `Driver` leverages the excellent-looking [https://github.com/go-rod/rod](https://github.com/go-rod/rod) module, which allows it to work with the website with a browser, just like a user would. Looking at the git history, you can see how I started not using any templating tools "just to make it work" Then, once I passed my acceptance test, I had the freedom to do so without fear of breaking things. -->