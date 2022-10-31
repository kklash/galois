package galois

import (
	"fmt"
	"testing"
)

func TestPolonomial_Mul(t *testing.T) {
	type TestCase struct {
		A, B    Polynomial
		Product Polynomial
	}

	testCases := []TestCase{
		{
			A:       0b00000000,
			B:       0b00000000,
			Product: 0b00000000,
		},
		{
			A:       0b00000001,
			B:       0b00000000,
			Product: 0b00000000,
		},
		{
			A:       0b00000000,
			B:       0b00000001,
			Product: 0b00000000,
		},
		{
			A:       0b00000001,
			B:       0b00000001,
			Product: 0b00000001,
		},
		{
			A:       0b00000010,
			B:       0b00000001,
			Product: 0b00000010,
		},
		{
			A:       0b00000010,
			B:       0b00000010,
			Product: 0b00000100,
		},
		{
			A:       0b00000010,
			B:       0b00000010,
			Product: 0b00000100,
		},
		{
			A:       0b00000010,
			B:       0b00000010,
			Product: 0b00000100,
		},
		{
			A:       0b00001011,
			B:       0b00000101,
			Product: 0b00100111,
		},
		{
			A:       0b0000000000101010,
			B:       0b0000000001110001,
			Product: 0b0000110101001010,
		},
	}

	for _, test := range testCases {
		product := test.A.Mul(test.B)
		if product != test.Product {
			t.Errorf(
				"expected polynomial product (%s) * (%s) = %s - got %s",
				test.A, test.B, test.Product, product,
			)
		}
	}
}

func TestPolonomial_Div(t *testing.T) {
	type TestCase struct {
		Numerator   Polynomial
		Denominator Polynomial
		Quotient    Polynomial
		Remainder   Polynomial
	}

	testCases := []TestCase{
		{
			Numerator:   0b1000,
			Denominator: 0b1011,
			Quotient:    0b1,
			Remainder:   0b0011,
		},
		{
			Numerator:   0b1001011,
			Denominator: 0b0000110,
			Quotient:    0b11100,
			Remainder:   0b0000011,
		},
		{
			Numerator:   0b1010001111010,
			Denominator: 0b0000100011101,
			Quotient:    0b10101,
			Remainder:   0b0000011000011,
		},
		{
			Numerator:   0b011111111,
			Denominator: 0b100011101,
			Quotient:    0,
			Remainder:   0b011111111,
		},
	}

	for _, test := range testCases {
		quo, rem := test.Numerator.Div(test.Denominator)
		if quo != test.Quotient {
			t.Errorf(
				"expected polynomial division (%s) / (%s) = %s (got %s)",
				test.Numerator, test.Denominator, test.Quotient, quo,
			)
		}

		if rem != test.Remainder {
			t.Errorf(
				"expected polynomial remainder (%s) %% (%s) = %s (got %s)",
				test.Numerator, test.Denominator, test.Remainder, rem,
			)
		}

		product := quo.Mul(test.Denominator).Add(rem)
		if product != test.Numerator {
			t.Errorf(
				"failed polynomial division validity check\n(%s)(%s) + %s != %s",
				quo, test.Denominator, rem, test.Numerator,
			)
		}
	}
}

func TestPolynomial_Exp(t *testing.T) {
	type TestCase struct {
		Base     Polynomial
		Exponent uint64
		Modulus  Polynomial
		Expected Polynomial
	}

	testCases := []TestCase{
		{
			Base:     0b10,
			Exponent: 0,
			Expected: 0b1,
		},
		{
			Base:     0b10,
			Exponent: 1,
			Expected: 0b10,
		},
		{
			Base:     0b10,
			Exponent: 2,
			Expected: 0b100,
		},
		{
			Base:     0b10,
			Exponent: 3,
			Expected: 0b1000,
		},
		{
			Base:     0b10,
			Exponent: 3,
			Modulus:  0b1011,
			Expected: 0b11,
		},
		{
			Base:     0b10,
			Exponent: 4,
			Expected: 0b10000,
		},
		{
			Base:     0b10,
			Exponent: 4,
			Modulus:  0b1011,
			Expected: 0b110,
		},
		{
			Base:     0b10,
			Exponent: 5,
			Modulus:  0b1011,
			Expected: 0b111,
		},
		{
			Base:     0b10,
			Exponent: 6,
			Modulus:  0b1011,
			Expected: 0b101,
		},
		{
			Base:     0b10,
			Exponent: 7,
			Expected: 0b10000000,
		},
		{
			Base:     0b10,
			Exponent: 7,
			Modulus:  0b1011,
			Expected: 0b1,
		},
		{
			Base:     0b11001, // (x^4 + x^3 + 1)^3
			Exponent: 3,
			Expected: 0b1111101011001, // x^12 + x^11 + x^10 + x^9 + x^8 + x^6 + x^4 + x^3 + 1
		},
		{
			Base:     0b11001, // (x^4 + x^3 + 1)^3
			Exponent: 3,
			Modulus:  0b100101, // x^5 + x^2 + 1
			Expected: 0b11100,  // x^4 + x^3 + x^2
		},
	}

	for _, test := range testCases {
		if actual := test.Base.Exp(test.Exponent, test.Modulus); actual != test.Expected {
			msg := fmt.Sprintf("expected (%s)^%d", test.Base, test.Exponent)
			if test.Modulus > 0 {
				msg += fmt.Sprintf(" mod %s", test.Modulus)
			}
			msg += fmt.Sprintf(" = %s (got %s)", test.Expected, actual)
			t.Error(msg)
		}
	}
}

func BenchmarkPolynomial_Exp_32(b *testing.B) {
	poly := Polynomial(0b1110000000011100000100101110110)
	modulus := Polynomial(PrimePolynomialDegree32)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		poly.Exp(0x7fffffff, modulus)
	}
}
