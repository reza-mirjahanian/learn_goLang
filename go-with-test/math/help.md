
These tests are sometimes called 'acceptance tests', sometimes called 'feature tests'. The idea is that you write a really high level test to describe what you're trying to achieve - a user clicks a button on a website, and they see a complete list of the Pokémon they've caught, for instance. When we've written that test, we can then write more tests - unit tests - that build towards a working system that will pass the acceptance test. So for our example these tests might be about rendering a webpage with a button, testing route handlers on a web server, performing database look ups, etc. All of these things will be TDD'd, and all of them will go towards making the original acceptance test pass.

## 

Parsing XML

[](#parsing-xml)

[`encoding/xml`](https://pkg.go.dev/encoding/xml) is the Go package that can handle all things to do with simple XML parsing.

The function [`xml.Unmarshal`](https://pkg.go.dev/encoding/xml#Unmarshal) takes a `[]byte` of XML data, and a pointer to a struct for it to get unmarshalled in to.