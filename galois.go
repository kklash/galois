// Package galois implements finite fields with non-prime order 2^m,
// up to a maximum of 2^32.
//
// To achieve this, we create fields whose elements are polynomials with binary
// coefficients (one and zero) with operations between coefficients taken modulo two.
// All operations between polynomials are taken modulo an irreducible prime.
// Operations within this domain form a finite field (Galois Field) with an order
// that is a power of two.
//
// We use prime polynomials sourced from https://www.partow.net/programming/polynomials/index.html
package galois
