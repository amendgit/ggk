package ggk

/** Paint
holds the style and color information about how to draw geometries, text
and bitmaps. */
type Paint struct {
	flags       uint32
	hinting     uint8
	xfermode    *Xfermode
	looper      *DrawLooper
	imageFilter *ImageFilter

	colorFilter *ColorFilter
	style       PaintStyle
	color       Color
}

func NewPaint() *Paint {
	var paint = &Paint{
		xfermode:    NewXfermode(),
		imageFilter: NewImageFilter(),
	}
	return paint
}

func NewPaint_Clone(otr *Paint) *Paint {
	toimpl()
	return nil
}

/** Equal may give false negatives: two paints that draw equivalently
may return false.  It will never give false positives: two paints that
are not equivalent always return false. */
func (paint *Paint) Equal(otr *Paint) bool {
	toimpl()
	return false
}

/** Hash is a shallow hash, with the same limitations as operator==.
If operator== returns true for two paints, getHash() returns the same value for each. */
func (paint *Paint) Hash() uint32 {
	toimpl()
	return 0
}

func (paint *Paint) Flatten(buffer *WriteBuffer) {
	toimpl()
}

func (paint *Paint) Unflatten(buffer *ReadBuffer) {
	toimpl()
}

/** Reset restores the paint to its initial settings. */
func Reset() {
	toimpl()
}

/** PaintHinting specifies the level of hinting to be performed. These names are
taken from the Gnome/Cairo names for the same. They are translated into
Freetype concepts the same as in cairo-ft-font.c:
KPaintHintingNo     -> FT_LOAD_NO_HINTING
KPaintHintingSlight -> FT_LOAD_TARGET_LIGHT
KPaintHintingNormal -> <default, no option>
KPaintHintingFull   -> <same as KPaintHintingNormal, unelss we are rendering
					   subpixel glyphs, in which case TARGET_LCD or
					   TARGET_LCD_V is used> */
type PaintHinting int

const (
	KPaintHintingNo     = 0
	KPaintHintingSlight = 1
	KPaintHintingNormal = 2 //< this is the default.
	KPaintHintingFull   = 3
)

func (paint *Paint) Hinting() PaintHinting {
	return PaintHinting(paint.hinting)
}

func (paint *Paint) SetHinting(hinting PaintHinting) {
	paint.hinting = uint8(hinting)
}

/** Specifies the bit values that are stored in the paint's flags. */
type PaintFlags uint32

const (
	KPaintFlagAntiAlias          PaintFlags = 0x01  //< mask to enable antialiasing
	KPaintFlagDither             PaintFlags = 0x04  //< mask to enable dithering
	KPaintFlagUnderline          PaintFlags = 0x08  //< mask to enable underline text
	KPaintFlagStrikeThruText     PaintFlags = 0x10  //< mask to enable strike-thru text
	KPaintFlagFakeBoldText       PaintFlags = 0x20  //< mask to enable fake-bold text
	KPaintFlagLinearText         PaintFlags = 0x40  //< mask to enable linear-text
	KPaintFlagSubpixelText       PaintFlags = 0x80  //< mask to enable subpixel text positioning
	KPaintFlagDevKernText        PaintFlags = 0x100 //< mask to enable device kerning text
	KPaintFlagLCDRenderText      PaintFlags = 0x200 //< mask to enable subpixel glyph renderering
	KPaintFlagEmbeddedBitmapText PaintFlags = 0x400 //< mask to enable embedded bitmap strikes
	KPaintFlagAutoHinting        PaintFlags = 0x800 //< mask to force Freetype's autohinter
	KPaintFlagVerticalText       PaintFlags = 0x1000

	/** hack for GDI -- do not use if you can help it
	when adding extra flags, note that the fFlags member is specified
	with a bit-width and you'll have to expand it. */
	KPaintFlagGenA8FromLCD PaintFlags = 0x2000

	KPaintFlagAllFlags PaintFlags = 0xFFFF
)

/** Return the paint's flags. Use the Flag enum to test flag values.
@return the paint's flags (see enums ending in _Flag for bit masks) */
func (paint *Paint) Flags() PaintFlags {
	return PaintFlags(paint.flags)
}

/** Set the paint's flags. Use the Flag enum to specific flag values.
@param flags    The new flag bits for the paint (see Flags enum) */
func (paint *Paint) SetFlags(flags PaintFlags) {
	paint.flags = flags
}

/** Helper for getFlags(), returning true if kAntiAlias_Flag bit is set
@return true if the antialias bit is set in the paint's flags. */
func (paint *Paint) IsAntiAlias() bool {
	return paint.flags & KPaintFlagAntiAlias != 0
}

/** Helper for setFlags(), setting or clearing the kAntiAlias_Flag bit
@param aa   true to enable antialiasing, false to disable it */
func (paint *Paint) SetAnitAlias(aa bool) {
	toimpl()
}

/** Helper for getFlags(), returning true if kDither_Flag bit is set
@return true if the dithering bit is set in the paint's flags. */
func (paint *Paint) IsDither() bool {
	return paint.flags & KPaintFlagDither != 0
}

/** Helper for setFlags(), setting or clearing the kDither_Flag bit
@param dither   true to enable dithering, false to disable it */
func (paint *Paint) SetDither(dither bool) {
	toimpl()
}

/** Helper for getFlags(), returning true if kLinearText_Flag bit is set
@return true if the lineartext bit is set in the paint's flags */
func (paint *Paint) IsLinearText() bool {
	return paint.flags & KPaintFlagLinearText != 0
}

/** Helper for setFlags(), setting or clearing the kLinearText_Flag bit
@param linearText true to set the linearText bit in the paint's flags,
				  false to clear it. */
func (paint *Paint) SetLinearText(linearText bool) {
	toimpl()
}

/** Helper for getFlags(), returning true if kSubpixelText_Flag bit is set
@return true if the lineartext bit is set in the paint's flags */
func (paint *Paint) IsSubpixelText() bool {
	return paint.flags & KPaintFlagSubpixelText != 0
}

/**
 *  Helper for setFlags(), setting or clearing the kSubpixelText_Flag.
 *  @param subpixelText true to set the subpixelText bit in the paint's
 *                      flags, false to clear it.
 */
func (paint *Paint) SetSubpixelText(subpixelText bool) {
	toimpl()
}

func (paint *Paint) IsLCDRenderText() bool {
	return uint32(paint.flags) & KPaintFlagLCDRenderText != 0
}

/**
 *  Helper for setFlags(), setting or clearing the kLCDRenderText_Flag.
 *  Note: antialiasing must also be on for lcd rendering
 *  @param lcdText true to set the LCDRenderText bit in the paint's flags,
 *                 false to clear it.
 */
func (paint *Paint) SetLCDRenderText(lcdRencderText bool) {
	toimpl()
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

func (paint *Paint) Color() Color {
	toimpl()
	return KColorBlack
}

func (paint *Paint) SetXfermodeMode(mode XfermodeMode) *Xfermode {
	paint.xfermode = NewXfermodeWithMode(mode)
	return paint.xfermode
}

type PaintStyle int

const (
	KPaintStyleFill          = PaintStyle(iota) // < fill the geometry
	KPaintStyleStroke                           // < stroke the geometry
	KPaintStyleStrokeAndFill                    // < fill and stroke the geometry
)

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

func (paint *Paint) Shader() *Shader {
	toimpl()
	return nil
}

/** Get the paint's colorfilter. If there is a colorfilter, its reference
count is not changed.
@return the paint's colorfilter (or NULL)
*/
func (paint *Paint) ColorFilter() *ColorFilter {
	return paint.colorFilter
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

func (paint *Paint) MaskFilter() *MaskFilter {
	toimpl()
	return nil
}

func (paint *Paint) SetShader(shader *Shader) {
	toimpl()
}

func (paint *Paint) SetAlpha(alpha uint8) {
	toimpl()
}

func (paint *Paint) Alpha() uint8 {
	toimpl()
	return 0
}
