## What are Test Doubles?

Test doubles are objects used during testing to mimic or replace real components in our code. They come in different flavors, each serving a specific purpose in the testing process.

### Types of Test Doubles:

1.  **Dummy Objects:**
    
    -   Dummies are placeholders. They are passed around but not actually used. For example, if a function requires an argument, you might pass a dummy object that won't affect the test.
2.  **Stubs:**
    
    -   Stubs provide specific responses to method calls. They simulate parts of the real system and help control the flow of the test. Stubs are useful when you want to avoid the complexity of a real object.
3.  **Spies:**
    
    -   Spies keep an eye on the interactions between the test code and the component being tested. They record how the component is used during the test, allowing you to verify later.
4.  **Mocks:**
    
    -   Mocks go a step further than spies. They not only observe interactions but also set expectations. You can define what calls you expect the code under test to make, and if those expectations are not met, the test fails.
5.  **Fakes:**
    
    -   Fakes are simplified implementations of a component. They behave like the real thing but in a more straightforward manner. For example, you might use a fake database that stores data in memory rather than on disk.

---------------
##  l;dr

[](#tl-dr)

-   Mocks, spies and stubs encourage you to encode assumptions of the behaviour of your dependencies ad-hocly in each test.
    

-   These assumptions are usually not validated beyond manual checking, so they threaten your test suite's usefulness.
    

-   Fakes and contracts give us a more sustainable method for creating test doubles with validated assumptions and better reuse than the alternatives.

**Test doubles** is the collective noun for the different ways you can construct dependencies that you can control for a **subject under test**  **(SUT)**, the thing you're testing. Test doubles are often a better alternative than using the real dependency as it can avoid issues like

-   Needing the internet to use an API
    

-   Avoid latency and other performance issues
    

-   Unable to exercise non-happy path cases
    

-   Decoupling your build from another team's.
    
    -   You wouldn't want to prevent deployments if an engineer in another team accidentally shipped a bug
------
In [Anti-patterns,](https://quii.gitbook.io/learn-go-with-tests/meta/anti-patterns) there are details on how using test doubles must be done carefully. Creating a messy test suite is easy if you don't use them tastefully. As a project grows though, other problems can creep in.

When you encode behaviour into test doubles, you are adding your assumptions as to how the real dependency works into the test. If there is a discrepancy between the behaviour of the double and the real dependency, or if one happens over time (e.g. the real dependency changes, which _has_ to be expected), **you may have passing tests but failing software**.

Stubs, spies and mocks, in particular, represent other challenges, mainly as a project grows. To illustrate this, I will describe a project I worked on.

Don't rely on weekly meetings or Slack threads to flesh out changes. **Codify your assumptions in contracts**. Run those contracts against the systems in your build pipelines so you get fast feedback if new information comes to light. These contracts, in conjunction with **fakes,** mean you can work independently and manage external changes sustainably.