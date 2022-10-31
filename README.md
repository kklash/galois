# galois

Implements finite fields of order 2^m up to 2^32 using polynomials.

## Finite Fields

[Finite fields](https://en.wikipedia.org/wiki/Finite_field) are mathematical objects used to construct finite, cyclical number systems which mirror the behaviors of normal infinite number systems in many common ways. They create domains where numbers can be added, subtracted, multiplied, divided, and exponentiated while still behaving according to the common rules of algebra, such as [commutativity](https://en.wikipedia.org/wiki/Commutativity), [associativity](https://en.wikipedia.org/wiki/Associativity), [distributivity](https://en.wikipedia.org/wiki/Distributivity), [identity elements](https://en.wikipedia.org/wiki/Multiplicative_identity), and [inverses](https://en.wikipedia.org/wiki/Multiplicative_inverse).

The most common way to construct a finite field is by performing operations on integers, and taking the result modulo some prime integer $p$. For example, the integers modulo 5 form a finite field. Integers already follow most of the simpler common rules for fields, and taking them modulo 5 allows you to find multiplicative inverses: For every value from 1 to 4, there is another number which multiplies mod 5 to produce the multiplicative identity element 1.

|Multiplication mod 5|1|2|3|4|
|--------------------|-|-|-|-|
|1|<b><i>1</i></b>|2|3|4|
|2|2|4|<b><i>1</i></b>|3|
|3|3|<b><i>1</i></b>|4|2|
|4|4|3|2|<b><i>1</i></b>|

You can now define division among the integers mod 5. For example, $4 \div 3 \equiv 3 \mod{5}$ because $3 \cdot 3 \equiv 4 \mod{5}$

## Non-Prime Order Finite Fields

The 'order' of a finite field is the number of elements in the field. For the integers mod 5, it is easy to see that it has four elements: the integers 1 through 4, and zero. This generalizes for integer finite fields generated by any prime number $p$: That field will always have order $p$.

**But what if you want a finite field with a non-prime order?**

It is common to use finite fields with orders that are powers of two. This helps us maximize the use of storage and communication bandwidth. After all, the fundamental unit of information is the binary bit. Frequently in computer systems, information is stored and transmitted in symbols which are some fixed number of bits.

If we have need of a finite number system to keep numbers bounded to within the space of such a symbol (for example, a byte), it would be best if that number system makes _full use_ of the symbol space (thus it would work with all possible byte values from 0 to 255). Since 256 is not a prime number, we can't use simple integers and modular arithmetic to construct a finite field of that order.

## Polynomials

This library uses [polynomials modulo prime polynomials](https://en.wikipedia.org/wiki/Finite_field#Non-prime_fields) (instead of integers modulo prime integers) to construct finite fields which have non-prime order. This allows you to define cyclic multiplication and division within fixed-bit-size symbols, ensuring the outputs of the operation do not overflow the maximum symbol size.

For example, this library can be used to perform byte-wise multiplication and division, where the output of every operation is also a byte, while the normal laws of algebra still apply. See [`example_test.go`](./example_test.go) to see an example.

# Usage

First, decide how many elements you need in your finite field. This library allows callers to create finite fields of size $2^m$, where $2 <= m <= 32$.

Based on this, pick an irreducible _prime polynomial_ whose degree is $m$. For example, if you need to do finite field arithmetic on 16-bit symbols, use a polynomial of degree 16. For convenience, this library comes pre-packaged with [a set of prime polynomials](./primes.go), which were sourced from https://www.partow.net/programming/polynomials/index.html. A prime polynomial of degree 16 is exported as `galois.PrimePolynomialDegree16`.

Next, instantiate a `galois.Field[T]`:

```go
field := galois.NewField[uint16](galois.PrimePolynomialDegree16)
```

The type parameter `T` on `galois.Field[T]` selects the size of unsigned integer which will represent the elements of the field in the interface of `galois.Field[T]`. Note that this doesn't affect the speed of internal operations, only the API of `galois.Field[T]`. Internally, `galois.Field[T]` always uses `uint64` to represent polynomials when performing mathematical operations.

This `field` can then be used to perform operations on `uint16` symbols:

```go
field.Mul(0x1234, 0x4567) // 0x6324
```

Note that the choice of prime polynomial matters very much for compatibility between implementations. Even if two different prime polynomials share the same degree, and thus generate finite fields of the same order, the results of arithmetic operations within their respective fields will be different.
