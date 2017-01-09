package ggk

var gLCDOrientation LCDOrientation = KLCDOrientationHorizontal

/** LCD color elements can vary in order. For subpixel text we need to know
	the order which the LCDs uses so that the color fringes are in the
	correct place.

	Note, if you change this after startup, you'll need to flush the glyph
	cache because it'll have the wrong type of masks cached.

	kNONE_LCDOrder means that the subpixel elements are not spatially
	separated in any usable fashion.

	@deprecated use SkPixelGeometry instead.
 */
type LCDOrder int

const (
	KLCDOrderRGB = iota //< this is the default.
	KLCDOrderBGR
	KLCDOrderNone
)

/** LCDs either have their color elements arranged horizontally or
	vertically. When rendering subpixel glyphs we need to know which way
	round they are.

	Note, if you change this after startup, you'll need to flush the glyph
	cache because it'll have the wrong type of masks cached.

	@deprecated use SkPixelGeometry instead.
*/
type LCDOrientation int

const (
	KLCDOrientationHorizontal = iota // this is the default
	KLCDOrientationVertical
)

type FontLCDConfig struct {
}

/** @deprecated set on Device creation. */
func FontLCDConfigSetSubpixelOrientation(orientation LCDOrientation) {
	toimpl()
}

/** @deprecated get from Device. */
func FontLCDConfigSubpixelOrientation() LCDOrientation {
	return gLCDOrientation
}