package galois

import (
	"math/bits"
	"strconv"
	"strings"
)

// Polynomial is a polynomial whose coefficients are either one or zero.
//
// The bits of this unsigned integer are used to represent the coefficients of a
// polynomial: least-significant bits are the lowest-degree coefficients,
// most-significant bits are the highest-degree coefficients.
//
// For example, the integer 361 (101101001 in base-2), would represent the polynomial:
//
//	x^8 + x^6 + x^5 + x^3 + 1
type Polynomial uint64

// String returns the string representation of the Polynomial in standard form.
func (p Polynomial) String() string {
	nTerms := bits.OnesCount64(uint64(p))
	terms := make([]string, 0, nTerms)

	for i := bits.Len64(uint64(p)) - 1; i >= 0; i-- {
		if (p>>i)&1 == 1 {
			if i == 0 {
				terms = append(terms, "1")
			} else if i == 1 {
				terms = append(terms, "x")
			} else {
				terms = append(terms, "x^"+strconv.Itoa(i))
			}
		}
	}

	return strings.Join(terms, " + ")
}

// Degree returns the degree of the highest non-zero term of the Polynomial.
func (p Polynomial) Degree() uint64 {
	if p == 0 {
		return 0
	}
	return uint64(bits.Len64(uint64(p)) - 1)
}

// Add returns the sum of the polynomials a and b, with coefficients taken
// modulo two.
func (a Polynomial) Add(b Polynomial) (sum Polynomial) {
	return a ^ b
}

// Mul returns the product of the polynomials a and b, with coefficients
// taken modulo two.
func (a Polynomial) Mul(b Polynomial) (product Polynomial) {
	aIsPowerOfX := bits.OnesCount64(uint64(a)) == 1
	bIsPowerOfX := bits.OnesCount64(uint64(b)) == 1

	// Special case for performance: If one factor is a lone power of x
	// (e.g. x, x^2, x^5, etc) then we can just left shift by the
	// degree to multiply.
	if aIsPowerOfX {
		return b << Polynomial(a.Degree())
	} else if bIsPowerOfX {
		return a << Polynomial(b.Degree())
	}

	aLen := bits.Len64(uint64(a))
	bLen := bits.Len64(uint64(b))

	maxLen := aLen
	if bLen > aLen {
		maxLen = bLen
	}

	if (aLen-1)+(bLen-1) > 63 {
		panic("overflow: cannot multiply polynomials with combined degree greater than 63")
	}

	for i := Polynomial(0); int(i) < maxLen; i++ {
		if (a>>i)&1 > 0 {
			product ^= b << i
		}
	}
	return
}

// Div divides the numerator polynomial by the given denominator polynomial and
// returns the quotient and remainder.
//
// Panics if denominator is zero.
func (numerator Polynomial) Div(denominator Polynomial) (quotient, remainder Polynomial) {
	if denominator == 0 {
		panic("divide by zero error; cannot divide polynomial by zero")
	}

	remainder = numerator
	denomDegree := int(denominator.Degree())

	var xFactor int
	for {
		xFactor = int(remainder.Degree()) - denomDegree
		if xFactor < 0 || remainder == 0 {
			return
		}
		quotient ^= 1 << Polynomial(xFactor)
		remainder ^= denominator << Polynomial(xFactor)
	}
}

// Mod divides the numerator polynomial by the given denominator polynomial and
// returns only the remainder.
//
// Panics if denominator is zero.
func (numerator Polynomial) Mod(denominator Polynomial) Polynomial {
	_, rem := numerator.Div(denominator)
	return rem
}

// Exp exponentiates the base Polynomial to the power of the given exponent,
// modulo the given modulus Polynomial, using the square & multiply algorithm.
//
// If modulus is the empty Polynomial (zero), no modular arithmetic is performed.
func (base Polynomial) Exp(exponent uint64, modulus Polynomial) Polynomial {
	if exponent == 0 {
		return MultIdentity
	}
	result := base

	exponentBitLen := int(bits.Len64(exponent))
	for i := 1; i < exponentBitLen; i++ {
		result = result.Mul(result)
		if modulus > 0 {
			result = result.Mod(modulus)
		}

		if exponent>>(exponentBitLen-i-1)&1 == 1 {
			result = result.Mul(base)
			if modulus > 0 {
				result = result.Mod(modulus)
			}
		}
	}

	return result
}
