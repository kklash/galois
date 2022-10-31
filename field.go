package galois

import "fmt"

// Generator is the primitive generator polynomial element used to generate all
// other elements in a finite field, by modular exponentiation of itself.
const Generator Polynomial = 2

// Field is a finite field of polynomials over an irreducible prime polynomial.
type Field[T ~uint8 | ~uint16 | ~uint32 | ~uint64] struct {
	// Prime should be an irreducible polynomial. All operations within the Field
	// are taken modulo this polynomial - That is to say, polynomials are divided
	// by this polynomial and the remainder is used as the final output.
	Prime Polynomial
}

// NewField creates a Field generated by the given prime polynomial.
//
// The generic type parameter T selects which size of unsigned integer will be
// used to represent elements of the Field. This parameter only affects the
// API of Field; Internally Field uses uint64 to represent polynomials during
// mathematical operations.
//
// Panics if the type parameter T is not of sufficient size to represent every
// element in the field.
func NewField[T ~uint8 | ~uint16 | ~uint32 | ~uint64](prime Polynomial) *Field[T] {
	max := T(0) - 1
	if fieldOrder(prime)-1 > uint64(max) {
		panic(fmt.Sprintf("cannot use %T to represent elements of GF(2^%d)", max, prime.Degree()))
	}

	return &Field[T]{Prime: prime}
}

// Order returns the order of the field (i.e. the number of elements, including zero).
func (field *Field[T]) Order() uint64 {
	return fieldOrder(field.Prime)
}

// Generate constructs a polynomial element in a finite field for the given
// prime Polynomial by exponentiating the base Generator element.
//
// If the exponent is larger than the field order, the exponent is reduced modulo
// that order.
func (field *Field[T]) Generate(exponent uint64) T {
	exponent %= field.Order()
	return T(Generator.Exp(exponent, field.Prime))
}

// Add computes the sum of the given elements within the finite field.
//
// Note that addition and subtraction within a Field is the same, because
// addition consists of adding polynomials with coefficients modulo two.
func (field *Field[T]) Add(values ...T) (sum T) {
	for _, v := range values {
		sum ^= v
	}
	return
}

// Mul multiplies a set of field field elements and returns the product. If called with no
// parameters, Mul returns zero.
//
// If any element is the additive identity element 0, Mul always returns zero.
func (field *Field[T]) Mul(values ...T) (product T) {
	if len(values) == 0 {
		return 0
	}
	for _, v := range values {
		if v == 0 {
			return 0
		}
	}
	p := Polynomial(values[0])
	for _, v := range values[1:] {
		p = p.Mul(Polynomial(v)).Mod(field.Prime)
	}
	return T(p)
}

// MultInverse computes the multiplicative inverse of y within the finite field, using the
// extended euclidean algorithm:
//
//	https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm#Simple_algebraic_field_extensions
//
// Multiplying an element with its multiplicative inverse always returns the multiplicative
// identity element 1. Every element in the field has an inverse element, including the
// multiplicative identity element whose inverse is itself, with the notable exception of the
// additive identity element zero. Zero has no inverse because multiplying it by anything
// always results in zero.
//
// Panics if y is the additive identity element zero.
func (field *Field[T]) MultInverse(y T) T {
	if y == 0 {
		panic("division by zero error")
	}

	t := Polynomial(0)
	r := field.Prime

	newt := Polynomial(1)
	newr := Polynomial(y)

	for newr != 0 {
		quotient, remainder := r.Div(newr)
		r, newr = newr, remainder
		t, newt = newt, t.Add(quotient.Mul(newt))
	}

	if r.Degree() > 0 {
		panic(
			fmt.Sprintf("failed to find inverse of GF(2^%d) element %d", field.Prime.Degree(), y),
		)
	}

	// t is now the multiplicative inverse of y in the field. x*t is the same as x/y.
	return T(t)
}

// Div returns the division of the numerator field element by the denominator element.
// As long as the field is using a real prime polynomial as its modulus, every
// field element will always be able to divide evenly into any other element, and
// still produce a defined quotient which is also a member of the field.
//
// The exception is dividing by zero. If denominator is the additive identity element
// zero, Div will panic when trying to calculate the multiplicative inverse of
// denominator.
func (field *Field[T]) Div(numerator, denominator T) T {
	denomInverse := Polynomial(field.MultInverse(denominator))
	return T(Polynomial(numerator).Mul(denomInverse).Mod(field.Prime))
}

// Exp multiplies the base element by itself the given number of times.
//
// If exponent is zero, returns the multiplicative identity 1.
func (field *Field[T]) Exp(base T, exponent uint64) T {
	if exponent == 0 {
		return 1
	}

	exponent %= field.Order()
	residue := Polynomial(base).Exp(exponent, field.Prime)
	return T(residue)
}