package ggk

import (
	"container/list"
	"time"
)

type CanvasInitFlags int

const (
	KCanvasInitFlagDefault = CanvasInitFlags(1 << iota)
	KCanvasInitFlagConservativeRasterClip
)

type CanvasShaderOverrideOpacity int

const (
	KCanvasShaderOverrideOpacityNone      = CanvasShaderOverrideOpacity(1 << iota) //!< there is no overriding shader (bitmap or image)
	KCanvasShaderOverrideOpacityOpaque                                             //!< the overriding shader is opaque
	KCanvasShaderOverrideOpacityNotOpaque                                          //!< the overriding shader may not be opaque
)

/** \class SkCanvas

A Canvas encapsulates all of the state about drawing into a device (bitmap).
This includes a reference to the device itself, and a stack of matrix/clip
values. For any given draw call (e.g. drawRect), the geometry of the object
being drawn is transformed by the concatenation of all the matrices in the
stack. The transformed geometry is clipped by the intersection of all of
the clips in the stack.

While the Canvas holds the state of the drawing device, the state (style)
of the object being drawn is held by the Paint, which is provided as a
parameter to each of the draw() methods. The Paint holds attributes such as
color, typeface, textSize, strokeWidth, shader (e.g. gradients, patterns),
etc.
*/
type Canvas struct {
	surfaceProps SurfaceProps
	saveCount    int
	metaData     *MetaData
	baseSurface  *BaseSurface
	mcStack      *list.List
	clipStack    *ClipStack
	mcRec        *tCanvasMCRec // points to top of the stack

	deviceCMDirty bool

	cachedLocalClipBounds      Rect
	cachedLocalClipBoundsDirty bool

	allowSoftClip          bool
	allowSimplifyClip      bool
	conservativeRasterClip bool

	isScaleTranslate bool
	deviceClipBounds Rect
}

/**
Attempt to allocate raster canvas, matching the ImageInfo, that will draw directly into the
specified pixels. To access the pixels after drawing to them, the caller should call
flush() or call peekPixels(...).

On failure, return NULL. This can fail for several reasons:
1. invalid ImageInfo (e.g. negative dimensions)
2. unsupported ImageInfo for a canvas
    - kUnknown_SkColorType, kIndex_8_SkColorType
    - kUnknown_SkAlphaType
    - this list is not complete, so others may also be unsupported

Note: it is valid to request a supported ImageInfo, but with zero
dimensions.
 */
func NewCanvasRasterDirect(imageInfo *ImageInfo, pixels []byte, rowBytes int) *Canvas {
	toimpl()
	return &Canvas{}
}

func NewCanvasRasterDirectN32(width, height int, pixels []PremulColor, rowBytes int) *Canvas {
	toimpl()
	return &Canvas{}
}

/**
 *  Creates an empty canvas with no backing device/pixels, and zero
 *  dimensions.
 */
func NewCanvasEmpty() *Canvas {
	toimpl()
	return &Canvas{}
}

/**
Construct a canvas with the specified bitmap to draw into.
@param bitmap   Specifies a bitmap for the canvas to draw into. Its
                structure are copied to the canvas.
*/
func NewCanvasFromBitmap(bmp *Bitmap) *Canvas {
	var canvas = new(Canvas)
	canvas.surfaceProps = MakeSurfaceProps(KSurfacePropsFlagNone, KSurfacePropsInitTypeLegacyFontHost)
	canvas.mcStack = list.New()
	var device = NewBitmapDevice(bmp, canvas.surfaceProps)
	canvas.init(device.BaseDevice, KCanvasInitFlagDefault)
	return canvas
}

/**
 *  Creates a canvas of the specified dimensions, but explicitly not backed
 *  by any device/pixels. Typically this use used by subclasses who handle
 *  the draw calls in some other way.
 */
func NewCanvas(width, height int, surfaceProps *SurfaceProps) *Canvas {
	toimpl()
	return &Canvas{}
}

/** Construct a canvas with the specified device to draw into.

	@param device   Specifies a device for the canvas to draw into.
*/
func NewCanvasFromDevice(device *BaseDevice) *Canvas {
	toimpl()
	return &Canvas{}
}

/** Construct a canvas with the specified bitmap to draw into.
	@param bitmap   Specifies a bitmap for the canvas to draw into. Its
					structure are copied to the canvas.
	@param props    New canvas surface properties.
*/
func NewCanvasFromBitmapSurfaceProps(bmp *Bitmap, surfaceProps *SurfaceProps) *Canvas {
	toimpl()
	return &Canvas{}
}

func (canvas *Canvas) MetaData() *MetaData {
	toimpl()
	return nil
}

/**
 *  Return ImageInfo for this canvas. If the canvas is not backed by pixels
 *  (cpu or gpu), then the info's ColorType will be kUnknown_SkColorType.
 */
func (canvas *Canvas) ImageInfo() *ImageInfo {
	toimpl()
	return nil
}

/**
 *  If the canvas is backed by pixels (cpu or gpu), this writes a copy of the SurfaceProps
 *  for the canvas to the location supplied by the caller, and returns true. Otherwise,
 *  return false and leave the supplied props unchanged.
 */
func (canvas *Canvas) SurfaceProps() *SurfaceProps {
	toimpl()
	return canvas.surfaceProps
}

/**
 *  Trigger the immediate execution of all pending draw operations. For the GPU
 *  backend this will resolve all rendering to the GPU surface backing the
 *  SkSurface that owns this canvas.
 */
func (canvas *Canvas) flush() {
	toimpl()
}

/**
 * Gets the size of the base or root layer in global canvas coordinates. The
 * origin of the base layer is always (0,0). The current drawable area may be
 * smaller (due to clipping or saveLayer).
 * TODO(abstract)
 */
func (c *Canvas) BaseLayerSize() Size {
	var device = c.Device()
	var size Size
	if device != nil {
		size = MakeSize(device.Width(), device.Height())
	}
	return size
}

/**
 *  Return the canvas' device object, which may be null. The device holds
 *  the bitmap of the pixels that the canvas draws into. The reference count
 *  of the returned device is not changed by this call.
 */
func (c *Canvas) Device() *BaseDevice {
	var rec = c.mcStack.Front().Value.(*tCanvasMCRec)
	return rec.Layer.Device
}

/**
 *  saveLayer() can create another device (which is later drawn onto
 *  the previous device). getTopDevice() returns the top-most device current
 *  installed. Note that this can change on other calls like save/restore,
 *  so do not access this device after subsequent canvas calls.
 *  The reference count of the device is not changed.
 *
 * @param updateMatrixClip If this is true, then before the device is
 *        returned, we ensure that its has been notified about the current
 *        matrix and clip. Note: this happens automatically when the device
 *        is drawn to, but is optional here, as there is a small perf hit
 *        sometimes.
 */
func (canvas *Canvas) TopDevice() *BaseDevice {
	toimpl()
	return nil
}

/**
 *  Create a new surface matching the specified info, one that attempts to
 *  be maximally compatible when used with this canvas. If there is no matching Surface type,
 *  NULL is returned.
 *
 *  If surfaceprops is specified, those are passed to the new surface, otherwise the new surface
 *  inherits the properties of the surface that owns this canvas. If this canvas has no parent
 *  surface, then the new surface is created with default properties.
 */
func (canvas *Canvas) NewSurface(imageInfo *ImageInfo, surfaceProps *SurfaceProps) *Surface {
	toimpl()
	return nil
}

/**
 * Return the GPU context of the device that is associated with the canvas.
 * For a canvas with non-GPU device, NULL is returned.
 */
func (canvas *Canvas) GrContext() *GrContext {
	toimpl()
	return nil
}

/**
 *  If the canvas has writable pixels in its top layer (and is not recording to a picture
 *  or other non-raster target) and has direct access to its pixels (i.e. they are in
 *  local RAM) return the address of those pixels, and if not null,
 *  return the ImageInfo, rowBytes and origin. The returned address is only valid
 *  while the canvas object is in scope and unchanged. Any API calls made on
 *  canvas (or its parent surface if any) will invalidate the
 *  returned address (and associated information).
 *
 *  On failure, returns NULL and the info, rowBytes, and origin parameters are ignored.
 */
func (canvas *Canvas) AccessTopLayerPixels(imageInfo *ImageInfo, rowBytes *int, origin *Point) []byte {
	toimpl()
	return nil
}

/**
 *  If the canvas has readable pixels in its base layer (and is not recording to a picture
 *  or other non-raster target) and has direct access to its pixels (i.e. they are in
 *  local RAM) return true, and if not null, return in the pixmap parameter information about
 *  the pixels. The pixmap's pixel address is only valid
 *  while the canvas object is in scope and unchanged. Any API calls made on
 *  canvas (or its parent surface if any) will invalidate the pixel address
 *  (and associated information).
 *
 *  On failure, returns false and the pixmap parameter will be ignored.
 */
func (canvas *Canvas) PeekPixels(pixmap *Pixmap) bool {
	toimpl()
	return false
}

/**
ReadPixels copy the pixels from the base-layer into the specified buffer
(pixels + rowBytes). converting them into the requested format (ImageInfo).
The base-layer are read starting at the specified (srcX, srcY) location in
the coordinate system of the base-layer.

The specified ImageInfo and (srcX, srcY) offset specifies a source rectangle.

    srcR.SetXYWH(srcX, srcY, dstInfo.Width(), dstInfo.Height())

srcR is intersected with the bounds of the base-layer. If this intersection
is not empty, then we have two sets of pixels (of equal size). Replace the
dst pixels with the corresponding src pixels, performing any
colortype/alphatype transformations needed (in the case where the src and dst
have different colortypes or alphatypes).

This call can fail, returning false, for serveral reasons:
- If srcR does not intersect the base-layer bounds.
- If the requested colortype/alphatype cannot be converted from the base-layer's types.
- If this canvas is not backed by pixels (e.g. picture or PDF)
*/
func (c *Canvas) ReadPixels(dstInfo *ImageInfo, dstData []byte, rowBytes int, x, y Scalar) error {
	var dev = c.Device()
	if dev == nil {
		return errorf("device is nil")
	}
	var size = c.BaseLayerSize()
	var rec = newReadPixelsRec(dstInfo, dstData, rowBytes, x, y)
	if err := rec.Trim(size.Width(), size.Height()); err != nil {
		return errorf("bad arg %v", err)
	}
	// The device can assert that the requested area is always contained in its
	// bounds.
	return dev.ReadPixels(rec.Info, rec.Pixels, rec.RowBytes, rec.X, rec.Y)
}

/**
 *  Helper for calling readPixels(info, ...). This call will check if bitmap has been allocated.
 *  If not, it will attempt to call allocPixels(). If this fails, it will return false. If not,
 *  it calls through to readPixels(info, ...) and returns its result.
 */
func (canvas *Canvas) ReadPixelsToBitmap(bmp *Bitmap, srcX, srcY int) bool {
	toimpl()
	return false
}

/**
 *  Helper for allocating pixels and then calling readPixels(info, ...). The bitmap is resized
 *  to the intersection of srcRect and the base-layer bounds. On success, pixels will be
 *  allocated in bitmap and true returned. On failure, false is returned and bitmap will be
 *  set to empty.
 */
func (canvas *Canvas) ReadPixelsInRectToBitmap(rect Rect, bmp *Bitmap) bool {
	toimpl()
	return false
}

/**
 *  This method affects the pixels in the base-layer, and operates in pixel coordinates,
 *  ignoring the matrix and clip.
 *
 *  The specified ImageInfo and (x,y) offset specifies a rectangle: target.
 *
 *      target.setXYWH(x, y, info.width(), info.height());
 *
 *  Target is intersected with the bounds of the base-layer. If this intersection is not empty,
 *  then we have two sets of pixels (of equal size), the "src" specified by info+pixels+rowBytes
 *  and the "dst" by the canvas' backend. Replace the dst pixels with the corresponding src
 *  pixels, performing any colortype/alphatype transformations needed (in the case where the
 *  src and dst have different colortypes or alphatypes).
 *
 *  This call can fail, returning false, for several reasons:
 *  - If the src colortype/alphatype cannot be converted to the canvas' types
 *  - If this canvas is not backed by pixels (e.g. picture or PDF)
 */
func (canvas *Canvas) WritePixels(info *ImageInfo, pixels []byte, rowBytes int, x, y int) bool {
	toimpl()
	return false
}

/**
 *  Helper for calling writePixels(info, ...) by passing its pixels and rowbytes. If the bitmap
 *  is just wrapping a texture, returns false and does nothing.
 */
func (canvas *Canvas) WritePixelsFromBitmap(bmp *Bitmap, x, y int) bool {
	toimpl()
	return false
}

/**
This call saves the current matrix, clip, and drawFilter, and pushes a
copy onto a private stack. Subsequent calls to translate, scale,
rotate, skew, concat or clipRect, clipPath, and setDrawFilter all
operate on this copy.
When the balancing call to restore() is made, the previous matrix, clip,
and drawFilter are restored.

@return The value to pass to restoreToCount() to balance this save()
*/
func (canvas *Canvas) Save() {
	toimpl()
}

/**
This behaves the same as save(), but in addition it allocates an
offscreen bitmap. All drawing calls are directed there, and only when
the balancing call to restore() is made is that offscreen transfered to
the canvas (or the previous layer).
@param bounds (may be null) This rect, if non-null, is used as a hint to
			  limit the size of the offscreen, and thus drawing may be
			  clipped to it, though that clipping is not guaranteed to
			  happen. If exact clipping is desired, use clipRect().
@param paint (may be null) This is copied, and is applied to the
			 offscreen when restore() is called
@return The value to pass to restoreToCount() to balance this save()
*/
func (canvas *Canvas) SaveLayer(bounds *Rect, paint *Paint) int {
	toimpl()
	return 0
}

/**
 *  Temporary name.
 *  Will allow any requests for LCD text to be respected, so the caller must be careful to
 *  only draw on top of opaque sections of the layer to get good results.
 */
func SaveLayerPreserveLCDTextRequests(bounds *Rect, paint *Paint) int {
	toimpl()
	return 0
}

/** This behaves the same as save(), but in addition it allocates an
	offscreen bitmap. All drawing calls are directed there, and only when
	the balancing call to restore() is made is that offscreen transfered to
	the canvas (or the previous layer).
	@param bounds (may be null) This rect, if non-null, is used as a hint to
				  limit the size of the offscreen, and thus drawing may be
				  clipped to it, though that clipping is not guaranteed to
				  happen. If exact clipping is desired, use clipRect().
	@param alpha  This is applied to the offscreen when restore() is called.
	@return The value to pass to restoreToCount() to balance this save()
*/
func (canvas *Canvas) SaveLayerAlphas(bounds *Rect, alpha uint8) int {
	toimpl()
	return 0
}

type CanvasSaveLayerFlags int

const (
	KCanvasSaveLayerFlagIsOpaque = 1 << iota
	KCanvasSaveLayerFlagPreserveLCDText
	kCanvasSaveLayerFlagDontClipToLayer   // private
	KCanvasSaveLayerDontClipToLayerLegacy = kCanvasSaveLayerFlagDontClipToLayer
)

type CanvasSaveLayerRec struct {
	bounds         *Rect
	paint          *Paint
	backdrop       *ImageFilter
	saveLayerFlags CanvasSaveLayerFlags
}

func newCanvasSaveLayerRec(bounds *Rect, paint *Paint, backdrop *ImageFilter,
saveLayerFlags CanvasSaveLayerFlags) *CanvasSaveLayerRec {
	return &CanvasSaveLayerRec{
		bounds:         bounds,
		paint:          paint,
		backdrop:       backdrop,
		saveLayerFlags: saveLayerFlags,
	}
}

func (canvas *Canvas) SaveLayerWithRec(rec *CanvasSaveLayerRec) int {
	toimpl()
	return 0
}

/**
This call balances a previous call to save(), and is used to remove all
modifications to the matrix/clip/drawFilter state since the last save
call.
It is an error to call restore() more times than save() was called.
*/
func (canvas *Canvas) Restore() {
	toimpl()
}

/**
Returns the number of matrix/clip states on the SkCanvas' private stack.
This will equal # save() calls - # restore() calls + 1. The save count on
a new canvas is 1.
*/
func (canvas *Canvas) SaveCount() int {
	return canvas.saveCount
}

/**
Efficient way to pop any calls to save() that happened after the save
count reached saveCount. It is an error for saveCount to be greater than
getSaveCount(). To pop all the way back to the initial matrix/clip context
pass saveCount == 1.
@param saveCount    The number of save() levels to restore from
*/
func (canvas *Canvas) RestoreToCount(saveCount int) {
	toimpl()
	return
}

/**
Preconcat the current matrix with the specified translation
@param dx   The distance to translate in X
@param dy   The distance to translate in Y
*/
func (canvas *Canvas) Translate(dx, dy Scalar) {
	toimpl()
}

/** Preconcat the current matrix with the specified scale.
	@param sx   The amount to scale in X
	@param sy   The amount to scale in Y
*/
func (canvas *Canvas) Scale(sx, sy Scalar) {
	toimpl()
}

/** Preconcat the current matrix with the specified rotation about the origin.
	@param degrees  The amount to rotate, in degrees
*/
func (canvas *Canvas) Rotate(degrees Scalar) {
	toimpl()
}

/** Preconcat the current matrix with the specified rotation about a given point.
	@param degrees  The amount to rotate, in degrees
	@param px  The x coordinate of the point to rotate about.
	@param py  The y coordinate of the point to rotate about.
*/
func (canvas *Canvas) RotateAt(degrees, px, py Scalar) {
	toimpl()
}

/** Preconcat the current matrix with the specified skew.
	@param sx   The amount to skew in X
	@param sy   The amount to skew in Y
*/
func (canvas *Canvas) Skew(sx, sy Scalar) {
	toimpl()
}

/** Preconcat the current matrix with the specified matrix.
	@param matrix   The matrix to preconcatenate with the current matrix
*/
func (canvas *Canvas) Concat(matrix *Matrix) {
	toimpl()
}

/** Replace the current matrix with a copy of the specified matrix.
	@param matrix The matrix that will be copied into the current matrix.
*/
func (canvas *Canvas) SetMatrix(matrix *Matrix) {
	toimpl()
}

/** Helper for setMatrix(identity). Sets the current matrix to identity.
*/
func (canvas *Canvas) ResetMatrix() {
	toimpl()
}

/** Add the specified translation to the current draw depth of the canvas.
	@param z    The distance to translate in Z.
				Negative into screen, positive out of screen.
				Without translation, the draw depth defaults to 0.
*/
func (canvas *Canvas) TranslateZ(z Scalar) {
	toimpl()
}

/** Set the current set of lights in the canvas.
	@param lights   The lights that we want the canvas to have.
*/
func (canvas *Canvas) SetLights(lights *Lights) {
	toimpl()
}

/** Returns the current set of lights the canvas uses
*/
func (canvas *Canvas) Lights() *Lights {
	return nil
}

/**
 *  Modify the current clip with the specified rectangle.
 *  @param rect The rect to combine with the current clip
 *  @param op The region op to apply to the current clip
 *  @param doAntiAlias true if the clip should be antialiased
 */
func (canvas *Canvas) ClipRect(rect Rect, op RegionOp, doAntiAlias bool) {
	toimpl()
}

/**
 *  Modify the current clip with the specified path.
 *  @param path The path to combine with the current clip
 *  @param op The region op to apply to the current clip
 *  @param doAntiAlias true if the clip should be antialiased
 */
func (canvas *Canvas) ClipPath(path *Path, op RegionOp, doAntiAlias bool) {
	toimpl()
}

/**
EXPERIMENTAL -- only used for testing
Set to false to force clips to be hard, even if doAntiAlias=true is
passed to clipRect or clipPath.
*/
func (canvas *Canvas) SetAllowSoftClip(allow bool) {
	toimpl()
}

/**
EXPERIMENTAL -- only used for testing
Set to simplify clip stack using path ops.
*/
func (canvas *Canvas) SetAllowSimplifyClip(allow bool) {
	toimpl()
}

/** Modify the current clip with the specified region. Note that unlike
	clipRect() and clipPath() which transform their arguments by the current
	matrix, clipRegion() assumes its argument is already in device
	coordinates, and so no transformation is performed.
	@param deviceRgn    The region to apply to the current clip
	@param op The region op to apply to the current clip
*/
func (canvas *Canvas) ClipRegion(deviceRgn *Region, op RegionOp) {
	toimpl()
}

/** Helper for clipRegion(rgn, kReplace_Op). Sets the current clip to the
	specified region. This does not intersect or in any other way account
	for the existing clip region.
	@param deviceRgn The region to copy into the current clip.
*/
func (canvas *Canvas) SetClipRegion(deviceRgn *Region) {
	toimpl()
}

/**
Return true if the specified rectangle, after being transformed by the
current matrix, would lie completely outside of the current clip. Call
this to check if an area you intend to draw into is clipped out (and
therefore you can skip making the draw calls).
@param rect the rect to compare with the current clip
@return true if the rect (transformed by the canvas' matrix) does not
			 intersect with the canvas' clip
*/
func (canvas *Canvas) QuickRejectRect(rect Rect) bool {
	toimpl()
	return false
}

/**
Return true if the specified path, after being transformed by the
current matrix, would lie completely outside of the current clip. Call
this to check if an area you intend to draw into is clipped out (and
therefore you can skip making the draw calls). Note, for speed it may
return false even if the path itself might not intersect the clip
(i.e. the bounds of the path intersects, but the path does not).
@param path The path to compare with the current clip
@return true if the path (transformed by the canvas' matrix) does not
			 intersect with the canvas' clip
*/
func (canvas *Canvas) QuickRejectPath(path *Path) bool {
	toimpl()
	return false
}

/**
Return the bounds of the current clip (in local coordinates) in the
bounds parameter, and return true if it is non-empty. This can be useful
in a way similar to quickReject, in that it tells you that drawing
outside of these bounds will be clipped out.
TODO(abstract)
*/
func (canvas *Canvas) ClipBounds(bounds *Rect) bool {
	toimpl()
	return false
}

/**
Return the bounds of the current clip, in device coordinates; returns
true if non-empty. Maybe faster than getting the clip explicitly and
then taking its bounds.
TODO(abstract)
*/
func (canvas *Canvas) ClipDeviceBounds(bounds *Rect) bool {
	toimpl()
	return false
}

/** Fill the entire canvas' bitmap (restricted to the current clip) with the
	specified ARGB color, using the specified mode.
	@param a    the alpha component (0..255) of the color to fill the canvas
	@param r    the red component (0..255) of the color to fill the canvas
	@param g    the green component (0..255) of the color to fill the canvas
	@param b    the blue component (0..255) of the color to fill the canvas
	@param mode the mode to apply the color in (defaults to SrcOver)
*/
func (canvas *Canvas) DrawARGB(a, r, g, b uint8, mode XfermodeMode) {
	toimpl()
}

/**
Fill the entire canvas' bitmap (restricted to the current clip) with the
specified color and mode.
@param color    the color to draw with
@param mode the mode to apply the color in (defaults to SrcOver)
*/
func (canvas *Canvas) DrawColor(color Color, mode XfermodeMode) {
	var paint = NewPaint()
	paint.SetColor(color)
	if KXfermodeModeSrcOver == mode {
		paint.SetXfermodeMode(mode)
	}
	canvas.DrawPaint(paint)
}

/**
 *  Helper method for drawing a color in SRC mode, completely replacing all the pixels
 *  in the current clip with this color.
 */
func (canvas *Canvas) Clear(color Color) {
	toimpl()
}

/**
 * This makes the contents of the canvas undefined. Subsequent calls that
 * require reading the canvas contents will produce undefined results. Examples
 * include blending and readPixels. The actual implementation is backend-
 * dependent and one legal implementation is to do nothing. This method
 * ignores the current clip.
 *
 * This function should only be called if the caller intends to subsequently
 * draw to the canvas. The canvas may do real work at discard() time in order
 * to optimize performance on subsequent draws. Thus, if you call this and then
 * never draw to the canvas subsequently you may pay a perfomance penalty.
 */
func (canvas *Canvas) Discard() {
	toimpl()
}

/**
 *  Fill the entire canvas (restricted to the current clip) with the
 *  specified paint.
 *  @param paint    The paint used to fill the canvas
 */
func (canvas *Canvas) DrawPaint(paint *Paint) {
	canvas.OnDrawPaint(paint)
}

type CanvasPointMode int

const (
	KCanvasPointModePoints  CanvasPointMode = iota // DrawPoints draws each point separately
	KCanvasPointModeLines                          // DrawPoints draws each pair of points as a line segment
	KCanvasPointModePolygon                        // DrawPoints draws the array of points as a polygon
)

/**
Draw a series of points, interpreted based on the PointMode mode. For
all modes, the count parameter is interpreted as the total number of
points. For kLine mode, count/2 line segments are drawn.
For kPoint mode, each point is drawn centered at its coordinate, and its
size is specified by the paint's stroke-width. It draws as a square,
unless the paint's cap-type is round, in which the points are drawn as
circles.
For kLine mode, each pair of points is drawn as a line segment,
respecting the paint's settings for cap/join/width.
For kPolygon mode, the entire array is drawn as a series of connected
line segments.
Note that, while similar, kLine and kPolygon modes draw slightly
differently than the equivalent path built with a series of moveto,
lineto calls, in that the path will draw all of its contours at once,
with no interactions if contours intersect each other (think XOR
xfermode). drawPoints always draws each element one at a time.
@param mode     PointMode specifying how to draw the array of points.
@param count    The number of points in the array
@param pts      Array of points to draw
@param paint    The paint used to draw the points
*/
func (canvas *Canvas) DrawPoints(mode CanvasPointMode, count int, pts []Point, paint *Paint) {
	canvas.OnDrawPoints(mode, count, pts, paint)
}

/** Draws a single pixel in the specified color.
	@param x        The X coordinate of which pixel to draw
	@param y        The Y coordiante of which pixel to draw
	@param color    The color to draw
*/
func (canvas *Canvas) DrawPoint(x, y Scalar, paint *Paint) {
	var pt Point
	pt.X, pt.Y = x, y
	canvas.DrawPoints(KCanvasPointModePoints, 1, []Point{pt}, paint)
}

/**
Draw a line segment with the specified start and stop x,y coordinates,
using the specified paint. NOTE: since a line is always "framed", the
paint's Style is ignored.
@param x0    The x-coordinate of the start point of the line
@param y0    The y-coordinate of the start point of the line
@param x1    The x-coordinate of the end point of the line
@param y1    The y-coordinate of the end point of the line
@param paint The paint used to draw the line
*/
func (canvas *Canvas) DrawLine(x0, y0, x1, y1 Scalar, paint *Paint) {
	toimpl()
}

/**
Draw the specified rectangle using the specified paint. The rectangle
will be filled or stroked based on the Style in the paint.
@param rect     The rect to be drawn
@param paint    The paint used to draw the rect
*/
func (canvas *Canvas) DrawRect(rect Rect, paint *Paint) {
	toimpl()
}

/**
Draw the specified rectangle using the specified paint. The rectangle
will be filled or framed based on the Style in the paint.
@param left     The left side of the rectangle to be drawn
@param top      The top side of the rectangle to be drawn
@param right    The right side of the rectangle to be drawn
@param bottom   The bottom side of the rectangle to be drawn
@param paint    The paint used to draw the rect
*/
func (canvas *Canvas) DrawRectCoords(left, top, right, bottom Scalar, paint *Paint) {
	toimpl()
}

/**
Draw the specified oval using the specified paint. The oval will be
filled or framed based on the Style in the paint.
@param oval     The rectangle bounds of the oval to be drawn
@param paint    The paint used to draw the oval
*/
func (canvas *Canvas) DrawOval(oval Rect, paint *Paint) {
	toimpl()
}

func (canvas *Canvas) init(device *BaseDevice, flags CanvasInitFlags) *BaseDevice {
	if device != nil && device.forceConservativeRasterClip() {
		flags = flags | KCanvasInitFlagConservativeRasterClip
	}

	canvas.conservativeRasterClip = (flags&KCanvasInitFlagConservativeRasterClip != 0)
	canvas.allowSoftClip = true
	canvas.allowSimplifyClip = false
	canvas.deviceCMDirty = true
	canvas.saveCount = 1
	canvas.metaData = nil
	canvas.clipStack.Reset(NewClipStack())
	canvas.mcRec = newCanvasMCRec(canvas.conservativeRasterClip)
	canvas.mcRec.Layer = newDeivceCM(nil, nil, nil, canvas.conservativeRasterClip, canvas.mcRec.Matrix)
	canvas.mcStack = list.New()
	canvas.mcStack.PushBack(canvas.mcRec)
	canvas.isScaleTranslate = true
	canvas.mcRec.TopLayer = canvas.mcRec.Layer
	canvas.baseSurface = nil

	if device != nil {
		canvas.mcRec.Layer.Device = device
		canvas.mcRec.RasterClip.SetRect(device.GlobalBounds())
		canvas.deviceClipBounds = quickRejectClipBounds(device.GlobalBounds())
	}

	return device
}

/** Return the current matrix on the canvas.
	This does not account for the translate in any of the devices.
	@return The current matrix on the canvas.
*/
func (canvas *Canvas) TotalMatrix() *Matrix {
	return canvas.mcRec.Matrix
}

func (canvas *Canvas) OnDrawPoints(mode CanvasPointMode, count int, pts []Point, paint *Paint) {
	toimpl()
}

type LazyPaint Lazy

func (canvas *Canvas) OnDrawPaint(paint *Paint) {
	canvas.internalDrawPaint(paint)
}

func (canvas *Canvas) internalDrawPaint(paint *Paint) {
	canvas.PredrawNotify(nil, paint, KCanvasShaderOverrideOpacityNotOpaque)

	var looper = newAutoDrawLooper(canvas, paint, false, nil)
	for looper.Next(KDrawFilterTypePaint) {
		var it = NewDrawIterator(canvas)
		for it.Next() {
			it.Device().DrawPaint(it.Draw, looper.Paint())
		}
	}
}

func (canvas *Canvas) PredrawNotify(rect *Rect, paint *Paint, overrideOpacity CanvasShaderOverrideOpacity) {
	toimpl()
}

func (canvas *Canvas) internalSaveLayer(rec *tCanvasSaveLayerRec, strategy CanvasSaveLayerStrategy) {
	toimpl()
}

func (canvas *Canvas) internalRestore() {
	toimpl()
}

// Returns the canvas to be used by DrawIterator. Default implementation
// returns this. Subclasses that encapsulate an indirect canvas may
// need to overload this method. The impl must keep track of this, as it
// is not released or deleted by the caller.
func (canvas *Canvas) CanvasForDrawIterator() *Canvas {
	return canvas
}

func (canvas *Canvas) UpdateDeviceCMCache() {
	if canvas.deviceCMDirty {
		var totalMatrix = canvas.TotalMatrix()
		var totalClip = canvas.mcRec.RasterClip
		var layer = canvas.mcRec.TopLayer

		if layer.Next == nil { // < only one layer.
			layer.UpdateMC(totalMatrix, totalClip, canvas.clipStack, nil)
		} else {
			var clip = NewRasterClipClone(totalClip)
			for layer.Next != nil {
				layer.UpdateMC(totalMatrix, clip, canvas.clipStack, clip)
				layer = layer.Next
			}
		}
	}
}

type CanvasSaveFlags int

const (
	KSaveFlagHasAlphaLayer   = 0x01
	KSaveFlagFullColorLayer  = 0x02
	KSaveFlagClipToLayer     = 0x10
	KSaveFlagARGBNoClipLayer = 0x0F
	KSaveFlagARGBClipLayer   = 0x1F
)

/**
tDeviceCM is the record we keep for each BaseDevice that the user installs.
The clip/matrix/proc are fields that reflect the top of the save/resotre
stack. Whenever the canvas changes, it makes a dirty flag, and then before
these are used (assuming we're not on a layer) we rebuild these cache values:
they reflect the top of the save stack, but translated and clipped by the
device's XY offset and bitmap-bounds.
*/
type tDeviceCM struct {
	Next          *tDeviceCM
	Device        *BaseDevice
	Clip          *RasterClip
	Paint         *Paint
	Matrix        *Matrix
	MatrixStroage *Matrix
	StashedMatrix *Matrix
}

func newDeivceCM(dev *BaseDevice, paint *Paint, canvas *Canvas, conservativeRasterClip bool, stashed *Matrix) *tDeviceCM {
	var deviceCM = new(tDeviceCM)
	deviceCM.Next = nil
	deviceCM.Clip = NewRasterClip(conservativeRasterClip)
	deviceCM.StashedMatrix = stashed
	deviceCM.Device = dev
	if paint != nil {
		deviceCM.Paint = paint
	} else {
		deviceCM.Paint = nil
	}
	return deviceCM
}

func (d *tDeviceCM) Reset(bounds Rect) {
	d.Clip.SetRect(bounds)
}

func (d *tDeviceCM) UpdateMC(totalMatrix *Matrix, totlaClip *RasterClip,
	clipStack *ClipStack, updateClip *RasterClip) {
	toimpl()
}

/**
tCanvasMCRec is the record we keep for each save/restore level in the stack.
Since a level optionally copies the matrix and/or stack, we have pointers
for these fields. If the value is copied for this level, the copy is stored
in the ...Storage field, and the pointer points to that. If the value is not
copied for this level, we ignore ...Storage, and just point at the
corresponding value in the previous level in the stack.
*/
type tCanvasMCRec struct {
	Filter *DrawFilter // the current filter (or nil)
	Layer  *tDeviceCM

	/**
	If there are any layers in the stack, this points to the top-most
	one that is at or below this level in the stack (so we know what
	bitmap/device to draw into from this level. This value is NOT
	reference counted, since the real owner is either our fLayer field,
	or a previous one in a lower level.)
	*/
	TopLayer          *tDeviceCM
	RasterClip        *RasterClip
	Matrix            *Matrix
	DeferredSaveCount int
	CurDrawDepth      int
}

func newCanvasMCRec(conservativeRasterClip bool) *tCanvasMCRec {
	var rec = &tCanvasMCRec{
		RasterClip:        NewRasterClip(conservativeRasterClip),
		Filter:            nil,
		Layer:             nil,
		TopLayer:          nil,
		Matrix:            NewMatrix(),
		DeferredSaveCount: 0,
		CurDrawDepth:      0,
	}

	// don't bother initializing fNext
	// todo: inc_rec()

	return rec
}

/**
Subclass save/restore notifiers.
Overriders should call the corresponding INHERITED method up the inheritance chain.
getSaveLayerStrategy()'s return value may suppress full layer allocation.
*/
type CanvasSaveLayerStrategy int

const (
	KCanvasSaveLayerStrategyFullLayer = iota
	KCanvasSaveLayerStrategyNoLayer
)

type tAutoDrawLooper struct {
	lazyPaintInit           *Lazy
	lazyPaintPerLooper      *Lazy
	canvas                  *Canvas
	origPaint               *Paint
	paint                   *Paint
	filter                  DrawFilter
	saveCount               int
	tempLayerForImageFilter bool
	done                    bool
	isSimple                bool
	looperContext           *tDrawLooperContext
}

func newAutoDrawLooper(canvas *Canvas, paint *Paint, skipLayerForImageFilter bool, rawBounds *Rect) *tAutoDrawLooper {
	var looper = &tAutoDrawLooper{
		canvas:                  canvas,
		origPaint:               paint,
		paint:                   paint,
		filter:                  nil,
		saveCount:               canvas.SaveCount(),
		tempLayerForImageFilter: false,
		done: false,
	}

	var simplifiedCF = imageToColorFilter(looper.origPaint)
	if simplifiedCF != nil {
		var paint = setIfNeeded(looper.lazyPaintInit, looper.origPaint)
		paint.SetColorFilter(simplifiedCF)
		paint.SetImageFilter(nil)
		looper.paint = paint
	}

	if !skipLayerForImageFilter && looper.paint.ImageFilter() != nil {
		/* We implement ImageFilters for a given draw by creating a layer, then applying the
		imagefilter to the pixels of that layer (its backing surface/image), and then
		we call restore() to xfer that layer to the main canvas.

		1. SaveLayer (with a paint containing the current imagefilter and xfermode)
		2. Generate the src pixels:
		    Remove the imagefilter and the xfermode from the paint that we (AutoDrawLooper)
		    return (fPaint). We then draw the primitive (using srcover) into a cleared
		    buffer/surface.
		3. Restore the layer created in #1
		    The imagefilter is passed the buffer/surface from the layer (now filled with the
		    src pixels of the primitive). It returns a new "filtered" buffer, which we
		    draw onto the previous layer using the xfermode from the original paint. */
		var tmp = NewPaint()
		tmp.SetImageFilter(looper.paint.ImageFilter())
		tmp.SetXfermode(looper.paint.Xfermode())
		var storage Rect
		if rawBounds != nil {
			// Make rawBounds include all paint outsets except for those due to image filters.
			*rawBounds = applyPaintToBoundsSansImageFilter(looper.paint, *rawBounds, &storage)
		}
		canvas.internalSaveLayer(newCanvasSaveLayerRec(rawBounds, tmp, nil, KCanvasSaveLayerFlagIsOpaque),
			KCanvasSaveLayerStrategyFullLayer)
		looper.tempLayerForImageFilter = false
		// we remove the imagefilter/xfermode inside doNext()
	}

	if paint.Looper() != nil {
		// looper.looperContext = paint.Looper().CreateContext(canvas)
		looper.isSimple = false
	} else {
		looper.looperContext = nil
		// can we be marked as simple?
		looper.isSimple = looper.filter != nil && !looper.tempLayerForImageFilter
	}

	return looper
}

func (looper *tAutoDrawLooper) Paint() *Paint {
	return looper.paint
}

func (looper *tAutoDrawLooper) Next(drawType DrawFilterType) bool {
	if looper.done {
		return false
	} else if looper.isSimple {
		looper.done = false
		return !looper.paint.NothingToDraw()
	} else {
		return looper.doNext(drawType)
	}
}

func (looper *tAutoDrawLooper) doNext(drawType DrawFilterType) bool {
	looper.paint = nil

	var origPaint = looper.origPaint
	if looper.lazyPaintInit.IsValid() {
		origPaint, _ = looper.lazyPaintInit.Get().(*Paint)
	}

	var paint, _ = looper.lazyPaintPerLooper.Set(origPaint).(*Paint)

	if looper.tempLayerForImageFilter {
		paint.SetImageFilter(nil)
		paint.SetXfermode(nil)
	}

	if looper.looperContext != nil && looper.looperContext.Next(looper.canvas, paint) {
		looper.done = true
		return false
	}

	if looper.filter != nil {
		if looper.filter.Filter(paint, drawType) == false {
			looper.done = true
			return false
		}
		if looper.looperContext == nil {
			// no looper means we only draw once.
			looper.done = true
		}
	}
	looper.paint = paint

	// if we only came in here for the imagefilter, mark us as done.
	if looper.looperContext == nil && looper.filter == nil {
		looper.done = true
	}

	// call this after any possible paint modifiers.
	if looper.paint.NothingToDraw() {
		looper.paint = nil
		return false
	}

	return true
}

func (looper *tAutoDrawLooper) Finalizer() {
	if looper.tempLayerForImageFilter {
		looper.canvas.internalRestore()
	}
}

type DrawIterator struct {
	*Draw
	canvas    *Canvas
	currLayer *tDeviceCM
	paint     *Paint // May be nil
}

func NewDrawIterator(canvas *Canvas) *DrawIterator {
	canvas = canvas.CanvasForDrawIterator()
	canvas.UpdateDeviceCMCache()
	var it = &DrawIterator{
		canvas:    canvas,
		currLayer: canvas.mcRec.TopLayer,
		Draw: &Draw{
			clipStack: canvas.clipStack,
		},
	}
	return it
}

func (it *DrawIterator) Next() bool {
	for it.currLayer != nil && it.currLayer.Clip.IsEmpty() {
		it.currLayer = it.currLayer.Next
	}

	var rec *tDeviceCM = it.currLayer
	if rec != nil && rec.Device != nil {
		it.matrix = rec.Matrix
		it.rasterClip = rec.Clip
		it.device = rec.Device
		if !it.device.AccessPixels(it.dst) {
			it.dst.Reset(it.device.imageInfo, nil, 0, nil)
		}
		it.paint = rec.Paint
		it.currLayer = rec.Next

		return true
	}
	return false
}

/**
If the paint has an imagefilter, but it can be simplified to just a colorfilter, return that
colorfilter, else return nullptr.
*/
func imageToColorFilter(paint *Paint) *ColorFilter {
	var imageFilter = paint.ImageFilter()
	if imageFilter == nil {
		return nil
	}

	var imageColorFilter, ok = imageFilter.AsAColorFilter()
	if !ok {
		return nil
	}

	var paintColorFilter *ColorFilter
	if paintColorFilter = paint.ColorFilter(); paintColorFilter == nil {
		// there is no existing paint colorfilter, so we can just return the imagefilter's
		return imageColorFilter
	}

	// The paint has both a colorfilter(paintCF) and an imagefilter-which-is-a-colorfilter(imgCF)
	// and we need to combine them into a single colorfilter.
	return NewColorFilterFromComposeFilter(imageColorFilter, paintColorFilter)
}

func setIfNeeded(lazyPaint *Lazy, paint *Paint) *Paint {
	toimpl()
	return nil
}

func applyPaintToBoundsSansImageFilter(paint *Paint, rowBounds Rect, storage *Rect) Rect {
	toimpl()
	return RectZero
}

func quickRejectClipBounds(bounds Rect) Rect {
	toimpl()
	return RectZero
}
