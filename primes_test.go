package galois

import (
	"testing"
)

func TestPrimes(t *testing.T) {
	type TestCase struct {
		PrimePolynomial Polynomial
		String          string
		Degree          uint64
		Order           uint64
	}

	testCases := []TestCase{
		{
			PrimePolynomial: PrimePolynomialDegree2,
			String:          "x^2 + x + 1",
			Degree:          2,
			Order:           0b11,
		},
		{
			PrimePolynomial: PrimePolynomialDegree3,
			String:          "x^3 + x + 1",
			Degree:          3,
			Order:           0b111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree4,
			String:          "x^4 + x + 1",
			Degree:          4,
			Order:           0b1111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree5,
			String:          "x^5 + x^2 + 1",
			Degree:          5,
			Order:           0b11111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree6,
			String:          "x^6 + x + 1",
			Degree:          6,
			Order:           0b111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree7,
			String:          "x^7 + x + 1",
			Degree:          7,
			Order:           0b1111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree8,
			String:          "x^8 + x^4 + x^3 + x^2 + 1",
			Degree:          8,
			Order:           0b11111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree9,
			String:          "x^9 + x^4 + 1",
			Degree:          9,
			Order:           0b111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree10,
			String:          "x^10 + x^3 + 1",
			Degree:          10,
			Order:           0b1111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree11,
			String:          "x^11 + x^2 + 1",
			Degree:          11,
			Order:           0b11111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree12,
			String:          "x^12 + x^6 + x^4 + x + 1",
			Degree:          12,
			Order:           0b111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree13,
			String:          "x^13 + x^4 + x^3 + x + 1",
			Degree:          13,
			Order:           0b1111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree14,
			String:          "x^14 + x^8 + x^6 + x + 1",
			Degree:          14,
			Order:           0b11111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree15,
			String:          "x^15 + x + 1",
			Degree:          15,
			Order:           0b111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree16,
			String:          "x^16 + x^12 + x^3 + x + 1",
			Degree:          16,
			Order:           0b1111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree17,
			String:          "x^17 + x^3 + 1",
			Degree:          17,
			Order:           0b11111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree18,
			String:          "x^18 + x^7 + 1",
			Degree:          18,
			Order:           0b111111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree19,
			String:          "x^19 + x^5 + x^2 + x + 1",
			Degree:          19,
			Order:           0b1111111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree20,
			String:          "x^20 + x^3 + 1",
			Degree:          20,
			Order:           0b11111111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree21,
			String:          "x^21 + x^2 + 1",
			Degree:          21,
			Order:           0b111111111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree22,
			String:          "x^22 + x + 1",
			Degree:          22,
			Order:           0b1111111111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree23,
			String:          "x^23 + x^5 + 1",
			Degree:          23,
			Order:           0b11111111111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree24,
			String:          "x^24 + x^7 + x^2 + x + 1",
			Degree:          24,
			Order:           0b111111111111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree25,
			String:          "x^25 + x^3 + 1",
			Degree:          25,
			Order:           0b1111111111111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree26,
			String:          "x^26 + x^6 + x^2 + x + 1",
			Degree:          26,
			Order:           0b11111111111111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree27,
			String:          "x^27 + x^5 + x^2 + x + 1",
			Degree:          27,
			Order:           0b111111111111111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree28,
			String:          "x^28 + x^3 + 1",
			Degree:          28,
			Order:           0b1111111111111111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree29,
			String:          "x^29 + x^2 + 1",
			Degree:          29,
			Order:           0b11111111111111111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree30,
			String:          "x^30 + x^23 + x^2 + x + 1",
			Degree:          30,
			Order:           0b111111111111111111111111111111,
		},
		{
			PrimePolynomial: PrimePolynomialDegree31,
			String:          "x^31 + x^3 + 1",
			Degree:          31,
			Order:           0b1111111111111111111111111111111,
		},
	}

	for _, test := range testCases {
		if actual := test.PrimePolynomial.String(); actual != test.String {
			t.Errorf("polynomial did not stringify as expected; wanted %q, got %q", test.String, actual)
		}
		if degree := test.PrimePolynomial.Degree(); degree != test.Degree {
			t.Errorf("polynomial did not return expected degree; wanted %d, got %d", test.Degree, degree)
		}
		if order := FieldOrder(test.PrimePolynomial); order != test.Order {
			t.Errorf("polynomial did not return expected order; wanted %d, got %d", test.Order, order)
		}
	}
}
