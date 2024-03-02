Go has two templating packages [text/template](https://pkg.go.dev/text/template) and [html/template](https://pkg.go.dev/html/template) and they share the same interface. What they both do is allow you to combine a template and some data to produce a string.

What's the difference with the HTML version?

> Package template (html/template) implements data-driven templates for generating HTML output safe against code injection. It provides the same interface as package text/template and should be used instead of text/template whenever the output is HTML.

The templating language is very similar to [Mustache](https://mustache.github.io) and allows you to dynamically generate content in a very clean fashion with a nice separation of concerns. Compared to other templating languages you may have used, it is very constrained or "logic-less" as Mustache likes to say. This is an important, **and deliberate** design decision.

### 

Not just for HTML

[](#not-just-for-html)

Remember that go has `text/template` to generate other kinds of data from a template. If you find yourself needing to transform data into some kind of structured output, the techniques laid out in this chapter can be useful.

### 

References and further material

[](#references-and-further-material)

-   [John Calhoun's 'Learn Web Development with Go'](https://www.calhoun.io/intro-to-templates-p1-contextual-encoding/) has a number of excellent articles on templating.
    

-   [Hotwire](https://hotwired.dev) - You can use these techniques to create Hotwire web applications. It has been built by Basecamp who are primarily a Ruby on Rails shop, but because it is server-side, we can use it with Go