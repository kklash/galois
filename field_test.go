package galois

import "testing"

func TestField_OverflowSafeConstructor(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Errorf(
				"expected to panic when creating Field of order 2^17 with only uint16 representations",
			)
		}
	}()

	NewField[uint16](PrimePolynomialDegree17)
}

func TestField_Generate(t *testing.T) {
	type TestCase struct {
		Modulus  Polynomial
		Exponent uint64
		Expected Polynomial
	}

	testCases := []TestCase{
		{
			Modulus:  PrimePolynomialDegree3,
			Exponent: 0,
			Expected: 0b001,
		},
		{
			Modulus:  PrimePolynomialDegree3,
			Exponent: 1,
			Expected: 0b010,
		},
		{
			Modulus:  PrimePolynomialDegree3,
			Exponent: 2,
			Expected: 0b100,
		},
		{
			Modulus:  PrimePolynomialDegree3,
			Exponent: 3,
			Expected: 0b011,
		},
		{
			Modulus:  PrimePolynomialDegree3,
			Exponent: 4,
			Expected: 0b110,
		},
		{
			Modulus:  PrimePolynomialDegree3,
			Exponent: 5,
			Expected: 0b111,
		},
		{
			Modulus:  PrimePolynomialDegree3,
			Exponent: 6,
			Expected: 0b101,
		},
		{
			Modulus:  PrimePolynomialDegree3,
			Exponent: 7,
			Expected: 0b001,
		},
		{
			Modulus:  PrimePolynomialDegree3,
			Exponent: 8,
			Expected: 0b010,
		},
	}

	for _, test := range testCases {
		field := NewField[Polynomial](test.Modulus)
		actual := field.Generate(test.Exponent)

		if actual != test.Expected {
			t.Errorf(
				"expected x^%d mod %s = %s - got %q",
				test.Exponent, test.Modulus, test.Expected, actual,
			)
		}
	}
}

func TestField_Exp(t *testing.T) {
	type TestCase struct {
		Base     Polynomial
		Exponent uint64
		Modulus  Polynomial
		Expected Polynomial
	}

	testCases := []TestCase{
		{
			Base:     0b100101,
			Exponent: 0,
			Modulus:  PrimePolynomialDegree6,
			Expected: 1,
		},
		{
			Base:     0b100101,
			Exponent: 1,
			Modulus:  PrimePolynomialDegree6,
			Expected: 0b100101,
		},
		{
			Base:     0b100101,
			Exponent: 4,
			Modulus:  PrimePolynomialDegree3,
			Expected: 0b110,
		},
		{
			Base:     0b11,
			Exponent: 7,
			Modulus:  PrimePolynomialDegree4,
			Expected: 0b1101,
		},
		{
			Base:     0b1011,
			Exponent: fieldOrder(PrimePolynomialDegree4) - 1,
			Modulus:  PrimePolynomialDegree4,
			Expected: 0b1,
		},
		{
			Base:     0b1011,
			Exponent: fieldOrder(PrimePolynomialDegree4),
			Modulus:  PrimePolynomialDegree4,
			Expected: 0b1011,
		},
		{
			Base:     0b10011,
			Exponent: fieldOrder(PrimePolynomialDegree5),
			Modulus:  PrimePolynomialDegree5,
			Expected: 0b10011,
		},
		{
			Base:     0b10001,
			Exponent: 2,
			Modulus:  PrimePolynomialDegree5,
			Expected: 0b1100,
		},
	}

	for _, test := range testCases {
		field := NewField[Polynomial](test.Modulus)
		actual := field.Exp(test.Base, test.Exponent)

		if actual != test.Expected {
			t.Errorf(
				"field element exponentiation failed; expected (%s)^%d mod %s = %s (got %s)",
				test.Base, test.Exponent, test.Modulus, test.Expected, actual,
			)
		}
	}
}

func TestField_Mul(t *testing.T) {
	type TestCase struct {
		ExponentA uint64
		ExponentB uint64
		Modulus   Polynomial
	}

	testCases := []TestCase{
		{
			ExponentA: 0,
			ExponentB: 0,
			Modulus:   PrimePolynomialDegree8,
		},
		{
			ExponentA: 0,
			ExponentB: 1,
			Modulus:   PrimePolynomialDegree8,
		},
		{
			ExponentA: 1,
			ExponentB: 1,
			Modulus:   PrimePolynomialDegree8,
		},
		{
			ExponentA: 1,
			ExponentB: 2,
			Modulus:   PrimePolynomialDegree8,
		},
		{
			ExponentA: 2,
			ExponentB: 2,
			Modulus:   PrimePolynomialDegree8,
		},
		{
			ExponentA: 3,
			ExponentB: 7,
			Modulus:   PrimePolynomialDegree8,
		},
		{
			ExponentA: 11,
			ExponentB: 254,
			Modulus:   PrimePolynomialDegree8,
		},
		{
			ExponentA: 374,
			ExponentB: 481,
			Modulus:   PrimePolynomialDegree8,
		},
	}

	for _, test := range testCases {
		field := NewField[Polynomial](test.Modulus)
		elementA := field.Generate(test.ExponentA)
		elementB := field.Generate(test.ExponentB)

		expectedExponent := (test.ExponentA + test.ExponentB) % (field.Order() - 1)
		expectedElement := field.Generate(expectedExponent)

		if product := field.Mul(elementA, elementB); product != expectedElement {
			t.Errorf(
				"field element multiplication failed; expected (%s) * (%s) = %s (got %s)",
				elementA, elementB, expectedElement, product,
			)
		}
	}
}

func TestField_UniqueInverse(t *testing.T) {
	primes := []Polynomial{
		PrimePolynomialDegree2,
		PrimePolynomialDegree3,
		PrimePolynomialDegree4,
		PrimePolynomialDegree5,
		PrimePolynomialDegree6,
		PrimePolynomialDegree7,
		PrimePolynomialDegree8,
		PrimePolynomialDegree9,
		PrimePolynomialDegree10,
		PrimePolynomialDegree11,
		PrimePolynomialDegree12,
		PrimePolynomialDegree13,
		PrimePolynomialDegree14,
		PrimePolynomialDegree15,
		PrimePolynomialDegree16,
	}

	for _, prime := range primes {
		field := NewField[Polynomial](prime)
		seen := make(map[Polynomial]uint64)

		for index := uint64(0); index < field.Order()-1; index++ {
			poly := field.Generate(index)
			if otherIndex, ok := seen[poly]; ok {
				t.Errorf(
					"unexpected duplicate polynomial %q in finite field %q, generated by x^%d and x^%d",
					poly, prime, index, otherIndex,
				)
			}
			seen[poly] = index

			inverted := field.MultInverse(poly)

			if inverseProduct := field.Mul(poly, inverted); inverseProduct != 1 {
				t.Errorf(
					"expected element to have an inverse which yields 1 when multiplied;\n"+
						"found (%s) * (%s) = %s", poly, inverted, inverseProduct,
				)
				return
			}
		}
	}
}

func TestDistributivity(t *testing.T) {
	t.Run("GF(256): 100(10 + 5) = 100*10 + 100*5", func(t *testing.T) {
		field := NewField[uint8](PrimePolynomialDegree8)

		// 100*(10 + 5)
		resultA := field.Mul(100, field.Add(10, 5))

		// 100*10 + 100*5
		resultB := field.Add(field.Mul(100, 10), field.Mul(100, 5))

		if resultA != resultB {
			t.Errorf(
				"failed distributivity test\n100*(10 + 5) = %d\n100*10 + 100*5 = %d",
				resultA, resultB,
			)
		}
	})

	t.Run("GF(32): 5(9 + 10 + 11) = 5*9 + 5*10", func(t *testing.T) {
		field := NewField[uint8](PrimePolynomialDegree4)

		// 5*(9 + 10 + 11)
		resultA := field.Mul(
			5,
			field.Add(9, 10, 11),
		)

		// 5*9 + 5*10 + 5*11
		resultB := field.Add(
			field.Mul(5, 9),
			field.Mul(5, 10),
			field.Mul(5, 11),
		)

		if resultA != resultB {
			t.Errorf(
				"failed distributivity test\n5*(9 + 10 + 11) = %d\n5*9 + 5*10 + 5*11 = %d",
				resultA, resultB,
			)
		}
	})
}

func BenchmarkField_Generate_8(b *testing.B) {
	field := NewField[uint8](PrimePolynomialDegree8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field.Generate(0b1111111)
	}
}
func BenchmarkField_Generate_12(b *testing.B) {
	field := NewField[uint16](PrimePolynomialDegree12)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		field.Generate(0b11111111111)
	}
}
func BenchmarkField_Generate_16(b *testing.B) {
	field := NewField[uint16](PrimePolynomialDegree16)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field.Generate(0b111111111111111)
	}
}
func BenchmarkField_Generate_20(b *testing.B) {
	field := NewField[uint32](PrimePolynomialDegree20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field.Generate(0b1111111111111111111)
	}
}
func BenchmarkField_Generate_32(b *testing.B) {
	field := NewField[uint32](PrimePolynomialDegree32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field.Generate(0b1111111111111111111111111111111)
	}
}

func BenchmarkField_Exp_32(b *testing.B) {
	field := NewField[uint32](PrimePolynomialDegree32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field.Exp(0b10101010101010101010101010101010, field.Order()-2)
	}
}

func BenchmarkField_MultInverse_8(b *testing.B) {
	field := NewField[uint8](PrimePolynomialDegree8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field.MultInverse(0b10101010)
	}
}
func BenchmarkField_MultInverse_16(b *testing.B) {
	field := NewField[uint16](PrimePolynomialDegree16)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field.MultInverse(0b1010101010101010)
	}
}
func BenchmarkField_MultInverse_24(b *testing.B) {
	field := NewField[uint32](PrimePolynomialDegree24)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field.MultInverse(0b101010101010101010101010)
	}
}
func BenchmarkField_MultInverse_32(b *testing.B) {
	field := NewField[uint32](PrimePolynomialDegree32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		field.MultInverse(0b10101010101010101010101010101010)
	}
}
