##  Just enough info about Kubernetes

[](#just-enough-info-about-kubernetes)

We run our software on [Kubernetes](https://kubernetes.io/) (K8s). K8s will terminate "pods" (in practice, our software) for various reasons, and a common one is when we push new code that we want to deploy.

We are setting ourselves high standards regarding [DORA metrics](https://cloud.google.com/blog/products/devops-sre/using-the-four-keys-to-measure-your-devops-performance), so we work in a way where we deploy small, incremental improvements and features to production multiple times per day.

When k8s wishes to terminate a pod, it initiates a ["termination lifecycle"](https://cloud.google.com/blog/products/containers-kubernetes/kubernetes-best-practices-terminating-with-grace), and a part of that is sending a SIGTERM signal to our software. This is k8s telling our code:

> You need to shut yourself down, finish whatever work you're doing because after a certain "grace period", I will send `SIGKILL`, and it's lights out for you.

On `SIGKILL` any work your program might've been doing will be immediately stopped.

-----------
##  When you have grace

[](#when-you-have-grace)

What we want to do is listen for `SIGTERM`, and rather than instantly killing the server, we want to:

-   Stop listening to any more requests
    

-   Allow any in-flight requests to finish
    

-   _Then_ terminate the process
- -----
## 

Acceptance tests

[](#acceptance-tests)

If you’ve read the rest of this book, you will have mostly written "unit tests". Unit tests are a fantastic tool for enabling fearless refactoring, driving good modular design, preventing regressions, and facilitating fast feedback.

By their nature, they only test small parts of your system. Usually, unit tests alone are _not enough_ for an effective testing strategy. Remember, we want our systems to **always be shippable**. We can't rely on manual testing, so we need another kind of testing: **acceptance tests**.

### 

What are they?

[](#what-are-they)

Acceptance tests are a kind of "black-box test". They are sometimes referred to as "functional tests". They should exercise the system as a user of the system would.

The term "black-box" refers to the idea that the test code has no access to the internals of the system, it can only use its public interface and make assertions on the behaviours it observes. This means they can only test the system as a whole.

This is an advantageous trait because it means the tests exercise the system the same as a user would, it can't use any special workarounds that could make a test pass, but not actually prove what you need to prove. This is similar to the principle of preferring your unit test files to live inside a separate test package, for example, `package mypkg_test` rather than `package mypkg`.

### 

Benefits of acceptance tests

[](#benefits-of-acceptance-tests)

-   When they pass, you know your entire system behaves how you want it to.
    

-   They are more accurate, quicker, and require less effort than manual testing.
    

-   When written well, they act as accurate, verified documentation of your system. It doesn't fall into the trap of documentation that diverges from the real behaviour of the system.
    

-   No mocking! It's all real.
    

### 

Potential drawbacks vs unit tests

[](#potential-drawbacks-vs-unit-tests)

-   They are expensive to write.
    

-   They take longer to run.
    

-   They are dependent on the design of the system.
    

-   When they fail, they typically don't give you a root cause, and can be difficult to debug.
    

-   They don't give you feedback on the internal quality of your system. You could write total garbage and still make an acceptance test pass.
    

-   Not all scenarios are practical to exercise due to the black-box nature.
    

For this reason, it is foolish to only rely on acceptance tests. They do not have many of the qualities unit tests have, and a system with a large number of acceptance tests will tend to suffer in terms of maintenance costs and poor lead time.

####  Lead time?

[](#lead-time)

Lead time refers to how long it takes from a commit being merged into your main branch to it being deployed in production. This number can vary from weeks and even months for some teams to a matter of minutes. Again, at `$WORK`, we value DORA's findings and want to keep our lead time to under 10 minutes.

A balanced testing approach is required for a reliable system with excellent lead time, and this is usually described in terms of the [Test Pyramid](https://martinfowler.com/articles/practical-test-pyramid.html).

----

The nature of _how_ to write acceptance tests depends on the system you're building, but the principles stay the same. Treat your system like a "black box". If you're making a website, your tests should act like a user, so you'll want to use a headless web browser like [Selenium](https://www.selenium.dev/), to click on links, fill in forms, etc. For a RESTful API, you'll send HTTP requests using a client.
