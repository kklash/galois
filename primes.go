package galois

// FieldOrder returns the number of unique elements in the field generated by
// repeatedly multiplying Generator modulo the given Polynomial, assuming
// the polynomial is irreducible.
func FieldOrder(prime Polynomial) uint64 {
	return (1 << prime.Degree()) - 1
}

const (
	PrimePolynomialDegree2  Polynomial = 0b000000000000000000000000000000111 // x^2 + x + 1
	PrimePolynomialDegree3  Polynomial = 0b000000000000000000000000000001011 // x^3 + x + 1
	PrimePolynomialDegree4  Polynomial = 0b000000000000000000000000000010011 // x^4 + x + 1
	PrimePolynomialDegree5  Polynomial = 0b000000000000000000000000000100101 // x^5 + x^2 + 1
	PrimePolynomialDegree6  Polynomial = 0b000000000000000000000000001000011 // x^6 + x + 1
	PrimePolynomialDegree7  Polynomial = 0b000000000000000000000000010000011 // x^7 + x + 1
	PrimePolynomialDegree8  Polynomial = 0b000000000000000000000000100011101 // x^8 + x^4 + x^3 + x^2 + 1
	PrimePolynomialDegree9  Polynomial = 0b000000000000000000000001000010001 // x^9 + x^4 + 1
	PrimePolynomialDegree10 Polynomial = 0b000000000000000000000010000001001 // x^10 + x^3 + 1
	PrimePolynomialDegree11 Polynomial = 0b000000000000000000000100000000101 // x^11 + x^2 + 1
	PrimePolynomialDegree12 Polynomial = 0b000000000000000000001000001010011 // x^12 + x^6 + x^4 + x + 1
	PrimePolynomialDegree13 Polynomial = 0b000000000000000000010000000011011 // x^13 + x^4 + x^3 + x + 1
	PrimePolynomialDegree14 Polynomial = 0b000000000000000000100000101000011 // x^14 + x^8 + x^6 + x + 1
	PrimePolynomialDegree15 Polynomial = 0b000000000000000001000000000000011 // x^15 + x + 1
	PrimePolynomialDegree16 Polynomial = 0b000000000000000010001000000001011 // x^16 + x^12 + x^3 + x + 1
	PrimePolynomialDegree17 Polynomial = 0b000000000000000100000000000001001 // x^17 + x^3 + 1
	PrimePolynomialDegree18 Polynomial = 0b000000000000001000000000010000001 // x^18 + x^7 + 1
	PrimePolynomialDegree19 Polynomial = 0b000000000000010000000000000100111 // x^19 + x^5 + x^2 + x + 1
	PrimePolynomialDegree20 Polynomial = 0b000000000000100000000000000001001 // x^20 + x^3 + 1
	PrimePolynomialDegree21 Polynomial = 0b000000000001000000000000000000101 // x^21 + x^2 + 1
	PrimePolynomialDegree22 Polynomial = 0b000000000010000000000000000000011 // x^22 + x + 1
	PrimePolynomialDegree23 Polynomial = 0b000000000100000000000000000100001 // x^23 + x^5 + 1
	PrimePolynomialDegree24 Polynomial = 0b000000001000000000000000010000111 // x^24 + x^7 + x^2 + x + 1
	PrimePolynomialDegree25 Polynomial = 0b000000010000000000000000000001001 // x^25 + x^3 + 1
	PrimePolynomialDegree26 Polynomial = 0b000000100000000000000000001000111 // x^26 + x^6 + x^2 + x + 1
	PrimePolynomialDegree27 Polynomial = 0b000001000000000000000000000100111 // x^27 + x^5 + x^2 + x + 1
	PrimePolynomialDegree28 Polynomial = 0b000010000000000000000000000001001 // x^28 + x^3 + 1
	PrimePolynomialDegree29 Polynomial = 0b000100000000000000000000000000101 // x^29 + x^2 + 1
	PrimePolynomialDegree30 Polynomial = 0b001000000100000000000000000000111 // x^30 + x^23 + x^2 + x + 1
	PrimePolynomialDegree31 Polynomial = 0b010000000000000000000000000001001 // x^31 + x^3 + 1
	PrimePolynomialDegree32 Polynomial = 0b100000000010000000000000000000111 // x^32 + x^22 + x^2 + x + 1
)
