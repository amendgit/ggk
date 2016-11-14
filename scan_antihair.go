package ggk

func antiFillF24d8(L, T, R, B F24d8, blitter Blitter, fillInner bool) {
	// Check for empty now that we're in our reduced precission space.
	toimpl()
}

func callHlineBlitter(blitter Blitter, x, y, count int, alpha uint8) {
	toimpl()
}

// calls blitRect() if the rectangle is non-empty
func fillCheckRect(L, T, R, B int, blitter Blitter) {
	if L < R && T < B {
		blitter.BlitRect(L, T, R-L, B-T)
	}
}

func innerStrokeF24d8(L, T, R, B F24d8, blitter Blitter) {
	toimpl()
}

func scanLine(L F24d8, top int, R F24d8, alpha uint8, blitter Blitter) {
	if (L >> 8) == ((R - 1) >> 8) { // < 1x1 pixel
		blitter.BlitV(int(L>>8), top, 1, Alpha(InvAlphaMul(alpha, uint8(R-L))))
		return
	}

	var left = int(L >> 8)

	if L&0xFF != 0 {
		blitter.BlitV(int(left), int(top), int(1), Alpha(InvAlphaMul(alpha, 255-(uint8(L)&0xFF)+1)))
		left++
	}

	var rite = int(R) >> 8
	var width = rite - left
	if width > 0 {
		callHlineBlitter(blitter, left, top, width, alpha)
	}

	if R&0xFF != 0 {
		blitter.BlitV(rite, top, 1, Alpha(InvAlphaMul(alpha, uint8(R)&0xFF)))
	}
}

func AlphaMulRound(a, b uint8) uint8 {
	return MulDiv255Round(a, b)
}

// 1 - (1 - a)*(1 - b)
func InvAlphaMul(a, b uint8) uint8 {
	// need precise rounding (not just SkAlphaMul) so that values like
	// a=228, b=252 don't overflow the result
	return uint8(a + b - AlphaMulRound(a, b))
}

func ScanAntiFrameRect(r Rect, strokeSize Point, clip *Region, blitter Blitter) {
	var (
		rx Scalar = ScalarHalf(strokeSize.X)
		ry Scalar = ScalarHalf(strokeSize.Y)
	)

	// Outset by the radius.
	var (
		outerL F24d8 = F24d8FromScalar(r.L() - rx)
		outerT F24d8 = F24d8FromScalar(r.T() - ry)
		outerR F24d8 = F24d8FromScalar(r.R() + rx)
		outerB F24d8 = F24d8FromScalar(r.B() + ry)
	)

	var outer Rect
	outer.SetLTRB(Scalar(F24d8Floor(outerL)), Scalar(F24d8Floor(outerT)), Scalar(F24d8Floor(outerR)), Scalar(F24d8Floor(outerB)))

	// var clipper *BlitterClipper
	rx = strokeSize.X - rx
	ry = strokeSize.Y - ry

	// Inset by the radius.
	var (
		innerL F24d8 = F24d8FromScalar(r.L() + rx)
		innerT F24d8 = F24d8FromScalar(r.T() + rx)
		innerR F24d8 = F24d8FromScalar(r.R() - rx)
		innerB F24d8 = F24d8FromScalar(r.B() - rx)
	)

	// Stroke the outter hull.
	antiFillF24d8(outerL, outerT, outerR, outerB, blitter, false)

	// Set outer to the outer rect of the middle section.
	outer.SetLTRB(Scalar(F24d8Ceil(outerL)), Scalar(F24d8Ceil(outerT)), Scalar(F24d8Floor(outerR)), Scalar(F24d8Floor(outerB)))

	if innerL >= innerR || innerT >= innerB {
		fillCheckRect(int(outer.L()), int(outer.T()), int(outer.R()), int(outer.B()), blitter)
	} else {
		var inner Rect
		// Set inner to the inner rect of the middle section.
		inner.SetLTRB(Scalar(F24d8Floor(innerL)), Scalar(F24d8Floor(innerT)), Scalar(F24d8Ceil(innerR)), Scalar(F24d8Ceil(innerB)))
		// Draw the frame in 4 pieces.
		fillCheckRect(int(outer.L()), int(outer.T()), int(outer.R()), int(inner.T()), blitter)
		fillCheckRect(int(outer.L()), int(inner.T()), int(inner.L()), int(inner.B()), blitter)
		fillCheckRect(int(inner.R()), int(inner.T()), int(outer.R()), int(inner.B()), blitter)
		fillCheckRect(int(outer.L()), int(inner.B()), int(outer.R()), int(outer.B()), blitter)

		// Now stroke the inner rect, which is similar to antifilldot8() except that
		// it treats the fractional coordinates with the inverse bias (since its
		// inner).
		innerStrokeF24d8(innerL, innerT, innerR, innerB, blitter)
	}
}
