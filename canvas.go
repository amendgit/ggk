package ggk

import (
	"container/list"
)

type CanvasImpl interface {
	/* BaseLayerSize
	Gets the size of the base or root layer in global canvas coordinates. The
	origin of the base layer is always (0,0). The current drawable area may be
	smaller (due to clipping or saveLayer).	*/
	BaseLayerSize() Size

	/** ClipBounds
	Return the bounds of the current clip (in local coordinates) in the
	bounds parameter, and return true if it is non-empty. This can be useful
	in a way similar to quickReject, in that it tells you that drawing
	outside of these bounds will be clipped out. */
	ClipBounds(bounds *Rect) bool

	/** ClipDeviceBounds
	Return the bounds of the current clip, in device coordinates; returns
	true if non-empty. Maybe faster than getting the clip explicitly and
	then taking its bounds. */
	ClipDeviceBounds(bounds *Rect) bool

	/** SetDrawFilter
	Set the new filter (or NULL). Pass NULL to clear any existing filter.
	As a convenience, the parameter is returned. If an existing filter
	exists, its refcnt is decrement. If the new filter is not null, its
	refcnt is incremented. The filter is saved/restored, just like the
	matrix and clip.
	@param filter the new filter (or NULL)
	@return the new filter */
	SetDrawFilter(filter *DrawFilter)

	/** IsClipEmpty
	Return true if the current clip is empty (i.e. nothing will draw).
	Note: this is not always a free call, so it should not be used
	more often than necessary. However, once the canvas has computed this
	result, subsequent calls will be cheap (until the clip state changes,
	which can happen on any clip..() or restore() call. */
	IsClipEmpty() bool

	OnNewSurface(imageInfo *ImageInfo, surfaceProps *SurfaceProps)
	OnPeekPixels(pixmap *Pixmap) bool
	OnAccessTopLayerPixles(pixmap *Pixmap) bool
	OnImageInfo() *ImageInfo
	OnGetProps() (*SurfaceProps, bool)

	WillSave()
	SaveLayerStrategy() CanvasSaveLayerStrategy
	WillRestore()
	DidRestore()
	DidConcat(matrix *Matrix)
	DidSetMatrix(matrix *Matrix)
	DidTranslate(dx, dy Scalar)
	DidTranslateZ(z Scalar)

	OnDrawAnnotation(rect Rect, kay []byte, value *Data)
	OnDrawDRect(outter, inner Rect, paint *Paint)
	OnDrawText(text string, x, y Scalar, paint *Paint)
	OnDrawTextAt(text string, xpos []Point, constY Scalar, paint *Paint)
	OnDrawTextAtH(text string, xpos []Point, constY Scalar, paint *Paint)
	OnDrawTextOnPath(text string, path *Path, matrix *Matrix, paint *Paint)
	OnDrawTextRSXform(text string, xform []RSXform, cullRect *Rect, paint *Paint)
	OnDrawTextBlob(blob *TextBlob, x, y Scalar, paint *Paint)
	OnDrawPatch(cubics [12]Point, colors [4]Color, texCoords [4]Point, xmode *Xfermode, paint *Paint)
	OnDrawDrawable(drawable *Drawable, matrixe *Matrix)
	OnDrawPaint(paint *Paint)
	OnDrawRect(rect Rect, paint *Paint)
	OnDrawOval(oval Rect, paint *Paint)
	OnDrawArc(oval Rect, startAngle, sweepAngle Scalar, useCenter bool, paint *Paint)
	OnDrawPoints(mode CanvasPointMode, count int, pts []Point, paint *Paint)
	OnDrawVertices(vertexMode CanvasVertexMode, vertexCount int, vertices []Point, texs []Point,
		colors []Color, xfermode *Xfermode, indices []uint16, indexCount int, paint *Paint)
	OnDrawAtlas(atlas *Image, xform []RSXform, tex []Rect, colors []Color, count int,
		mode XfermodeMode, cull *Rect, paint *Paint)
	OnDrawPath(path *Path, paint *Paint)
	OnDrawImage(image *Image, dx, dy Scalar, paint *Paint)
	OnDrawImageRect(image *Image, src *Rect, dst Rect, paint *Paint,
		constraint CanvasSrcRectConstraint)
	OnDrawImageNine(image *Image, center Rect, dst Rect, paint *Paint)
	OnDrawImageLattice(image *Image, lattice *CanvasLattice, dst Rect, paint *Paint)
	OnDrawBitmap(bmp *Bitmap, dx, dy Scalar, paint *Paint)
	OnDrawBitmapRect(bmp *Bitmap, src *Rect, dst Rect, paint *Paint,
		constraint CanvasSrcRectConstraint)
	OnDrawBitmapNine(bmp *Bitmap, center Rect, dst Rect, paint *Paint)

	OnClipRect(rect Rect, op RegionOp, edgeStyle ClipEdgeStyle)
	OnClipPath(path *Path, op RegionOp, edgeStyle ClipEdgeStyle)
	OnClipRegion(deviceRgn *Region, op RegionOp)
	OnDiscard()
	OnDrawPicture(pic *Picture, matrix *Matrix, paint *Paint)
	OnDrawShadowedPicture(pic *Picture, matrix *Matrix, paint *Paint)

	/** CanvasForDrawIterator
	Returns the canvas to be used by DrawIter. Default implementation
	returns this. Subclasses that encapsulate an indirect canvas may
	need to overload this method. The impl must keep track of this, as it
	is not released or deleted by the caller. */
	CanvasForDrawIterator() *Canvas
}

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
etc. */
type Canvas struct {
	Impl CanvasImpl

	surfaceProps *SurfaceProps
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
dimensions. */
func NewCanvasRasterDirect(imageInfo *ImageInfo, pixels []byte, rowBytes int) *Canvas {
	toimpl()
	return &Canvas{}
}

func NewCanvasRasterDirectN32(width, height int, pixels []PremulColor, rowBytes int) *Canvas {
	toimpl()
	return &Canvas{}
}

/** NewCanvasEmpty
Creates an empty canvas with no backing device/pixels, and zero
dimensions. */
func NewCanvasEmpty() *Canvas {
	toimpl()
	return &Canvas{}
}

/**
Construct a canvas with the specified bitmap to draw into.
@param bitmap   Specifies a bitmap for the canvas to draw into. Its
                structure are copied to the canvas. */
func NewCanvasBitmap(bmp *Bitmap) *Canvas {
	var canvas = new(Canvas)
	canvas.Impl = canvas
	canvas.surfaceProps = NewSurfaceProps(KSurfacePropsFlagNone, KSurfacePropsInitTypeLegacyFontHost)
	canvas.mcStack = list.New()

	var device = NewBitmapDevice(bmp, canvas.surfaceProps)
	canvas.init(device.BaseDevice, KCanvasInitFlagDefault)

	return canvas
}

/**
Construct a canvas with the specified bitmap to draw into.
@param bitmap   Specifies a bitmap for the canvas to draw into. Its
				structure are copied to the canvas.
@param props    New canvas surface properties. */
func NewCanvasBitmapSurfaceProps(bmp *Bitmap, surfaceProps *SurfaceProps) *Canvas {
	toimpl()
	return &Canvas{}
}

func (canvas *Canvas) MetaData() *MetaData {
	toimpl()
	return nil
}

/**
Return ImageInfo for this canvas. If the canvas is not backed by pixels
(cpu or gpu), then the info's ColorType will be kUnknown_SkColorType. */
func (canvas *Canvas) ImageInfo() *ImageInfo {
	toimpl()
	return nil
}

/**
If the canvas is backed by pixels (cpu or gpu), this writes a copy of the SurfaceProps
for the canvas to the location supplied by the caller, and returns true. Otherwise,
return false and leave the supplied props unchanged. */
func (canvas *Canvas) SurfaceProps() *SurfaceProps {
	toimpl()
	return nil
}

/**
Trigger the immediate execution of all pending draw operations. For the GPU
backend this will resolve all rendering to the GPU surface backing the
SkSurface that owns this canvas. */
func (canvas *Canvas) Flush() {
	toimpl()
}

/** impl CanvasImpl */
func (canvas *Canvas) BaseLayerSize() Size {
	var device = canvas.Device()
	var size Size
	if device != nil {
		size = MakeSize(device.Width(), device.Height())
	}
	return size
}

/**
Return the canvas' device object, which may be null. The device holds
the bitmap of the pixels that the canvas draws into. The reference count
of the returned device is not changed by this call. */
func (canvas *Canvas) Device() *BaseDevice {
	var rec = canvas.mcStack.Front().Value.(*tCanvasMCRec)
	return rec.Layer.Device
}

/**
saveLayer() can create another device (which is later drawn onto
the previous device). getTopDevice() returns the top-most device current
installed. Note that this can change on other calls like save/restore,
so do not access this device after subsequent canvas calls.
The reference count of the device is not changed.

@param updateMatrixClip If this is true, then before the device is
      returned, we ensure that its has been notified about the current
      matrix and clip. Note: this happens automatically when the device
      is drawn to, but is optional here, as there is a small perf hit
      sometimes. */
func (canvas *Canvas) TopDevice() *BaseDevice {
	toimpl()
	return nil
}

/**
Create a new surface matching the specified info, one that attempts to
be maximally compatible when used with this canvas. If there is no matching Surface type,
NULL is returned.

If surfaceprops is specified, those are passed to the new surface, otherwise the new surface
inherits the properties of the surface that owns this canvas. If this canvas has no parent
surface, then the new surface is created with default properties. */
func (canvas *Canvas) NewSurface(imageInfo *ImageInfo, surfaceProps *SurfaceProps) *Surface {
	toimpl()
	return nil
}

/**
Return the GPU context of the device that is associated with the canvas.
For a canvas with non-GPU device, NULL is returned. */
func (canvas *Canvas) GrContext() *GrContext {
	toimpl()
	return nil
}

/**
If the canvas has writable pixels in its top layer (and is not recording to a picture
or other non-raster target) and has direct access to its pixels (i.e. they are in
local RAM) return the address of those pixels, and if not null,
return the ImageInfo, rowBytes and origin. The returned address is only valid
while the canvas object is in scope and unchanged. Any API calls made on
canvas (or its parent surface if any) will invalidate the
returned address (and associated information).

On failure, returns NULL and the info, rowBytes, and origin parameters are ignored. */
func (canvas *Canvas) AccessTopLayerPixels(imageInfo *ImageInfo, rowBytes *int, origin *Point) []byte {
	toimpl()
	return nil
}

/**
If the canvas has readable pixels in its base layer (and is not recording to a picture
or other non-raster target) and has direct access to its pixels (i.e. they are in
local RAM) return true, and if not null, return in the pixmap parameter information about
the pixels. The pixmap's pixel address is only valid
while the canvas object is in scope and unchanged. Any API calls made on
canvas (or its parent surface if any) will invalidate the pixel address
(and associated information).

On failure, returns false and the pixmap parameter will be ignored. */
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
- If this canvas is not backed by pixels (e.g. picture or PDF) */
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
Helper for calling readPixels(info, ...). This call will check if bitmap has been allocated.
If not, it will attempt to call allocPixels(). If this fails, it will return false. If not,
it calls through to readPixels(info, ...) and returns its result. */
func (canvas *Canvas) ReadPixelsToBitmap(bmp *Bitmap, srcX, srcY int) bool {
	toimpl()
	return false
}

/**
Helper for allocating pixels and then calling readPixels(info, ...). The bitmap is resized
to the intersection of srcRect and the base-layer bounds. On success, pixels will be
allocated in bitmap and true returned. On failure, false is returned and bitmap will be
set to empty. */
func (canvas *Canvas) ReadPixelsInRectToBitmap(rect Rect, bmp *Bitmap) bool {
	toimpl()
	return false
}

/**
This method affects the pixels in the base-layer, and operates in pixel coordinates,
ignoring the matrix and clip.

The specified ImageInfo and (x,y) offset specifies a rectangle: target.

    target.setXYWH(x, y, info.width(), info.height());

Target is intersected with the bounds of the base-layer. If this intersection is not empty,
then we have two sets of pixels (of equal size), the "src" specified by info+pixels+rowBytes
and the "dst" by the canvas' backend. Replace the dst pixels with the corresponding src
pixels, performing any colortype/alphatype transformations needed (in the case where the
src and dst have different colortypes or alphatypes).

This call can fail, returning false, for several reasons:
- If the src colortype/alphatype cannot be converted to the canvas' types
- If this canvas is not backed by pixels (e.g. picture or PDF) */
func (canvas *Canvas) WritePixels(info *ImageInfo, pixels []byte, rowBytes int, x, y int) bool {
	toimpl()
	return false
}

/**
Helper for calling writePixels(info, ...) by passing its pixels and rowbytes. If the bitmap
is just wrapping a texture, returns false and does nothing. */
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

@return The value to pass to restoreToCount() to balance this save() */
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
@return The value to pass to restoreToCount() to balance this save() */
func (canvas *Canvas) SaveLayer(bounds *Rect, paint *Paint) int {
	toimpl()
	return 0
}

/**
Temporary name.
Will allow any requests for LCD text to be respected, so the caller must be careful to
only draw on top of opaque sections of the layer to get good results. */
func SaveLayerPreserveLCDTextRequests(bounds *Rect, paint *Paint) int {
	toimpl()
	return 0
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
@param alpha  This is applied to the offscreen when restore() is called.
@return The value to pass to restoreToCount() to balance this save() */
func (canvas *Canvas) SaveLayerAlpha(bounds *Rect, alpha uint8) int {
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
It is an error to call restore() more times than save() was called. */
func (canvas *Canvas) Restore() {
	toimpl()
}

/**
Returns the number of matrix/clip states on the SkCanvas' private stack.
This will equal # save() calls - # restore() calls + 1. The save count on
a new canvas is 1. */
func (canvas *Canvas) SaveCount() int {
	return canvas.saveCount
}

/**
Efficient way to pop any calls to save() that happened after the save
count reached saveCount. It is an error for saveCount to be greater than
getSaveCount(). To pop all the way back to the initial matrix/clip context
pass saveCount == 1.
@param saveCount    The number of save() levels to restore from */
func (canvas *Canvas) RestoreToCount(saveCount int) {
	toimpl()
	return
}

/**
Preconcat the current matrix with the specified translation
@param dx   The distance to translate in X
@param dy   The distance to translate in Y */
func (canvas *Canvas) Translate(dx, dy Scalar) {
	toimpl()
}

/**
Preconcat the current matrix with the specified scale.
@param sx   The amount to scale in X
@param sy   The amount to scale in Y */
func (canvas *Canvas) Scale(sx, sy Scalar) {
	toimpl()
}

/**
Preconcat the current matrix with the specified rotation about the origin.
@param degrees  The amount to rotate, in degrees */
func (canvas *Canvas) Rotate(degrees Scalar) {
	toimpl()
}

/**
Preconcat the current matrix with the specified rotation about a given point.
@param degrees  The amount to rotate, in degrees
@param px  The x coordinate of the point to rotate about.
@param py  The y coordinate of the point to rotate about. */
func (canvas *Canvas) RotateAt(degrees, px, py Scalar) {
	toimpl()
}

/**
Preconcat the current matrix with the specified skew.
@param sx   The amount to skew in X
@param sy   The amount to skew in Y */
func (canvas *Canvas) Skew(sx, sy Scalar) {
	toimpl()
}

/**
Preconcat the current matrix with the specified matrix.
@param matrix   The matrix to preconcatenate with the current matrix */
func (canvas *Canvas) Concat(matrix *Matrix) {
	toimpl()
}

/**
Replace the current matrix with a copy of the specified matrix.
@param matrix The matrix that will be copied into the current matrix. */
func (canvas *Canvas) SetMatrix(matrix *Matrix) {
	toimpl()
}

/**
Helper for setMatrix(identity). Sets the current matrix to identity. */
func (canvas *Canvas) ResetMatrix() {
	toimpl()
}

/**
Add the specified translation to the current draw depth of the canvas.
@param z    The distance to translate in Z.
			Negative into screen, positive out of screen.
			Without translation, the draw depth defaults to 0. */
func (canvas *Canvas) TranslateZ(z Scalar) {
	toimpl()
}

/**
Set the current set of lights in the canvas.
@param lights   The lights that we want the canvas to have. */
func (canvas *Canvas) SetLights(lights *Lights) {
	toimpl()
}

/** Returns the current set of lights the canvas uses
 */
func (canvas *Canvas) Lights() *Lights {
	return nil
}

/**
Modify the current clip with the specified rectangle.
@param rect The rect to combine with the current clip
@param op The region op to apply to the current clip
@param doAntiAlias true if the clip should be antialiased */
func (canvas *Canvas) ClipRect(rect Rect, op RegionOp, doAntiAlias bool) {
	toimpl()
}

/**
Modify the current clip with the specified path.
@param path The path to combine with the current clip
@param op The region op to apply to the current clip
@param doAntiAlias true if the clip should be antialiased */
func (canvas *Canvas) ClipPath(path *Path, op RegionOp, doAntiAlias bool) {
	toimpl()
}

/**
EXPERIMENTAL -- only used for testing
Set to false to force clips to be hard, even if doAntiAlias=true is
passed to clipRect or clipPath. */
func (canvas *Canvas) SetAllowSoftClip(allow bool) {
	toimpl()
}

/**
EXPERIMENTAL -- only used for testing
Set to simplify clip stack using path ops. */
func (canvas *Canvas) SetAllowSimplifyClip(allow bool) {
	toimpl()
}

/**
Modify the current clip with the specified region. Note that unlike
clipRect() and clipPath() which transform their arguments by the current
matrix, clipRegion() assumes its argument is already in device
coordinates, and so no transformation is performed.
@param deviceRgn    The region to apply to the current clip
@param op The region op to apply to the current clip */
func (canvas *Canvas) ClipRegion(deviceRgn *Region, op RegionOp) {
	toimpl()
}

/** Helper for clipRegion(rgn, kReplace_Op). Sets the current clip to the
specified region. This does not intersect or in any other way account
for the existing clip region.
@param deviceRgn The region to copy into the current clip. */
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
			 intersect with the canvas' clip */
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
			 intersect with the canvas' clip */
func (canvas *Canvas) QuickRejectPath(path *Path) bool {
	toimpl()
	return false
}

/** ClipBounds
Return the bounds of the current clip (in local coordinates) in the
bounds parameter, and return true if it is non-empty. This can be useful
in a way similar to quickReject, in that it tells you that drawing
outside of these bounds will be clipped out.
Impl CanvasImpl */
func (canvas *Canvas) ClipBounds(bounds *Rect) bool {
	toimpl()
	return false
}

/** ClipDeviceBounds
Return the bounds of the current clip, in device coordinates; returns
true if non-empty. Maybe faster than getting the clip explicitly and
then taking its bounds.
Impl CanvasImpl */
func (canvas *Canvas) ClipDeviceBounds(bounds *Rect) bool {
	toimpl()
	return false
}

/** DrawARGB
Fill the entire canvas' bitmap (restricted to the current clip) with the
specified ARGB color, using the specified mode.
@param a    the alpha component (0..255) of the color to fill the canvas
@param r    the red component (0..255) of the color to fill the canvas
@param g    the green component (0..255) of the color to fill the canvas
@param b    the blue component (0..255) of the color to fill the canvas
@param mode the mode to apply the color in (defaults to SrcOver) */
func (canvas *Canvas) DrawARGB(a, r, g, b uint8, mode XfermodeMode) {
	toimpl()
}

/** DrawColor
Fill the entire canvas' bitmap (restricted to the current clip) with the
specified color and mode.
@param color    the color to draw with
@param mode the mode to apply the color in (defaults to SrcOver) */
func (canvas *Canvas) DrawColor(color Color, mode XfermodeMode) {
	var paint = NewPaint()
	paint.SetColor(color)
	if KXfermodeModeSrcOver == mode {
		paint.SetXfermodeMode(mode)
	}
	canvas.DrawPaint(paint)
}

/** Clear
Helper method for drawing a color in SRC mode, completely replacing all the pixels
in the current clip with this color. */
func (canvas *Canvas) Clear(color Color) {
	toimpl()
}

/** Discard
This makes the contents of the canvas undefined. Subsequent calls that
require reading the canvas contents will produce undefined results. Examples
include blending and readPixels. The actual implementation is backend-
dependent and one legal implementation is to do nothing. This method
ignores the current clip.

This function should only be called if the caller intends to subsequently
draw to the canvas. The canvas may do real work at discard() time in order
to optimize performance on subsequent draws. Thus, if you call this and then
never draw to the canvas subsequently you may pay a perfomance penalty. */
func (canvas *Canvas) Discard() {
	toimpl()
}

/** DrawPaint
Fill the entire canvas (restricted to the current clip) with the
specified paint.
@param paint    The paint used to fill the canvas */
func (canvas *Canvas) DrawPaint(paint *Paint) {
	canvas.Impl.OnDrawPaint(paint)
}

type CanvasPointMode int

const (
	KCanvasPointModePoints  CanvasPointMode = iota // DrawPoints draws each point separately
	KCanvasPointModeLines                          // DrawPoints draws each pair of points as a line segment
	KCanvasPointModePolygon                        // DrawPoints draws the array of points as a polygon
)

/** DrawPoints
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
@param paint    The paint used to draw the points */
func (canvas *Canvas) DrawPoints(mode CanvasPointMode, count int, pts []Point, paint *Paint) {
	canvas.OnDrawPoints(mode, count, pts, paint)
}

/** DrawPoint
Draws a single pixel in the specified color.
@param x        The X coordinate of which pixel to draw
@param y        The Y coordiante of which pixel to draw
@param color    The color to draw */
func (canvas *Canvas) DrawPoint(x, y Scalar, paint *Paint) {
	var pt Point
	pt.X, pt.Y = x, y
	canvas.DrawPoints(KCanvasPointModePoints, 1, []Point{pt}, paint)
}

/** DrawLine
Draw a line segment with the specified start and stop x,y coordinates,
using the specified paint. NOTE: since a line is always "framed", the
paint's Style is ignored.
@param x0    The x-coordinate of the start point of the line
@param y0    The y-coordinate of the start point of the line
@param x1    The x-coordinate of the end point of the line
@param y1    The y-coordinate of the end point of the line
@param paint The paint used to draw the line */
func (canvas *Canvas) DrawLine(x0, y0, x1, y1 Scalar, paint *Paint) {
	toimpl()
}

/** DrawRect
Draw the specified rectangle using the specified paint. The rectangle
will be filled or stroked based on the Style in the paint.
@param rect     The rect to be drawn
@param paint    The paint used to draw the rect */
func (canvas *Canvas) DrawRect(rect Rect, paint *Paint) {
	toimpl()
}

/** DrawRectCoords
Draw the specified rectangle using the specified paint. The rectangle
will be filled or framed based on the Style in the paint.
@param left     The left side of the rectangle to be drawn
@param top      The top side of the rectangle to be drawn
@param right    The right side of the rectangle to be drawn
@param bottom   The bottom side of the rectangle to be drawn
@param paint    The paint used to draw the rect */
func (canvas *Canvas) DrawRectCoords(left, top, right, bottom Scalar, paint *Paint) {
	toimpl()
}

/** DrawOval
Draw the specified oval using the specified paint. The oval will be
filled or framed based on the Style in the paint.
@param oval     The rectangle bounds of the oval to be drawn
@param paint    The paint used to draw the oval */
func (canvas *Canvas) DrawOval(oval Rect, paint *Paint) {
	toimpl()
}

/** DrawDRect
Draw the annulus formed by the outer and inner rrects. The results
are undefined if the outer does not contain the inner. */
func (canvas *Canvas) DrawDRect(outer, inner Rect, paint *Paint) {
	toimpl()
}

/** DrawCircle
Draw the specified circle using the specified paint. If radius is <= 0,
then nothing will be drawn. The circle will be filled
or framed based on the Style in the paint.
@param cx       The x-coordinate of the center of the cirle to be drawn
@param cy       The y-coordinate of the center of the cirle to be drawn
@param radius   The radius of the cirle to be drawn
@param paint    The paint used to draw the circle */
func (canvas *Canvas) DrawCircle(cx, cy, radius Scalar, paint *Paint) {
	toimpl()
}

/** DrawArc
Draw the specified arc, which will be scaled to fit inside the
specified oval. If the sweep angle is >= 360, then the oval is drawn
completely. Note that this differs slightly from SkPath::arcTo, which
treats the sweep angle mod 360.
@param oval The bounds of oval used to define the shape of the arc.
@param startAngle Starting angle (in degrees) where the arc begins
@param sweepAngle Sweep angle (in degrees) measured clockwise.
@param useCenter true means include the center of the oval. For filling
				 this will draw a wedge. False means just use the arc.
@param paint    The paint used to draw the arc */
func (canvas *Canvas) DrawArc(oval Rect, startAngle, sweepAngle Scalar, useCenter bool, paint *Paint) {
	toimpl()
}

/** DrawRoundRect
Draw the specified round-rect using the specified paint. The round-rect
will be filled or framed based on the Style in the paint.
@param rect     The rectangular bounds of the roundRect to be drawn
@param rx       The x-radius of the oval used to round the corners
@param ry       The y-radius of the oval used to round the corners
@param paint    The paint used to draw the roundRect */
func (canvas *Canvas) DrawRoundRect(rect Rect, rx, ry Scalar, paint *Paint) {
	toimpl()
}

/** DrawPath
Draw the specified path using the specified paint. The path will be
filled or framed based on the Style in the paint.
@param path     The path to be drawn
@param paint    The paint used to draw the path */
func (canvas *Canvas) DrawPath(path *Path, paint *Paint) {
	toimpl()
}

/** DrawImage
Draw the specified image, with its top/left corner at (x,y), using the
specified paint, transformed by the current matrix.

@param image    The image to be drawn
@param left     The position of the left side of the image being drawn
@param top      The position of the top side of the image being drawn
@param paint    The paint used to draw the image, or NULL */
func (canvas *Canvas) DrawImage(image *Image, left, top Scalar, paint *Paint) {
	toimpl()
}

/** CavasSrcRectConstraint
Controls the behavior at the edge of the src-rect, when specified in drawImageRect,
trading off speed for exactness.

When filtering is enabled (in the Paint), skia may need to sample in a neighborhood around
the pixels in the image. If there is a src-rect specified, it is intended to restrict the
pixels that will be read. However, for performance reasons, some implementations may slow
down if they cannot read 1-pixel past the src-rect boundary at times.

This enum allows the caller to specify if such a 1-pixel "slop" will be visually acceptable.
If it is, the caller should pass kFast, and it may result in a faster draw. If the src-rect
must be strictly respected, the caller should pass kStrict. */
type CanvasSrcRectConstraint int

const (
	/**
	If kStrict is specified, the implementation must respect the src-rect
	(if specified) strictly, and will never sample outside of those bounds during sampling
	even when filtering. This may be slower than kFast. */
	KCanvasSrcRectConstraintStrict = iota

	/**
	If kFast is specified, the implementation may sample outside of the src-rect
	(if specified) by half the width of filter. This allows greater flexibility
	to the implementation and can make the draw much faster. */
	KCanvasSrcRectConstraintFast
)

/** DrawImageRect
Draw the specified image, scaling and translating so that it fills the specified
dst rect. If the src rect is non-null, only that subset of the image is transformed
and drawn.

@param image      The image to be drawn
@param src        Optional: specify the subset of the image to be drawn
@param dst        The destination rectangle where the scaled/translated
                  image will be drawn
@param paint      The paint used to draw the image, or NULL
@param constraint Control the tradeoff between speed and exactness w.r.t. the src-rect. */
func (canvas *Canvas) DrawImageRect(image *Image, srcRect, dstRect Rect, paint *Paint, constraint CanvasSrcRectConstraint) {
	toimpl()
}

/** DrawImageNine
Draw the image stretched differentially to fit into dst.
center is a rect within the image, and logically divides the image
into 9 sections (3x3). For example, if the middle pixel of a [5x5]
image is the "center", then the center-rect should be [2, 2, 3, 3].

If the dst is >= the image size, then...
- The 4 corners are not stretched at all.
- The sides are stretched in only one axis.
- The center is stretched in both axes.
Else, for each axis where dst < image,
- The corners shrink proportionally
- The sides (along the shrink axis) and center are not drawn */
func (canvas *Canvas) DrawImageNine(image *Image, enter Rect, dst Rect, paint *Paint) {
	toimpl()
}

/** DrawBitmap
Draw the specified bitmap, with its top/left corner at (x,y), using the
specified paint, transformed by the current matrix. Note: if the paint
contains a maskfilter that generates a mask which extends beyond the
bitmap's original width/height, then the bitmap will be drawn as if it
were in a Shader with CLAMP mode. Thus the color outside of the original
width/height will be the edge color replicated.

If a shader is present on the paint it will be ignored, except in the
case where the bitmap is kAlpha_8_SkColorType. In that case, the color is
generated by the shader.

@param bitmap   The bitmap to be drawn
@param left     The position of the left side of the bitmap being drawn
@param top      The position of the top side of the bitmap being drawn
@param paint    The paint used to draw the bitmap, or NULL */
func (canvas *Canvas) DrawBitmap(bmp *Bitmap, left, top Scalar, paint *Paint) {
	toimpl()
}

/** DrawBitmapRect
Draw the specified bitmap, scaling and translating so that it fills the specified
dst rect. If the src rect is non-null, only that subset of the bitmap is transformed
and drawn.

@param bitmap     The bitmap to be drawn
@param src        Optional: specify the subset of the bitmap to be drawn
@param dst        The destination rectangle where the scaled/translated
                  bitmap will be drawn
@param paint      The paint used to draw the bitmap, or NULL
@param constraint Control the tradeoff between speed and exactness w.r.t. the src-rect. */
func (canvas *Canvas) DrawBitmapRect(bmp *Bitmap, src, dst Rect, paint *Paint, constraint CanvasSrcRectConstraint) {
	toimpl()
}

/** DrawBitmapNine
Draw the bitmap stretched or shrunk differentially to fit into dst.
center is a rect within the bitmap, and logically divides the bitmap
into 9 sections (3x3). For example, if the middle pixel of a [5x5]
bitmap is the "center", then the center-rect should be [2, 2, 3, 3].

If the dst is >= the bitmap size, then...
- The 4 corners are not stretched at all.
- The sides are stretched in only one axis.
- The center is stretched in both axes.
Else, for each axis where dst < bitmap,
- The corners shrink proportionally
- The sides (along the shrink axis) and center are not drawn */
func (canvas *Canvas) DrawBitmapNine(bmp *Bitmap, center Rect, dst Rect, paint *Paint) {
	toimpl()
}

/** CanvasLattice
Specifies coordinates to divide a bitmap into (xCount*yCount) rects. */
type CanvasLattice struct {
	// An array of x-coordinates that divide the bitmap vertically.
	// These must be unique, increasing, and in the set [0, width].
	// Does not have ownership.
	XDivs *int

	// The number of fXDivs.
	XCount int

	// An array of y-coordinates that divide the bitmap horizontally.
	// These must be unique, increasing, and in the set [0, height].
	// Does not have ownership.
	YDivs *int

	// The number of fYDivs.
	YCount int
}

/** DrawBitmapLattice
Draw the bitmap stretched or shrunk differentially to fit into dst.

Moving horizontally across the bitmap, alternating rects will be "scalable"
(in the x-dimension) to fit into dst or must be left "fixed".  The first rect
is treated as "fixed", but it's possible to specify an empty first rect by
making lattice.fXDivs[0] = 0.

The scale factor for all "scalable" rects will be the same, and may be greater
than or less than 1 (meaning we can stretch or shrink).  If the number of
"fixed" pixels is greater than the width of the dst, we will collapse all of
the "scalable" regions and appropriately downscale the "fixed" regions.

The same interpretation also applies to the y-dimension. */
func (canvas *Canvas) DrawBitmapLattice(bmp *Bitmap, lattice *CanvasLattice, dst Rect, paint *Paint) {
	toimpl()
}

func (canvas *Canvas) DrawImageLattice(bmp *Bitmap, lattice *CanvasLattice, dst Rect, paint *Paint) {
	toimpl()
}

/** DrawText
Draw the text, with origin at (x,y), using the specified paint.
The origin is interpreted based on the Align setting in the paint.
@param text The text to be drawn
@param byteLength   The number of bytes to read from the text parameter
@param x        The x-coordinate of the origin of the text being drawn
@param y        The y-coordinate of the origin of the text being drawn
@param paint    The paint used for the text (e.g. color, size, style) */
func (canvas *Canvas) DrawText(text string, x, y Scalar, paint *Paint) {
	toimpl()
}

/** DrawTextAt
Draw the text, with each character/glyph origin specified by the pos[]
array. The origin is interpreted by the Align setting in the paint.
@param text The text to be drawn
@param byteLength   The number of bytes to read from the text parameter
@param pos      Array of positions, used to position each character
@param paint    The paint used for the text (e.g. color, size, style) */
func (canvas *Canvas) DrawTextAt(text string, pos []Point, paint *Paint) {
	toimpl()
}

/** DrawTextAtH
Draw the text, with each character/glyph origin specified by the x
coordinate taken from the xpos[] array, and the y from the constY param.
The origin is interpreted by the Align setting in the paint.
@param text The text to be drawn
@param byteLength   The number of bytes to read from the text parameter
@param xpos     Array of x-positions, used to position each character
@param constY   The shared Y coordinate for all of the positions
@param paint    The paint used for the text (e.g. color, size, style) */
func (canvas *Canvas) DrawTextAtH(text string, xpos []Scalar, constY Scalar, paint *Paint) {
	toimpl()
}

/** DrawTextOnPathHV
Draw the text, with origin at (x,y), using the specified paint, along
the specified path. The paint's Align setting determins where along the
path to start the text.
@param text The text to be drawn
@param byteLength   The number of bytes to read from the text parameter
@param path         The path the text should follow for its baseline
@param hOffset      The distance along the path to add to the text's
					starting position
@param vOffset      The distance above(-) or below(+) the path to
					position the text
@param paint        The paint used for the text */
func (canvas *Canvas) DrawTextOnPathHV(text string, path *Path, hOffset, vOffset Scalar, paint *Paint) {
	toimpl()
}

/** DrawTextOnPath
Draw the text, with origin at (x,y), using the specified paint, along
the specified path. The paint's Align setting determins where along the
path to start the text.
@param text The text to be drawn
@param byteLength   The number of bytes to read from the text parameter
@param path         The path the text should follow for its baseline
@param matrix       (may be null) Applied to the text before it is
					mapped onto the path
@param paint        The paint used for the text */
func (canvas *Canvas) DrawTextOnPath(text string, path *Path, matrix *Matrix, paint *Paint) {
	toimpl()
}

/** DrawTextRSXform
Draw the text with each character/glyph individually transformed by its xform.
If cullRect is not null, it is a conservative bounds of what will be drawn
taking into account the xforms and the paint, and will be used to accelerate culling. */
func (canvas *Canvas) DrawTextRSXform(text string, rsxform []RSXform, cullRect Rect, paint *Paint) {
	toimpl()
}

/** DrawPicture
Draw the picture into this canvas. This method effective brackets the
playback of the picture's draw calls with save/restore, so the state
of this canvas will be unchanged after this call.
@param picture The recorded drawing commands to playback into this
			   canvas.

If matrix is non-null, apply that matrix to the CTM when drawing this picture. This is
logically equivalent to
    save/concat/drawPicture/restore

If paint is non-null, draw the picture into a temporary buffer, and then apply the paint's
alpha/colorfilter/imagefilter/xfermode to that buffer as it is drawn to the canvas.
This is logically equivalent to
    saveLayer(paint)/drawPicture/restore */
func (canvas *Canvas) DrawPicture(pic *Picture, matrix *Matrix, paint *Paint) {
	toimpl()
}

/** DrawShadowedPicture
Draw the picture into this canvas.

We will use the canvas's lights along with the picture information (draw depths of
objects, etc) to first create a set of shadowmaps for the light-picture pairs, and
then use that set of shadowmaps to render the scene with shadows.

If matrix is non-null, apply that matrix to the CTM when drawing this picture. This is
logically equivalent to
    save/concat/drawPicture/restore

If paint is non-null, draw the picture into a temporary buffer, and then apply the paint's
alpha/colorfilter/imagefilter/xfermode to that buffer as it is drawn to the canvas.
This is logically equivalent to
    saveLayer(paint)/drawPicture/restore */
func (canvas *Canvas) DrawShadowedPicture(picture *Picture, matrix *Matrix, paint *Paint) {
	toimpl()
}

type CanvasVertexMode int

const (
	KCanvasVertexModeTriangles CanvasVertexMode = iota
	KCanvasVertexModeTriangleStrip
	KCanvasVertexModeTriangleFan
)

/** DrawVertices
Draw the array of vertices, interpreted as triangles (based on mode).

If both textures and vertex-colors are NULL, it strokes hairlines with
the paint's color. This behavior is a useful debugging mode to visualize
the mesh.

@param vmode How to interpret the array of vertices
@param vertexCount The number of points in the vertices array (and
			corresponding texs and colors arrays if non-null)
@param vertices Array of vertices for the mesh
@param texs May be null. If not null, specifies the coordinate
			in _texture_ space (not uv space) for each vertex.
@param colors May be null. If not null, specifies a color for each
			  vertex, to be interpolated across the triangle.
@param xmode Used if both texs and colors are present. In this
			case the colors are combined with the texture using mode,
			before being drawn using the paint. If mode is null, then
			kModulate_Mode is used.
@param indices If not null, array of indices to reference into the
			vertex (texs, colors) array.
@param indexCount number of entries in the indices array (if not null)
@param paint Specifies the shader/texture if present. */
func (canvas *Canvas) DrawVertices(vmode CanvasVertexMode, vertexCount int, vertices []Point, texs []Point, colors []Color,
	mode *Xfermode, indices []uint16, indexCount int, paint *Paint) {
	toimpl()
}

/** DrawPatch
Draw a cubic coons patch

@param cubic specifies the 4 bounding cubic bezier curves of a patch with clockwise order
			   starting at the top left corner.
@param colors specifies the colors for the corners which will be bilerp across the patch,
			   their order is clockwise starting at the top left corner.
@param texCoords specifies the texture coordinates that will be bilerp across the patch,
			   their order is the same as the colors.
@param xmode specifies how are the colors and the textures combined if both of them are
			   present.
@param paint Specifies the shader/texture if present. */
func (canvas *Canvas) DrawPatch(cubics [12]Point, colors [4]Color, texCoords [4]Point, xmode *Xfermode, paint *Paint) {
	toimpl()
}

/** DrawAtlas
Draw a set of sprites from the atlas. Each is specified by a tex rectangle in the
coordinate space of the atlas, and a corresponding xform which transforms the tex rectangle
into a quad.

    xform maps [0, 0, tex.width, tex.height] -> quad

The color array is optional. When specified, each color modulates the pixels in its
corresponding quad (via the specified SkXfermode::Mode).

The cullRect is optional. When specified, it must be a conservative bounds of all of the
resulting transformed quads, allowing the canvas to skip drawing if the cullRect does not
intersect the current clip.

The paint is optional. If specified, its antialiasing, alpha, color-filter, image-filter
and xfermode are used to affect each of the quads. */
func (canvas *Canvas) DrawAtlas(atlas *Image, form []RSXform, tex []Rect, colors []Color, count int, mode XfermodeMode,
	cullRect Rect, paint *Paint) {
	toimpl()
}

/** DrawDrawable
Draw the contents of this drawable into the canvas. If the canvas is async
(e.g. it is recording into a picture) then the drawable will be referenced instead,
to have its draw() method called when the picture is finalized.

If the intent is to force the contents of the drawable into this canvas immediately,
then drawable->draw(canvas) may be called. */
func (canvas *Canvas) DrawDrawable(drawable *Drawable, matrix *Matrix) {
	toimpl()
}

func (canvas *Canvas) DrawDrawableAt(drawable *Drawable, x, y Scalar) {
	toimpl()
}

/** DrawAnnotation
Send an "annotation" to the canvas. The annotation is a key/value pair, where the key is
a null-terminated utf8 string, and the value is a blob of data stored in an SkData
(which may be null). The annotation is associated with the specified rectangle.

The caller still retains its ownership of the data (if any).

Note: on may canvas types, this information is ignored, but some canvases (e.g. recording
a picture or drawing to a PDF document) will pass on this information. */
func (canvas *Canvas) DrawAnnotation(rect Rect, key []byte, value *Data) {
	toimpl()
}

/** DrawFilter
Get the current filter object. The filter's reference count is not
affected. The filter is saved/restored, just like the matrix and clip.
@return the canvas' filter (or NULL). */
func (canvas *Canvas) DrawFilter() *DrawFilter {
	toimpl()
	return nil
}

/** SetDrawFilter
Set the new filter (or NULL). Pass NULL to clear any existing filter.
As a convenience, the parameter is returned. If an existing filter
exists, its refcnt is decrement. If the new filter is not null, its
refcnt is incremented. The filter is saved/restored, just like the
matrix and clip.
@param filter the new filter (or NULL)
@return the new filter
Impl CanvasImpl */
func (canvas *Canvas) SetDrawFilter(filter *DrawFilter) {
	toimpl()
}

/** IsClipEmpty
Return true if the current clip is empty (i.e. nothing will draw).
Note: this is not always a free call, so it should not be used
more often than necessary. However, once the canvas has computed this
result, subsequent calls will be cheap (until the clip state changes,
which can happen on any clip..() or restore() call.
Impl CanvasImpl */
func (canvas *Canvas) IsClipEmpty() bool {
	toimpl()
	return false
}

/** IsClipRect
Returns true if the current clip is just a (non-empty) rectangle.
Returns false if the clip is empty, or if it is complex. */
func (canvas *Canvas) IsClipRect() bool {
	toimpl()
	return false
}

/** TotalMatrix
Return the current matrix on the canvas.
This does not account for the translate in any of the devices.
@return The current matrix on the canvas. */
func (canvas *Canvas) TotalMatrix() *Matrix {
	return canvas.mcRec.Matrix
}

/** ClipStack
Return the clip stack. The clip stack stores all the individual
clips organized by the save/restore frame in which they were
added.
@return the current clip stack ("list" of individual clip elements) */
func (canvas *Canvas) ClipStack() *ClipStack {
	toimpl()
	return nil
}

type CanvasClipVisitor interface {
}

/** ReplayClips
Replays the clip operations, back to front, that have been applied to
the canvas, calling the appropriate method on the visitor for each
clip. All clips have already been transformed into device space. */
func (canvas *Canvas) ReplayClips(clipVisitor CanvasClipVisitor) {
	toimpl()
}

func (canvas *Canvas) internalAccessTopLayerDrawContext() *GrDrawContext {
	toimpl()
	return nil
}

func (canvas *Canvas) internalSetIgnoreSaveLayerBounds(ignore bool) {
	toimpl()
	return
}

func (canvas *Canvas) internalGetIgnoreSaveLayerBounds() bool {
	toimpl()
	return false
}

func (canvas *Canvas) internalSetTreatSpriteAsBitmap(treatSpriteAsBitmap bool) {
	toimpl()
}

func (canvas *Canvas) internalGetTreatSpriteAsBitmap() bool {
	toimpl()
	return false
}

// TEMP helpers until we switch virtual over to const& for src-rect
func (canvas *Canvas) legacyDrawImageRect(image *Image, src, dst *Rect, paint *Paint) {
	toimpl()
}

func (canvas *Canvas) legacyDrawBitmapRect(bmp *Bitmap, src, dst *Rect, paint *Paint) {
	toimpl()
}

// expose minimum amount of information necessary for transitional refactoring
/**
Returns CTM and clip bounds, translated from canvas coordinates to top layer coordinates. */
func (canvas *Canvas) temporaryInternalDescribeTopLayer(matrix *Matrix, clipBounds *Rect) {
	toimpl()
}

func (canvas *Canvas) Z() Scalar {
	toimpl()
	return 0
}

/** OnNewSurface
default impl defers to getDevice()->newSurface(info)
Impl CanvasImpl */
func (canvas *Canvas) OnNewSurface(imageInfo *ImageInfo, surfaceProps *SurfaceProps) {
	toimpl()
}

/** OnPeekPixels
default impl defers to its device
Impl CanvasImpl */
func (canvas *Canvas) OnPeekPixels(pixmap *Pixmap) bool {
	toimpl()
	return false
}

/** OnAccessTopLayerPixles Impl CanvasImpl */
func (canvas *Canvas) OnAccessTopLayerPixles(pixmap *Pixmap) bool {
	toimpl()
	return false
}

/** OnImageInfo Impl CanvasImpl */
func (canvas *Canvas) OnImageInfo() *ImageInfo {
	toimpl()
	return nil
}

/** OnGetProps Impl CanvasImpl */
func (canvas *Canvas) OnGetProps() (*SurfaceProps, bool) {
	toimpl()
	return nil, false
}

/** WillSave Impl CanvasImpl */
func (canvas *Canvas) WillSave() {
	toimpl()
}

/**
Subclass save/restore notifiers.
Overriders should call the corresponding INHERITED method up the inheritance chain.
getSaveLayerStrategy()'s return value may suppress full layer allocation. */
type CanvasSaveLayerStrategy int

const (
	KCanvasSaveLayerStrategyFullLayer = iota
	KCanvasSaveLayerStrategyNoLayer
)

/** SaveLayerStrategy
Overriders should call the corresponding INHERITED method up the inheritance chain.
Impl CanvasImpl */
func (canvas *Canvas) SaveLayerStrategy() CanvasSaveLayerStrategy {
	toimpl()
	return KCanvasSaveLayerStrategyFullLayer
}

/** WillRestore Impl CanvasImpl */
func (canvas *Canvas) WillRestore() {
	toimpl()
}

/** DidRestore Impl CanvasImpl */
func (canvas *Canvas) DidRestore() {
	toimpl()
}

/** DidConcat Impl CanvasImpl */
func (canvas *Canvas) DidConcat(matrix *Matrix) {
	toimpl()
}

/** DidSetMatrix Impl CanvasImpl */
func (canvas *Canvas) DidSetMatrix(matrix *Matrix) {
	toimpl()
}

/** DidTranslate Impl CanvasImpl */
func (canvas *Canvas) DidTranslate(dx, dy Scalar) {
	toimpl()
}

/** DidTranslateZ Impl CanvasImpl */
func (canvas *Canvas) DidTranslateZ(z Scalar) {
	toimpl()
}

/** OnDrawAnnotation Impl CanvasImpl */
func (canvas *Canvas) OnDrawAnnotation(rect Rect, kay []byte, value *Data) {
	toimpl()
}

/** OnDrawDRect Impl CanvasImpl */
func (canvas *Canvas) OnDrawDRect(outter, inner Rect, paint *Paint) {
	toimpl()
}

/** OnDrawText Impl CanvasImpl */
func (canvas *Canvas) OnDrawText(text string, x, y Scalar, paint *Paint) {
	toimpl()
}

/** OnDrawTextAt Impl CanvasImpl */
func (canvas *Canvas) OnDrawTextAt(text string, xpos []Point, constY Scalar, paint *Paint) {
	toimpl()
}

/** OnDrawTextAtH Impl CanvasImpl */
func (canvas *Canvas) OnDrawTextAtH(text string, xpos []Point, constY Scalar, paint *Paint) {
	toimpl()
}

/** OnDrawTextOnPath Impl CanvasImpl */
func (canvas *Canvas) OnDrawTextOnPath(text string, path *Path, matrix *Matrix, paint *Paint) {
	toimpl()
}

/** OnDrawTextRSXform Impl CanvasImpl */
func (canvas *Canvas) OnDrawTextRSXform(text string, xform []RSXform, cullRect *Rect, paint *Paint) {
	toimpl()
}

/** OnDrawTextBlob Impl CanvasImpl */
func (canvas *Canvas) OnDrawTextBlob(blob *TextBlob, x, y Scalar, paint *Paint) {
	toimpl()
}

/** OnDrawPatch Impl CanvasImpl */
func (canvas *Canvas) OnDrawPatch(cubics [12]Point, colors [4]Color, texCoords [4]Point, xmode *Xfermode, paint *Paint) {
	toimpl()
}

/** OnDrawDrawable Impl CanvasImpl */
func (canvas *Canvas) OnDrawDrawable(drawable *Drawable, matrixe *Matrix) {
	toimpl()
}

/** OnDrawPaint Impl CanvasImpl */
func (canvas *Canvas) OnDrawPaint(paint *Paint) {
	canvas.PredrawRectNotify(nil, paint, KCanvasShaderOverrideOpacityNotOpaque)

	var looper = newAutoDrawLooper(canvas, paint, false, nil)
	for looper.Next(KDrawFilterTypePaint) {
		var it = NewDrawIterator(canvas)
		for it.Next() {
			it.Device().DrawPaint(it.Draw, looper.Paint())
		}
	}
}

/** OnDrawRect Impl CanvasImpl */
func (canvas *Canvas) OnDrawRect(rect Rect, paint *Paint) {
	toimpl()
}

/** OnDrawOval Impl CanvasImpl */
func (canvas *Canvas) OnDrawOval(oval Rect, paint *Paint) {
	toimpl()
}

/** OnDrawArc Impl CanvasImpl */
func (canvas *Canvas) OnDrawArc(oval Rect, startAngle, sweepAngle Scalar, useCenter bool, paint *Paint) {
	toimpl()
}

/** OnDrawPoints Impl CanvasImpl */
func (canvas *Canvas) OnDrawPoints(mode CanvasPointMode, count int, pts []Point, paint *Paint) {
	toimpl()
}

/** OnDrawVertices Impl CanvasImpl */
func (canvas *Canvas) OnDrawVertices(vertexMode CanvasVertexMode, vertexCount int, vertices []Point, texs []Point,
	colors []Color, xfermode *Xfermode, indices []uint16, indexCount int, paint *Paint) {
	toimpl()
}

/** OnDrawAtlas Impl CanvasImpl */
func (canvas *Canvas) OnDrawAtlas(atlas *Image, xform []RSXform, tex []Rect, colors []Color, count int,
	mode XfermodeMode, cull *Rect, paint *Paint) {
	toimpl()
}

/** OnDrawPath Impl CanvasImpl */
func (canvas *Canvas) OnDrawPath(path *Path, paint *Paint) {
	toimpl()
}

/** OnDrawImage Impl CanvasImpl */
func (canvas *Canvas) OnDrawImage(image *Image, dx, dy Scalar, paint *Paint) {
	toimpl()
}

/** OnDrawImageRect Impl CanvasImpl */
func (canvas *Canvas) OnDrawImageRect(image *Image, src *Rect, dst Rect, paint *Paint,
	constraint CanvasSrcRectConstraint) {
	toimpl()
}

/** OnDrawImageNine Impl CanvasImpl */
func (canvas *Canvas) OnDrawImageNine(image *Image, center Rect, dst Rect, paint *Paint) {
	toimpl()
}

/** OnDrawImageLattice Impl CanvasImpl */
func (canvas *Canvas) OnDrawImageLattice(image *Image, lattice *CanvasLattice, dst Rect, paint *Paint) {
	toimpl()
}

/** OnDrawBitmap Impl CanvasImpl */
func (canvas *Canvas) OnDrawBitmap(bmp *Bitmap, dx, dy Scalar, paint *Paint) {
	toimpl()
}

/** OnDrawBitmapRect Impl CanvasImpl */
func (canvas *Canvas) OnDrawBitmapRect(bmp *Bitmap, src *Rect, dst Rect, paint *Paint,
	constraint CanvasSrcRectConstraint) {
	toimpl()
}

/** OnDrawBitmapNine Impl CanvasImpl */
func (canvas *Canvas) OnDrawBitmapNine(bmp *Bitmap, center Rect, dst Rect, paint *Paint) {
	toimpl()
}

type ClipEdgeStyle int

const (
	KClipEdgeStyleHard = iota
	KClipEdgeStyleSoft
)

/** OnClipRect Impl CanvasImpl */
func (canvas *Canvas) OnClipRect(rect Rect, op RegionOp, edgeStyle ClipEdgeStyle) {
	toimpl()
}

/** OnClipPath Impl CanvasImpl */
func (canvas *Canvas) OnClipPath(path *Path, op RegionOp, edgeStyle ClipEdgeStyle) {
	toimpl()
}

/** OnClipRegion Impl CanvasImpl */
func (canvas *Canvas) OnClipRegion(deviceRgn *Region, op RegionOp) {
	toimpl()
}

/** OnDiscard Impl CanvasImpl */
func (canvas *Canvas) OnDiscard() {
	toimpl()
}

/** OnDrawPicture Impl CanvasImpl */
func (canvas *Canvas) OnDrawPicture(pic *Picture, matrix *Matrix, paint *Paint) {
	toimpl()
}

/** OnDrawShadowedPicture Impl CanvasImpl */
func (canvas *Canvas) OnDrawShadowedPicture(pic *Picture, matrix *Matrix, paint *Paint) {
	toimpl()
}

/** CanvasForDrawIterator Impl CanvasImpl */
func (canvas *Canvas) CanvasForDrawIterator() *Canvas {
	return canvas
}

/** ClipRectBounds
Clip rectangle bounds. Called internally by saveLayer.
returns false if the entire rectangle is entirely clipped out
If non-NULL, The imageFilter parameter will be used to expand the clip
and offscreen bounds for any margin required by the filter DAG. */
func (canvas *Canvas) ClipRectBounds(bounds *Rect, saveLayerFlags CanvasSaveLayerFlags, intersection *Rect,
	imageFilter *ImageFilter) {
	toimpl()
}

/** LayerIterator
After calling saveLayer(), there can be any number of devices that make
up the top-most drawing area. LayerIter can be used to iterate through
those devices. Note that the iterator is only valid until the next API
call made on the canvas. Ownership of all pointers in the iterator stays
with the canvas, so none of them should be modified or deleted. */
type LayerIterator struct {
	defaultPaint *Paint
	done         bool
	impl         *DrawIterator
}

/** NewLayerIterator
Initialize iterator with canvas, and set values for 1st device. */
func NewLayerIterator(canvas *Canvas) *LayerIterator {
	toimpl()
	return nil
}

/** Done
Return true if the iterator is done */
func (iter *LayerIterator) Done() bool {
	toimpl()
	return false
}

/** Next
Cycle to the next device */
func (iter *LayerIterator) Next() {
	toimpl()
}

/** Device
These reflect the current device in the iterator */
func (iter *LayerIterator) Device() *BaseDevice {
	toimpl()
	return nil
}

func (iter *LayerIterator) Matrix() *Matrix {
	toimpl()
	return nil
}

func (iter *LayerIterator) Clip() *RasterClip {
	toimpl()
	return nil
}

func (iter *LayerIterator) Paint() *Paint {
	toimpl()
	return nil
}

func (iter *LayerIterator) X() int {
	toimpl()
	return 0
}

func (iter *LayerIterator) Y() int {
	toimpl()
	return 0
}

func CanvasBoundsAffectsClip(saveLayerFlags CanvasSaveLayerFlags) {
	toimpl()
}

func CanvasLegacySaveFlagsToSaveLayerFlags(legacySaveFlags uint32) CanvasSaveLayerFlags {
	toimpl()
	return 0
}

func CanvasDrawDeviceWithFilter(src *BaseDevice, filter *ImageFilter, dst *BaseDevice, ctm *Matrix,
	clipStack *ClipStack) {
	toimpl()
}

type CanvasShaderOverrideOpacity int

const (
	KCanvasShaderOverrideOpacityNone      = 1 << iota // there is no overriding shader (bitmap or image)
	KCanvasShaderOverrideOpacityOpaque                // the overriding shader is opaque
	KCanvasShaderOverrideOpacityNotOpaque             // the overriding shader may not be opaque
)

/** PredrawNotify
notify our surface (if we have one) that we are about to draw, so it
can perform copy-on-write or invalidate any cached images */
func (canvas *Canvas) PredrawRectNotify(rect *Rect, paint *Paint, overrideOpacity CanvasShaderOverrideOpacity) {
	if canvas.surfaceProps != nil {
		var mode = KSurfaceContentChangeModeRetain
		/*
		Since willOverwriteAllPixels() may not be complete free to call, we only do so if
		there is an outstanding snapshot, since w/o that, there will be no copy-on-write
		and therefore we don't care which mode we're in. */
		if canvas.surfaceProps.OutstandingImageSnapshot() != nil {
			if canvas.WouldOverwriteEntireSurface(rect, paint, overrideOpacity) {
				mode = KSurfaceContentChangeModeDiscard
			}
		}
		canvas.surfaceProps.AboutToDraw(mode)
	}
}

/** the first N recs that can fit here mean we won't call malloc */
const (
	KMCRecSize    = 128 // < most recent measurement
	KMCRecCount   = 32  // < common depth for save/restores
	KDeviceCMSize = 176 // < most recent measurement
)

func (canvas *Canvas) updateDeviceCMCache() {
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

func (canvas *Canvas) doSave() {
	toimpl()
}

func (canvas *Canvas) checkForDeferredSave() {
	toimpl()
}

func (canvas *Canvas) internalSetMatrix(matrix *Matrix) {
	toimpl()
}

type CanvasInitFlags int

const (
	KCanvasInitFlagDefault = CanvasInitFlags(1 << iota)
	KCanvasInitFlagConservativeRasterClip
)

/**
Creates a canvas of the specified dimensions, but explicitly not backed
by any device/pixels. Typically this use used by subclasses who handle
the draw calls in some other way. */
func NewCanvas(width, height int, surfaceProps *SurfaceProps) *Canvas {
	toimpl()
	return &Canvas{}
}

/**
Construct a canvas with the specified device to draw into.
@param device   Specifies a device for the canvas to draw into. */
func NewCanvasFromDevice(device *BaseDevice) *Canvas {
	toimpl()
	return &Canvas{}
}

func (canvas *Canvas) resetForNextPicture(bounds Rect) {
	toimpl()
}

/**
call this each time we attach ourselves to a device
 - constructor
 - internalSaveLayer */
func (canvas *Canvas) setupDevice(device *BaseDevice) {
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

/** getTopLayerBounds gets the bounds of the top level layer in global canvas coordinates.
We don't want this to be public because it exposes decisions about layer sizes that are
internal to the canvas. */
func (canvas *Canvas) getTopLayerBounds() Rect {
	toimpl()
	return RectZero
}

func (canvas *Canvas) internalSaveLayer(rec *CanvasSaveLayerRec, strategy CanvasSaveLayerStrategy) {
	toimpl()
}

func (canvas *Canvas) internalRestore() {
	toimpl()
}

type LazyPaint Lazy

/** WouldOverwriteEntireSurface Returns true if drawing the specified rect (or all
if it is null) with the specified paint (or default if null) would overwrite the
entire root device of the canvas (i.e. the canvas' surface if it had one). */
func (canvas *Canvas) WouldOverwriteEntireSurface(rect *Rect, paint *Paint,
	overrideOpacity CanvasShaderOverrideOpacity) bool {
	toimpl()
	return false
}

/** CanDrawBitmapAsSprite Returns true if the paint's imagefilter can be invoked
directly, without needed a layer. */
func (canvas *Canvas) CanDrawBitmapAsSprite(x, y Scalar, w, h int, paint *Paint) {
	toimpl()
}

/** tDeviceCM is the record we keep for each BaseDevice that the user installs.
The clip/matrix/proc are fields that reflect the top of the save/resotre
stack. Whenever the canvas changes, it makes a dirty flag, and then before
these are used (assuming we're not on a layer) we rebuild these cache values:
they reflect the top of the save stack, but translated and clipped by the
device's XY offset and bitmap-bounds. */
type tDeviceCM struct {
	Next          *tDeviceCM
	Device        *BaseDevice
	Clip          *RasterClip
	Paint         *Paint
	Matrix        *Matrix
	MatrixStroage *Matrix
	StashedMatrix *Matrix
}

func newDeivceCM(device *BaseDevice, paint *Paint, canvas *Canvas, conservativeRasterClip bool, stashed *Matrix) *tDeviceCM {
	var deviceCM = new(tDeviceCM)
	deviceCM.Next = nil
	deviceCM.Clip = NewRasterClip(conservativeRasterClip)
	deviceCM.StashedMatrix = stashed
	deviceCM.Device = device
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

func (deviceCM *tDeviceCM) UpdateMC(totalMatrix *Matrix, totalClip *RasterClip,
	clipStack *ClipStack, updateClip *RasterClip) {
	var x, y = deviceCM.Device.Origin().X, deviceCM.Device.Origin().Y
	var w, h = deviceCM.Device.Width(), deviceCM.Device.Height()
	if x == 0 || y == 0 {
		deviceCM.Matrix = totalMatrix
		deviceCM.Clip = totalClip
	} else {
		deviceCM.Matrix = NewMatrixClone(totalMatrix)
		deviceCM.Matrix.PostTranslate(-x, -y)
		totalClip.Translate(-x, -y, deviceCM.Clip)
	}
	deviceCM.Clip.Op(MakeRectWH(w, h), KRegionOpIntersect)

	// Intersect clip, but don't translate it (yet)

	if updateClip != nil {
		updateClip.Op(MakeRect(x, y, w, h), KRegionOpDifference)
	}

	deviceCM.Device.SetMatrixClip(deviceCM.Matrix, deviceCM.Clip.ForceGetBW(), clipStack)
}

/** tCanvasMCRec is the record we keep for each save/restore level in the stack.
Since a level optionally copies the matrix and/or stack, we have pointers
for these fields. If the value is copied for this level, the copy is stored
in the ...Storage field, and the pointer points to that. If the value is not
copied for this level, we ignore ...Storage, and just point at the
corresponding value in the previous level in the stack. */
type tCanvasMCRec struct {
	Filter *DrawFilter // the current filter (or nil)
	Layer  *tDeviceCM

	/** If there are any layers in the stack, this points to the top-most
	one that is at or below this level in the stack (so we know what
	bitmap/device to draw into from this level. This value is NOT
	reference counted, since the real owner is either our fLayer field,
	or a previous one in a lower level.) */
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
	canvas.updateDeviceCMCache()
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

/** If the paint has an imagefilter, but it can be simplified to just a colorfilter, return that
colorfilter, else return nullptr. */
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
	if bounds.IsEmpty() {
		return MakeRectEmpty()
	}

	// Expand bounds out by 1 in case we are anti-aliasing. We store the bounds as floats to
	// enable a faster quick reject implementation.
	bounds.InsetLTRB(-1, -1, 1, 1)

	return bounds
}
