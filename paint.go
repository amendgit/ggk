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
	return paint.flags&KPaintFlagAntiAlias != 0
}

/** Helper for setFlags(), setting or clearing the kAntiAlias_Flag bit
@param aa   true to enable antialiasing, false to disable it */
func (paint *Paint) SetAnitAlias(aa bool) {
	toimpl()
}

/** Helper for getFlags(), returning true if kDither_Flag bit is set
@return true if the dithering bit is set in the paint's flags. */
func (paint *Paint) IsDither() bool {
	return paint.flags&KPaintFlagDither != 0
}

/** Helper for setFlags(), setting or clearing the kDither_Flag bit
@param dither   true to enable dithering, false to disable it */
func (paint *Paint) SetDither(dither bool) {
	toimpl()
}

/** Helper for getFlags(), returning true if kLinearText_Flag bit is set
@return true if the lineartext bit is set in the paint's flags */
func (paint *Paint) IsLinearText() bool {
	return paint.flags&KPaintFlagLinearText != 0
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
	return paint.flags&KPaintFlagSubpixelText != 0
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
	return uint32(paint.flags)&KPaintFlagLCDRenderText != 0
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

func (paint *Paint) IsEmbeddedBitmapText() bool {
	toimpl()
	return false
}

/** Helper for setFlags(), setting or clearing the kEmbeddedBitmapText_Flag bit
@param useEmbeddedBitmapText true to set the kEmbeddedBitmapText bit in the paint's flags,
							 false to clear it.
*/
func (paint *Paint) SetEmbeddedBitmapText(useEmbeddedBitmapText bool) {
	toimpl()
}

func (paint *Paint) IsAutohinted() bool {
	toimpl()
	return false
}

/** Helper for setFlags(), setting or clearing the kAutoHinting_Flag bit
@param useAutohinter true to set the kEmbeddedBitmapText bit in the
						  paint's flags,
					 false to clear it.
*/
func (paint *Paint) SetAutohinted(useAutohinted bool) {
	toimpl()
}

func (paint *Paint) IsVerticalText() bool {
	toimpl()
	return false
}

/**
 *  Helper for setting or clearing the kVerticalText_Flag bit in
 *  setFlags(...).
 *
 *  If this bit is set, then advances are treated as Y values rather than
 *  X values, and drawText will places its glyphs vertically rather than
 *  horizontally.
 */
func (paint *Paint) SetVerticalText(useVerticalText bool) {
	toimpl()
}

/** Helper for getFlags(), returning true if kUnderlineText_Flag bit is set
@return true if the underlineText bit is set in the paint's flags. */
func (paint *Paint) IsUnderlineText() bool {
	toimpl()
	return false
}

/** Helper for setFlags(), setting or clearing the kUnderlineText_Flag bit
@param underlineText true to set the underlineText bit in the paint's
					 flags, false to clear it. */
func (paint *Paint) SetUnderlineText(underlineText bool) {
	toimpl()
}

/** Helper for getFlags(), returns true if kStrikeThruText_Flag bit is set
@return true if the strikeThruText bit is set in the paint's flags. */
func (paint *Paint) IsStrikeThruText() bool {
	toimpl()
	return false
}

/** Helper for setFlags(), setting or clearing the kStrikeThruText_Flag bit
@param strikeThruText   true to set the strikeThruText bit in the
						paint's flags, false to clear it. */
func (paint *Paint) SetStrikeThruText(strikeThruText bool) {
	toimpl()
}

/** Helper for getFlags(), returns true if kFakeBoldText_Flag bit is set
@return true if the kFakeBoldText_Flag bit is set in the paint's flags.
*/
func (paint *Paint) IsFakeBoldText() bool {
	toimpl()
	return false
}

/** Helper for setFlags(), setting or clearing the kFakeBoldText_Flag bit
@param fakeBoldText true to set the kFakeBoldText_Flag bit in the paint's
					flags, false to clear it.
*/
func (paint *Paint) SetFakeBoldText(fakeBoldText bool) {
	toimpl()
}

/** Helper for getFlags(), returns true if kDevKernText_Flag bit is set
@return true if the kernText bit is set in the paint's flags.
*/
func (paint *Paint) IsDevKernText() bool {
	toimpl()
	return false
}

/** Helper for setFlags(), setting or clearing the kKernText_Flag bit
@param kernText true to set the kKernText_Flag bit in the paint's
					flags, false to clear it.
*/
func (paint *Paint) SetDevKernText(devKernText bool) {
	toimpl()
}

/**
 *  Return the filter level. This affects the quality (and performance) of
 *  drawing scaled images.
 */
func (paint *Paint) FilterQuality() FilterQuality {
	toimpl()
	return KFilterQualityNone
}

/**
 *  Set the filter quality. This affects the quality (and performance) of
 *  drawing scaled images.
 */
func (paint *Paint) SetFilterQuality(quality FilterQuality) {
	toimpl()
}

/** Styles apply to rect, oval, path, and text.
  Bitmaps are always drawn in "fill", and lines are always drawn in
  "stroke".

  Note: strokeandfill implicitly draws the result with
  SkPath::kWinding_FillType, so if the original path is even-odd, the
  results may not appear the same as if it was drawn twice, filled and
  then stroked.
*/
type PaintStyle int

const (
	KPaintStyleFill          = PaintStyle(iota) // < fill the geometry
	KPaintStyleStroke                           // < stroke the geometry
	KPaintStyleStrokeAndFill                    // < fill and stroke the geometry
	KPaintStyleCount         = KPaintStyleStrokeAndFill + 1
)

/** Return the paint's style, used for controlling how primitives'
geometries are interpreted (except for drawBitmap, which always assumes
kFill_Style).
@return the paint's Style
*/
func (paint *Paint) Style() PaintStyle {
	toimpl()
	return KPaintStyleFill
}

/** Set the paint's style, used for controlling how primitives'
geometries are interpreted (except for drawBitmap, which always assumes
Fill).
@param style    The new style to set in the paint */
func (paint *Paint) SetStyle(style PaintStyle) {
	paint.style = style
}

/** Return the paint's color. Note that the color is a 32bit value
containing alpha as well as r,g,b. This 32bit value is not
premultiplied, meaning that its alpha can be any value, regardless of
the values of r,g,b.
@return the paint's color (and alpha). */
func (paint *Paint) Color() Color {
	toimpl()
	return KColorBlack
}

/** Set the paint's color. Note that the color is a 32bit value containing
alpha as well as r,g,b. This 32bit value is not premultiplied, meaning
that its alpha can be any value, regardless of the values of r,g,b.
@param color    The new color (including alpha) to set in the paint. */
func (paint *Paint) SetColor(color Color) {
	paint.color = color
}

/** Helper to getColor() that just returns the color's alpha value.
@return the alpha component of the paint's color. */
func (paint *Paint) Alpha() uint8 {
	toimpl()
	return 0
}

/** Helper to setColor(), that only assigns the color's alpha value,
leaving its r,g,b values unchanged.
@param a    set the alpha component (0..255) of the paint's color. */
func (paint *Paint) SetAlpha(alpha uint8) {
	toimpl()
}

/** Helper to setColor(), that takes a,r,g,b and constructs the color value
using SkColorSetARGB()
@param a    The new alpha component (0..255) of the paint's color.
@param r    The new red component (0..255) of the paint's color.
@param g    The new green component (0..255) of the paint's color.
@param b    The new blue component (0..255) of the paint's color.
*/
func (paint *Paint) SetARGB(a, r, g, b uint8) {
	toimpl()
}

/** Return the width for stroking.
<p />
A value of 0 strokes in hairline mode.
Hairlines always draw 1-pixel wide, regardless of the matrix.
@return the paint's stroke width, used whenever the paint's style is
		Stroke or StrokeAndFill. */
func (paint *Paint) StrokeWidth() int {
	toimpl()
	return 0
}

/** Set the width for stroking.
Pass 0 to stroke in hairline mode.
Hairlines always draw 1-pixel wide, regardless of the matrix.
@param width set the paint's stroke width, used whenever the paint's
			 style is Stroke or StrokeAndFill. */
func (paint *Paint) SetStrokeWidth(width Scalar) {
	toimpl()
}

/** Return the paint's stroke miter value. This is used to control the
behavior of miter joins when the joins angle is sharp.
@return the paint's miter limit, used whenever the paint's style is
		Stroke or StrokeAndFill. */
func (paint *Paint) StrokeMiter() Scalar {
	toimpl()
	return Scalar(0)
}

/** Set the paint's stroke miter value. This is used to control the
behavior of miter joins when the joins angle is sharp. This value must
be >= 0.
@param miter    set the miter limit on the paint, used whenever the
				paint's style is Stroke or StrokeAndFill.
*/
func (paint *Paint) SetStrokeMiter(miter Scalar) {
	toimpl()
}

/** Cap enum specifies the settings for the paint's strokecap. This is the
treatment that is applied to the beginning and end of each non-closed
contour (e.g. lines).

If the cap is round or square, the caps are drawn when the contour has
a zero length. Zero length contours can be created by following moveTo
with a lineTo at the same point, or a moveTo followed by a close.

A dash with an on interval of zero also creates a zero length contour.

The zero length contour draws the square cap without rotation, since
the no direction can be inferred. */
type PaintCap int

const (
	KPaintCapButt   = PaintCap(iota) //< begin/end contours with no extension
	KPaintCapRound                   //< begin/end contours with a semi-circle extension
	KPaintCapSquare                  //< begin/end contours with a half square extension

	KPaintCapLast    = KPaintCapSquare
	KPaintCapDefault = KPaintCapButt
	KPaintCapCount   = KPaintCapLast + 1
)

/** Return the paint's stroke cap type, controlling how the start and end
of stroked lines and paths are treated.
@return the line cap style for the paint, used whenever the paint's
		style is Stroke or StrokeAndFill.
*/
func (paint *Paint) StrokeCap() PaintCap {
	toimpl()
	return KPaintCapDefault
}

/** Set the paint's stroke cap type.
@param cap  set the paint's line cap style, used whenever the paint's
			style is Stroke or StrokeAndFill. */
func (paint *Paint) SetStrokeCap(cap PaintCap) {
	toimpl()
}

type PaintJoin int

const (
	KPaintJoinMiter = PaintJoin(iota) //< connect path segments with a sharp join
	KPaintJoinRound                   //< connect path segments with a round join
	KPaintJoinBevel                   //< connect path segments with a flat bevel join

	KPaintJoinLast    = KPaintJoinBevel
	KPaintJoinDefault = KPaintJoinMiter
	KPaintJoinCount   = KPaintJoinLast + 1
)

/** Return the paint's stroke join type.
@return the paint's line join style, used whenever the paint's style is
		Stroke or StrokeAndFill. */
func (paint *Paint) StrokeJoin() PaintJoin {
	toimpl()
	return KPaintJoinDefault
}

/** Set the paint's stroke join type.
@param join set the paint's line join style, used whenever the paint's
			style is Stroke or StrokeAndFill. */
func (paint *Paint) SetStrokeJoin(join PaintJoin) {
	toimpl()
}

/**
 *  Applies any/all effects (patheffect, stroking) to src, returning the
 *  result in dst. The result is that drawing src with this paint will be
 *  the same as drawing dst with a default paint (at least from the
 *  geometric perspective).
 *
 *  @param src  input path
 *  @param dst  output path (may be the same as src)
 *  @param cullRect If not null, the dst path may be culled to this rect.
 *  @param resScale If > 1, increase precision, else if (0 < res < 1) reduce precision
 *              in favor of speed/size.
 *  @return     true if the path should be filled, or false if it should be
 *              drawn with a hairline (width == 0)
 */
func (paint *Paint) FillPath(src *Path, dst *Path, cullRect *Rect, resScale Scalar) {
	toimpl()
}

/** Get the paint's shader object.
	<p />
  The shader's reference count is not affected.
	@return the paint's shader (or NULL)
*/
func (paint *Paint) Shader() *Shader {
	toimpl()
	return nil
}

/** Set or clear the shader object.
 *  Shaders specify the source color(s) for what is being drawn. If a paint
 *  has no shader, then the paint's color is used. If the paint has a
 *  shader, then the shader's color(s) are use instead, but they are
 *  modulated by the paint's alpha. This makes it easy to create a shader
 *  once (e.g. bitmap tiling or gradient) and then change its transparency
 *  w/o having to modify the original shader... only the paint's alpha needs
 *  to be modified.
 *
 *  There is an exception to this only-respect-paint's-alpha rule: If the shader only generates
 *  alpha (e.g. SkShader::CreateBitmapShader(bitmap, ...) where bitmap's colortype is kAlpha_8)
 *  then the shader will use the paint's entire color to "colorize" its output (modulating the
 *  bitmap's alpha with the paint's color+alpha).
 *
 *  Pass NULL to clear any previous shader.
 *  As a convenience, the parameter passed is also returned.
 *  If a previous shader exists, its reference count is decremented.
 *  If shader is not NULL, its reference count is incremented.
 *  @param shader   May be NULL. The shader to be installed in the paint
 *  @return         shader
 */
func (paint *Paint) SetShader(shader *Shader) {
	toimpl()
}

/** Get the paint's colorfilter. If there is a colorfilter, its reference
count is not changed.
@return the paint's colorfilter (or NULL)
*/
func (paint *Paint) ColorFilter() *ColorFilter {
	return paint.colorFilter
}

/** Set or clear the paint's colorfilter, returning the parameter.
<p />
If the paint already has a filter, its reference count is decremented.
If filter is not NULL, its reference count is incremented.
@param filter   May be NULL. The filter to be installed in the paint
@return         filter
*/
func (paint *Paint) SetColorFilter(colorFilter *ColorFilter) {
	toimpl()
}

/** Get the paint's xfermode object.
	<p />
  The xfermode's reference count is not affected.
	@return the paint's xfermode (or NULL)
*/
func (paint *Paint) Xfermode() *Xfermode {
	return paint.xfermode
}

/** Set or clear the xfermode object.
<p />
Pass NULL to clear any previous xfermode.
As a convenience, the parameter passed is also returned.
If a previous xfermode exists, its reference count is decremented.
If xfermode is not NULL, its reference count is incremented.
@param xfermode May be NULL. The new xfermode to be installed in the
				paint
@return         xfermode
*/
func (paint *Paint) SetXfermode(xfermode *Xfermode) {
	paint.xfermode = xfermode
}

/** Create an xfermode based on the specified Mode, and assign it into the
paint, returning the mode that was set. If the Mode is SrcOver, then
the paint's xfermode is set to null.
*/
func (paint *Paint) SetXfermodeMode(mode XfermodeMode) *Xfermode {
	paint.xfermode = NewXfermodeWithMode(mode)
	return paint.xfermode
}

/** Get the paint's patheffect object.
	<p />
  The patheffect reference count is not affected.
	@return the paint's patheffect (or NULL)
*/
func (paint *Paint) PathEffect() *PathEffect {
	toimpl()
	return nil
}

/** Set or clear the patheffect object.
<p />
Pass NULL to clear any previous patheffect.
As a convenience, the parameter passed is also returned.
If a previous patheffect exists, its reference count is decremented.
If patheffect is not NULL, its reference count is incremented.
@param effect   May be NULL. The new patheffect to be installed in the
				paint
@return         effect
*/
func (paint *Paint) SetPathEffect(effect *PathEffect) {
	toimpl()
}

/** Get the paint's maskfilter object.
	<p />
  The maskfilter reference count is not affected.
	@return the paint's maskfilter (or NULL)
*/
func (paint *Paint) MaskFilter() *MaskFilter {
	toimpl()
	return nil
}

/** Set or clear the maskfilter object.
<p />
Pass NULL to clear any previous maskfilter.
As a convenience, the parameter passed is also returned.
If a previous maskfilter exists, its reference count is decremented.
If maskfilter is not NULL, its reference count is incremented.
@param maskfilter   May be NULL. The new maskfilter to be installed in
					the paint
@return             maskfilter
*/
func (paint *Paint) SetMastFilter(maskfilter *MaskFilter) {
	toimpl()
}

// These attributes are for text/fonts

/** Get the paint's typeface object.
<p />
The typeface object identifies which font to use when drawing or
measuring text. The typeface reference count is not affected.
@return the paint's typeface (or NULL)
*/
func (paint *Paint) Typeface() *Typeface {
	toimpl()
	return nil
}

/** Set or clear the typeface object.
<p />
Pass NULL to clear any previous typeface.
As a convenience, the parameter passed is also returned.
If a previous typeface exists, its reference count is decremented.
If typeface is not NULL, its reference count is incremented.
@param typeface May be NULL. The new typeface to be installed in the
				paint
@return         typeface
*/
func (paint *Paint) SetTypeface(typeface *Typeface) {
	toimpl()
}

/** Get the paint's rasterizer (or NULL).
<p />
The raster controls how paths/text are turned into alpha masks.
@return the paint's rasterizer (or NULL)
*/
func (paint *Paint) Rasterizer() *Rasterizer {
	toimpl()
	return nil
}

/** Set or clear the rasterizer object.
<p />
Pass NULL to clear any previous rasterizer.
As a convenience, the parameter passed is also returned.
If a previous rasterizer exists in the paint, its reference count is
decremented. If rasterizer is not NULL, its reference count is
incremented.
@param rasterizer May be NULL. The new rasterizer to be installed in
				  the paint.
@return           rasterizer
*/
func (paint *Paint) SetRasterizer(rasterizer *Rasterizer) {
	toimpl()
}

func (paint *Paint) ImageFilter() *ImageFilter {
	return paint.imageFilter
}

func (paint *Paint) SetImageFilter(imageFilter *ImageFilter) {
	toimpl()
}

/**
 *  Return the paint's SkDrawLooper (if any). Does not affect the looper's
 *  reference count.
 */
func (paint *Paint) Looper() *DrawLooper {
	return paint.looper
}

/**
 *  Set or clear the looper object.
 *  <p />
 *  Pass NULL to clear any previous looper.
 *  As a convenience, the parameter passed is also returned.
 *  If a previous looper exists in the paint, its reference count is
 *  decremented. If looper is not NULL, its reference count is
 *  incremented.
 *  @param looper May be NULL. The new looper to be installed in the paint.
 *  @return looper
 */
func (paint *Paint) SetLooper(looper *DrawLooper) {
	paint.looper = looper
}

type PaintAlign int

const (
	KPaintAlignLeft = PaintAlign(iota)
	KPaintAlignCenter
	KPaintAlignRight
	KPaintAlignCount = 3
)

/** Return the paint's Align value for drawing text.
@return the paint's Align value for drawing text.
*/
func (paint *Paint) TextAlign() PaintAlign {
	toimpl()
	return KPaintAlignLeft
}

/** Set the paint's text alignment.
@param align set the paint's Align value for drawing text.
*/
func (paint *Paint) SetTextAlign(align PaintAlign) {
	toimpl()
}

/** Return the paint's text size.
@return the paint's text size.
*/
func (paint *Paint) TextSize() Scalar {
	toimpl()
	return Scalar(0)
}

/** Set the paint's text size. This value must be > 0
@param textSize set the paint's text size.
*/
func (paint *Paint) SetTextSize(textSize Scalar) {
	toimpl()
}

/** Return the paint's horizontal scale factor for text. The default value
is 1.0.
@return the paint's scale factor in X for drawing/measuring text
*/
func (paint *Paint) TextScaleX() Scalar {
	toimpl()
	return Scalar(0)
}

/** Set the paint's horizontal scale factor for text. The default value
is 1.0. Values > 1.0 will stretch the text wider. Values < 1.0 will
stretch the text narrower.
@param scaleX   set the paint's scale factor in X for drawing/measuring
				text.
*/
func (paint *Paint) SetTextScaleX(scaleX Scalar) {
	toimpl()
}

/** Return the paint's horizontal skew factor for text. The default value
is 0.
@return the paint's skew factor in X for drawing text.
*/
func (paint *Paint) TextSkewX() Scalar {
	toimpl()
	return Scalar(0)
}

/** Set the paint's horizontal skew factor for text. The default value
is 0. For approximating oblique text, use values around -0.25.
@param skewX set the paint's skew factor in X for drawing text.
*/
func (paint *Paint) SetTextSkewX(skewX Scalar) {
	toimpl()
}

/** Describes how to interpret the text parameters that are passed to paint
methods like measureText() and getTextWidths().
*/
type PaintTextEncoding int

const (
	KPaintTextEncodingUTF8    = PaintTextEncoding(iota) //< the text parameters are UTF8
	KPaintTextEncodingUTF16                             //< the text parameters are UTF16
	KPaintTextEncodingUTF32                             //< the text parameters are UTF32
	KPaintTextEncodingGlyphID                           //< the text parameters are glyph indices
)

func (paint *Paint) TextEncoding() PaintTextEncoding {
	toimpl()
	return KPaintTextEncodingUTF8
}

func (paint *Paint) SetTextEncoding(encoding PaintTextEncoding) {
	toimpl()
}

/** Flags which indicate the confidence level of various metrics.
A set flag indicates that the metric may be trusted.
*/
type PaintFontMetricsFlags int

const (
	KPaintFontMetricsFlagUnderlineThinknessIsValid = 1 << iota
	KPaintFontMetricsFlagUnderlinePositionIsValid
)

type PaintFontMetrics struct {
	Flags              uint32 //< Bit field to identify which values are unknown
	Top                Scalar //< The greatest distance above the baseline for any glyph (will be <= 0)
	Ascent             Scalar //< The recommended distance above the baseline (will be <= 0)
	Descent            Scalar //< The recommended distance below the baseline (will be >= 0)
	Bottom             Scalar //< The greatest distance below the baseline for any glyph (will be >= 0)
	Leading            Scalar //< The recommended distance to add between lines of text (will be >= 0)
	AvgCharWidth       Scalar //< the average character width (>= 0)
	MaxCharWidth       Scalar //< the max character width (>= 0)
	XMin               Scalar //< The minimum bounding box x value for all glyphs
	XMax               Scalar //< The maximum bounding box x value for all glyphs
	XHeight            Scalar //< The height of an 'x' in px, or 0 if no 'x' in face
	CapHeight          Scalar //< The cap height (> 0), or 0 if cannot be determined.
	UnderlineThickness Scalar //< underline thickness, or 0 if cannot be determined

	/**  Underline Position - position of the top of the Underline stroke
	  relative to the baseline, this can have following values
	  - Negative - means underline should be drawn above baseline.
	  - Positive - means below baseline.
	  - Zero     - mean underline should be drawn on baseline. */
	UnderlinePosition Scalar
}

/**  If the fontmetrics has a valid underlinethickness, return true, and set the
		thickness param to that value. If it doesn't return false and ignore the
		thickness param.
*/
func (metrics *PaintFontMetrics) HasUnderlineThickness(thickness *Scalar) bool {
	toimpl()
	return false
}

/**  If the fontmetrics has a valid underlineposition, return true, and set the
		thickness param to that value. If it doesn't return false and ignore the
		thickness param.
*/
func (metrics *PaintFontMetrics) HasUnderlinePosition(position *Scalar) bool {
	toimpl()
	return false
}

/** Return the recommend spacing between lines (which will be
	fDescent - fAscent + fLeading).
	If metrics is not null, return in it the font metrics for the
	typeface/pointsize/etc. currently set in the paint.
	@param metrics      If not null, returns the font metrics for the
						current typeface/pointsize/etc setting in this
						paint.
	@param scale        If not 0, return width as if the canvas were scaled
						by this value
	@param return the recommended spacing between lines
*/
func (paint *Paint) FontMetrics(metrics *PaintFontMetrics, scale Scalar) Scalar {
	toimpl()
	return Scalar(0)
}

/** Return the recommend line spacing. This will be
	fDescent - fAscent + fLeading
*/
func (paint *Paint) FontSpacing() Scalar {
	toimpl()
	return Scalar(0)
}

/** Convert the specified text into glyph IDs, returning the number of
	glyphs ID written. If glyphs is NULL, it is ignore and only the count
	is returned.
*/
func (paint *Paint) TextToGlyphs(text string, byteLength int, glyphs []GlyphID) int {
	toimpl()
	return 0
}

/** Return true if all of the specified text has a corresponding non-zero
	glyph ID. If any of the code-points in the text are not supported in
	the typeface (i.e. the glyph ID would be zero), then return false.

	If the text encoding for the paint is kGlyph_TextEncoding, then this
	returns true if all of the specified glyph IDs are non-zero.
 */
func (paint *Paint) ContainsText(text string, byteLength int) bool {
	toimpl()
	return false
}

/** Convert the glyph array into Unichars. Unconvertable glyphs are mapped
	to zero. Note: this does not look at the text-encoding setting in the
	paint, only at the typeface.
*/
func (paint *Paint) GlyphsToUnichars(glyphs []GlyphID, count int, text []Unichar) int {
	toimpl()
	return Unichar(0)
}

/** Return the number of drawable units in the specified text buffer.
	This looks at the current TextEncoding field of the paint. If you also
	want to have the text converted into glyph IDs, call textToGlyphs
	instead.
*/
func (paint *Paint) CountText(text string, byteLength int) int {
	toimpl()
	return 0
}

// ------

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

func (paint *Paint) NothingToDraw() bool {
	toimpl()
	return false
}

func (paint *Paint) Clone() *Paint {
	toimpl()
	return nil
}

func (paint *Paint) SetShader(shader *Shader) {
	toimpl()
}
