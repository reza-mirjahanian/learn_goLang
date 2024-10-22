You may not have used [`strings.Builder`](https://golang.org/pkg/strings/#Builder) before

> A Builder is used to efficiently build a string using Write methods. It minimizes memory copying.
### 

Property based tests

[](#property-based-tests)

-   Built into the standard library
    

-   If you can think of ways to describe your domain rules in code, they are an excellent tool for giving you more confidence
    

-   Force you to think about your domain deeply
    

-   Potentially a nice complement to your test suite

Just going to write up why a value of type `int` isn't really an 'arabic numeral'. This might be me being way too precise so I'll completely understand if you tell me to f off.

A _digit_ is a character used in the representation of numbers - from the Latin for 'finger', as we usually have ten of them. In the Arabic (also called Hindu-Arabic) number system there are ten of them. These Arabic digits are:

0 1 2 3 4 5 6 7 8 9

A _numeral_ is the representation of a number using a collection of digits. An Arabic numeral is a number represented by Arabic digits in a base 10 positional number system. We say 'positional' because each digit has a different value based upon its position in the numeral. So

1337

The `1` has a value of one thousand because its the first digit in a four digit numeral.

Roman are built using a reduced number of digits (`I`, `V` etc...) mainly as values to produce the numeral. There's a bit of positional stuff but it's mostly `I` always representing 'one'.

So, given this, is `int` an 'Arabic number'? The idea of a number is not at all tied to its representation - we can see this if we ask ourselves what the correct representation of this number is:

255

11111111

two-hundred and fifty-five

FF

377

Yes, this is a trick question. They're all correct. They're the representation of the same number in the decimal, binary, English, hexadecimal and octal number systems respectively.

The representation of a number as a numeral is _independent_ of its properties as a number - and we can see this when we look at integer literals in Go:

0xFF  ==  255  // true

And how we can print integers in a format string:

n :=  255

fmt.Printf("%b %c %d %o %q %x %X %U", n, n, n, n, n, n, n, n)

// 11111111 ÿ 255 377 'ÿ' ff FF U+00FF

We can write the same integer both as a hexadecimal and an Arabic (decimal) numeral.

So when the function signature looks like `ConvertToRoman(arabic int) string` it's making a bit of an assumption about how it's being called. Because sometimes `arabic` will be written as a decimal integer literal

ConvertToRoman(255)

But it could just as well be written

ConvertToRoman(0xFF)

Really, we're not 'converting' from an Arabic numeral at all, we're 'printing' - representing - an `int` as a Roman numeral - and `int`s are not numerals, Arabic or otherwise; they're just numbers. The `ConvertToRoman` function is more like `strconv.Itoa` in that it's turning an `int` into a `string`.

But every other version of the kata doesn't care about this distinction so :shrug: