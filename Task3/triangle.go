// Package triangle contains mathods to detect what kind given triangle is.
//
package triangle

// Kind triangle type representation
type Kind int

const (
	// NaT is not a triangle
	NaT = iota
	// Equ equilateral
	Equ
	// Iso isosceles
	Iso
	// Sca scalene
	Sca
)
