package ggk

// Paint holds the style and color information about how to draw geometries, text
// and bitmaps.
type Paint struct {
	flags       uint16
	hinting     uint8
	xfermode    *Xfermode
	looper      *DrawLooper
	imageFilter *ImageFilter

	style PaintStyle
	color Color
}

func NewPaint() *Paint {
	var paint = &Paint{
		xfermode:    NewXfermode(),
		imageFilter: NewImageFilter(),
	}
	return paint
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

func (paint *Paint) Hinting() PaintHinting {
	return PaintHinting(paint.hinting)
}

func (paint *Paint) SetHinting(hinting PaintHinting) {
	paint.hinting = uint8(hinting)
}

func (paint *Paint) Looper() *DrawLooper {
	return paint.looper
}

func (paint *Paint) SetLooper(looper *DrawLooper) {
	paint.looper = looper
}

func (paint *Paint) SetColor(color Color) {
	paint.color = color
}

func (paint *Paint) SetXfermodeMode(mode XfermodeMode) *Xfermode {
	paint.xfermode = NewXfermodeWithMode(mode)
	return paint.xfermode
}

type PaintFlags int

const (
	KPaintFlagAntiAlias          PaintFlags = 0x01
	KPaintFlagDither             PaintFlags = 0x04
	KPaintFlagUnderline          PaintFlags = 0x08
	KPaintFlagStrikeThruText     PaintFlags = 0x10
	KPaintFlagFakeBoldText       PaintFlags = 0x20
	KPaintFlagLinearText         PaintFlags = 0x40
	KPaintFlagSubpixelText       PaintFlags = 0x80
	KPaintFlagDevKernText        PaintFlags = 0x100
	KPaintFlagLCDRenderText      PaintFlags = 0x200
	KPaintFlagEmbeddedBitmapText PaintFlags = 0x400
	KPaintFlagAutoHinting        PaintFlags = 0x800
	KPaintFlagVerticalText       PaintFlags = 0x1000

	// hack for GDI -- do not use if you can help it when adding extra flags,
	// note that the flags member is specified with a bit-width and you'll have
	// expand it.
	KPaintFlagGenA8FromLCD PaintFlags = 0x2000

	KPaintFlagAllFlags PaintFlags = 0xFFFF
)

type PaintStyle int

const (
	KPaintStyleFill          = PaintStyle(iota) // < fill the geometry
	KPaintStyleStroke                           // < stroke the geometry
	KPaintStyleStrokeAndFill                    // < fill and stroke the geometry
)

func (paint *Paint) Flags() PaintFlags {
	return PaintFlags(paint.flags)
}

func (paint *Paint) SetFlags(flags PaintFlags) {
	paint.flags = uint16(flags)
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

func (paint *Paint) ComputeFastStrokeBounds(orig Rect, storage *Rect) Rect {
	return paint.doComputeFastStrokeBounds(orig, storage, KPaintStyleStroke)
}

// Take the style explicitly, so the caller can force us to be stroked
// without having to make a copy of the paint just to change that field.
func (paint *Paint) doComputeFastStrokeBounds(orig Rect, storage *Rect, style PaintStyle) Rect {
	toimpl()
	return RectZero
}

func (paint *Paint) ImageFilter() *ImageFilter {
	return paint.imageFilter
}

func (paint *Paint) Rasterizer() bool {
	toimpl()
	return false
}

/** Get the paint's colorfilter. If there is a colorfilter, its reference
count is not changed.
@return the paint's colorfilter (or NULL)
*/
func (paint *Paint) ColorFilter() *ColorFilter {
	return paint.ColorFilter()
}

func (paint *Paint) SetColorFilter(colorFilter *ColorFilter) {
	toimpl()
}

func (paint *Paint) SetImageFilter(imageFilter *ImageFilter) {
	toimpl()
}

func (paint *Paint) SetXfermode(xfermode *Xfermode) {
	paint.xfermode = xfermode
}

/** Xfermode
@return the paint's xfermode or nil. */
func (paint *Paint) Xfermode() *Xfermode {
	return paint.xfermode
}

func (paint *Paint) NothingToDraw() bool {
	toimpl()
	return false
}

func (paint *Paint) Clone() *Paint {
	toimpl()
	return nil
}

func (paint *Paint) SetStyle(style PaintStyle) {
	paint.style = style
}

func (paint *Paint) StrokeWidth() int {
	toimpl()
	return 0
}
