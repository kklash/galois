package galois

import "fmt"

// MultIdentity is the additive identity element in any finite field of polynomials.
// Adding AddIdentity with any other Polynomial N always returns the same polynomial N.
const AddIdentity Polynomial = 0

// MultIdentity is the multiplicative identity element in any finite field of polynomials.
// Multiplying MultIdentity with any other Polynomial N always returns the same polynomial N.
const MultIdentity Polynomial = 1

// Generator is the primitive generator polynomial element used to generate all
// other elements in a finite field, by modular exponentiation of itself.
const Generator Polynomial = 2

// FieldElement is a single element in a finite field of polynomials over an
// irreducible prime polynomial.
type FieldElement struct {
	// Modulus should be an irreducible polynomial. All operations within the field
	// are taken modulo this polynomial - That is to say, polynomials are divided
	// by this polynomial and the remainder is used as the final output.
	Modulus Polynomial

	// Polynomial is the polynomial representation of the finite field element.
	Polynomial Polynomial
}

// GenerateFieldElement constructs a polynomial element in a FiniteField for the given
// prime Polynomial by exponentiating the base Generator element.
//
// If the exponent is larger than the field order, the exponent is reduced modulo
// that order.
func GenerateFieldElement(exponent uint64, prime Polynomial) *FieldElement {
	return &FieldElement{
		Modulus:    prime,
		Polynomial: Generator.Exp(exponent%FieldOrder(prime), prime),
	}
}

// NewFieldElement constructs a polynomial element in a FiniteField for the given
// prime Polynomial.
func NewFieldElement(polynomial Polynomial, prime Polynomial) *FieldElement {
	return &FieldElement{
		Modulus:    prime,
		Polynomial: polynomial.Mod(prime),
	}
}

// String returns the stringified representation of the polynomial in standard form.
func (elem *FieldElement) String() string {
	return elem.Polynomial.String() + " mod " + elem.Modulus.String()
}

// Equal returns true if this element represents the same polynomial as the other
// FieldElement.
func (self *FieldElement) Equal(other *FieldElement) bool {
	if other.Modulus != self.Modulus {
		panic("checking equality on two elements of different galois fields")
	}

	return self.Polynomial == other.Polynomial
}

// IsMultIdentity returns true if the FieldElement represents the multiplicative identity
// element in the field.
func (elem *FieldElement) IsMultIdentity() bool {
	return elem.Polynomial == MultIdentity
}

// IsAddIdentity returns true if the FieldElement represents the additive identity element
// in the field.
func (elem *FieldElement) IsAddIdentity() bool {
	return elem.Polynomial == AddIdentity
}

// Mul multiplies this FieldElement with another polynomial in the same field.
//
// If either element is the MultIdentity element, Mul returns a copy of the other
// element.
func (self *FieldElement) Mul(other *FieldElement) *FieldElement {
	if other.Modulus != self.Modulus {
		panic("multiplying two elements of different galois fields")
	}
	if other.IsMultIdentity() {
		copy := *self
		return &copy
	} else if self.IsMultIdentity() {
		copy := *other
		return &copy
	}

	return NewFieldElement(
		self.Polynomial.Mul(other.Polynomial).Mod(self.Modulus),
		self.Modulus,
	)
}

// MultInverse computes the multiplicative inverse of the FieldElement using the extended
// euclidean algorithm:
//
//	https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm#Simple_algebraic_field_extensions
//
// Multiplying an element with its multiplicative inverse always returns the MultIdentity
// element 1. Every element in the field has an inverse element, including the MultIdentity
// element whose inverse is itself, with the notable exception of the AddIdentity element
// zero. Zero has no inverse because multiplying it by anything always results in zero.
//
// Panics if the FieldElement is the AddIdentity element zero.
func (elem *FieldElement) MultInverse() *FieldElement {
	if elem.IsAddIdentity() {
		panic("cannot take inverse of the additive identity element zero")
	}

	t := Polynomial(0)
	r := elem.Modulus

	newt := Polynomial(1)
	newr := elem.Polynomial

	for newr != 0 {
		quotient, remainder := r.Div(newr)
		r, newr = newr, remainder
		t, newt = newt, t.Add(quotient.Mul(newt))
	}

	if r.Degree() > 0 {
		panic(fmt.Sprintf("failed to find inverse of element %s", elem))
	}

	return NewFieldElement(t, elem.Modulus)
}

// Div returns the division of the numerator FieldElement by the denominator element.
// As long as the FieldElement is using a real prime polynomial as its modulus, every
// FieldElement will always be able to divide any other and still produce a defined
// quotient result which is also a member of the field.
//
// The exception is dividing by zero. If denominator is the AddIdentity element zero,
// Div will panic when trying to calculate the multiplicative inverse of denominator.
func (numerator *FieldElement) Div(denominator *FieldElement) *FieldElement {
	return numerator.Mul(denominator.MultInverse())
}

// Exp multiplies the FieldElement by itself the given number of times.
//
// If exponent is zero, returns the multiplicative MultIdentity for the field.
func (elem *FieldElement) Exp(exponent uint64) *FieldElement {
	if exponent == 0 {
		return NewFieldElement(MultIdentity, elem.Modulus)
	}

	exponent %= FieldOrder(elem.Modulus)
	residue := elem.Polynomial.Exp(exponent, elem.Modulus)
	return NewFieldElement(residue, elem.Modulus)
}
