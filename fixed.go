package ggk

// F24d8 is a 24.8 integer fixed point.
type F24d8 int32

func F24d8Floor(x F24d8) int {
	return int(x >> 8)
}

func F24d8Ceil(x F24d8) int {
	return int(x+0xFF) >> 8
}

func F24d8FromScalar(x Scalar) F24d8 {
	return F24d8(x * 256)
}
