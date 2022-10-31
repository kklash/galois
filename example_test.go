package galois_test

import (
	"fmt"

	"github.com/kklash/galois"
)

// This example demonstrates the construction and use of a Finite
// Field of order 256, AKA GF(256).
//
// This is a very commonly used finite field, as the elements of GF(256)
// have a bijective correspondence (one-to-one mapping) to all possible values
// of a byte: 0 to 255.
//
// GF(256) is useful because it allows you to define byte-wise division and
// multiplication with unique outputs which stay within the size of a byte.
func ExampleField() {
	field := galois.NewField[uint8](galois.PrimePolynomialDegree8)

	x := uint8(0xBC)
	y := uint8(0xDE)
	z := uint8(0xFF)

	xMulY := field.Mul(x, y)

	fmt.Printf(
		"0x%X * 0x%X = 0x%X\n",
		x, y, xMulY,
	)
	fmt.Printf(
		"0x%X / 0x%X = 0x%X\n",
		xMulY, x, field.Div(xMulY, x),
	)

	zMulY := field.Mul(z, y)

	fmt.Printf(
		"0x%X * 0x%X = 0x%X\n",
		z, y, zMulY,
	)
	fmt.Printf(
		"0x%X / 0x%X = 0x%X\n",
		zMulY, y, field.Div(zMulY, y),
	)

	// output:
	// 0xBC * 0xDE = 0x6D
	// 0x6D / 0xBC = 0xDE
	// 0xFF * 0xDE = 0x8B
	// 0x8B / 0xDE = 0xFF
}

// This example demonstrates basic multiplication of two elements using a Field.
//
// Note how the choice of prime polynomial matters very much. Even if two prime
// polynomials have the same degree, and thus generate a field of the same order,
// the results of arithmetic operations within those fields will be different.
func ExampleField_Mul() {
	field1 := galois.NewField[uint16](galois.PrimePolynomialDegree16)
	a := uint16(0x1234)
	b := uint16(0x5678)
	fmt.Printf("field1: 0x%X * 0x%X = 0x%X\n", a, b, field1.Mul(a, b))

	// x^16 + x^13 + x^12 + x^10 + x^9 + x^7 + x^6 + x^1 + 1
	field2 := galois.NewField[uint16](0b10011011011000011)
	fmt.Printf("field2: 0x%X * 0x%X = 0x%X\n", a, b, field2.Mul(a, b))

	// output:
	// field1: 0x1234 * 0x5678 = 0x6324
	// field2: 0x1234 * 0x5678 = 0xA051
}
