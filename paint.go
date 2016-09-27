package ggk

// Paint holds the style and color information about how to draw geometries, text
// and bitmaps.
type Paint struct {
	flags   uint16
	hinting uint8
	looper  *DrawLooper
}

func NewPaint() *Paint {
	toimpl()
	return nil
}

// PaintHinting specifies the level of hinting to be performed. These names are
// taken from the Gnome/Cairo names for the same. They are translated into
// Freetype concepts the same as in cairo-ft-font.c:
// KPaintHintingNo     -> FT_LOAD_NO_HINTING
// KPaintHintingSlight -> FT_LOAD_TARGET_LIGHT
// KPaintHintingNormal -> <default, no option>
// KPaintHintingFull   -> <same as KPaintHintingNormal, unelss we are rendering
//                         subpixel glyphs, in which case TARGET_LCD or
//                         TARGET_LCD_V is used>
type PaintHinting int

const (
	KPaintHintingNo     = 0
	KPaintHintingSlight = 1
	KPaintHintingNormal = 2 // this is the default.
	KPaintHintingFull   = 3
)

func (p *Paint) Hinting() PaintHinting {
	return PaintHinting(p.hinting)
}

func (p *Paint) SetHinting(hinting PaintHinting) {
	p.hinting = uint8(hinting)
}

func (p *Paint) Looper() *DrawLooper {
	return p.looper
}

func (p *Paint) SetLooper(looper *DrawLooper) {
	p.looper = looper
}

type PaintFlags int

const (
	KPaintFlagAntiAlias          = 0x01
	KPaintFlagDither             = 0x04
	KPaintFlagUnderline          = 0x08
	KPaintFlagStrikeThruText     = 0x10
	KPaintFlagFakeBoldText       = 0x20
	KPaintFlagLinearText         = 0x40
	KPaintFlagSubpixelText       = 0x80
	KPaintFlagDevKernText        = 0x100
	KPaintFlagLCDRenderText      = 0x200
	KPaintFlagEmbeddedBitmapText = 0x400
	KPaintFlagAutoHinting        = 0x800
	KPaintFlagVerticalText       = 0x1000

	// hack for GDI -- do not use if you can help it when adding extra flags,
	// note that the flags member is specified with a bit-width and you'll have
	// expand it.
	KPaintFlagGenA8FromLCD = 0x2000

	KPaintFlagAllFlags = 0xFFFF
)

type PaintStyle int

const (
	kPaintStyleFill = iota     //!< fill the geometry
	kPaintStyleStroke          //!< stroke the geometry
	kPaintStyleStrokeAndFill   //!< fill and stroke the geometry
)

func (p *Paint) Flags() PaintFlags {
	return PaintFlags(p.flags)
}

func (p *Paint) SetFlags(flags PaintFlags) {
	p.flags = uint16(flags)
}

func (paint *Paint) CanComputeFastBounds() bool {
	// if paint.Looper() != nil {
	// 	return paint.Looper().CanComputeFastBounds()
	// }
	// if paint.ImageFilter() != nil && paint.ImageFilter().CanComputeFastBounds() {
	// 	return false
	// }
	// return !paint.Rasterizer()
	toimpl()
	return false
}

func (paint *Paint)ComputeFastStrokeBounds(orig Rect, storage *Rect) Rect {
	return paint.doComputeFastStrokeBounds(orig, storage, kPaintStyleStroke)
}

// Take the style explicitly, so the caller can force us to be stroked
// without having to make a copy of the paint just to change that field.
func (paint *Paint)doComputeFastStrokeBounds(orig Rect, storage *Rect, style PaintStyle) Rect {
	toimpl()
	return RectZero
}

func (paint *Paint) ImageFilter() *ImageFilter {
	toimpl()
	return nil
}

func (paint *Paint) Rasterizer() bool {
	toimpl()
	return false
}

func (paint *Paint) SetColorFilter(colorFilter *ColorFilter) {
	toimpl()
}

func (paint *Paint) SetImageFilter(imageFilter *ImageFilter) {
	toimpl()
}

func (paint *Paint) SetXfermode(xfermode *tXfermode) {
	toimpl()
}

func (paint *Paint) Xfermode() *tXfermode {
	toimpl()
	return nil
}

func (paint *Paint) NothingToDraw() bool {
	toimpl()
	return false
}