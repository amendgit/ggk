package ggk

// Return a*b/255, rounding any fractional bits.
// Only valid if a and b are unsigned and <= 0x7fff
func MulDiv255Round(a uint8, b uint8) uint8 {

	var prod uint8 = uint8(uint16(a) * uint16(b) + 128)
	return (prod + (prod >> 8)) >> 8
}
