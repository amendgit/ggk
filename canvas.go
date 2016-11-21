package ggk

import (
	"container/list"
	"runtime"
)

type CanvasInitFlags int

const (
	KCanvasInitFlagDefault = 1 << iota
	KCanvasInitFlagConservativeRasterClip
)

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
}

// Construct a canvas with the specified bitmap to draw into.
// @param bitmap   Specifies a bitmap for the canvas to draw into. Its
//                 structure are copied to the canvas.
func NewCanvas(bmp *Bitmap) *Canvas {
	var canvas = new(Canvas)
	canvas.surfaceProps = MakeSurfaceProps(KSurfacePropsFlagNone, KSurfacePropsInitTypeLegacyFontHost)
	canvas.mcStack = list.New()
	var device = NewBitmapDevice(bmp, canvas.surfaceProps)
	canvas.init(device.BaseDevice, KCanvasInitFlagDefault)
	return canvas
}

func (canvas *Canvas) init(dev *BaseDevice, flags CanvasInitFlags) {
	toimpl()
}

func (canvas *Canvas) ReadPixelsToBitmap(bmp *Bitmap, x, y Scalar) error {
	toimpl()
	return nil
}

func (canvas *Canvas) ReadPixelsInRectToBitmap(bmp *Bitmap, srcRect Rect) error {
	toimpl()
	return nil
}

// ReadPixels copy the pixels from the base-layer into the specified buffer
// (pixels + rowBytes). converting them into the requested format (ImageInfo).
// The base-layer are read starting at the specified (srcX, srcY) location in
// the coordinate system of the base-layer.
//
// The specified ImageInfo and (srcX, srcY) offset specifies a source rectangle.
//
//     srcR.SetXYWH(srcX, srcY, dstInfo.Width(), dstInfo.Height())
//
// srcR is intersected with the bounds of the base-layer. If this intersection
// is not empty, then we have two sets of pixels (of equal size). Replace the
// dst pixels with the corresponding src pixels, performing any
// colortype/alphatype transformations needed (in the case where the src and dst
// have different colortypes or alphatypes).
//
// This call can fail, returning false, for serveral reasons:
// - If srcR does not intersect the base-layer bounds.
// - If the requested colortype/alphatype cannot be converted from the base-layer's types.
// - If this canvas is not backed by pixels (e.g. picture or PDF)
func (c *Canvas) ReadPixels(dstInfo *ImageInfo, dstData []byte, rowBytes int,
	x, y Scalar) error {
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

// WritePixels affects the pixels in the base-layer, and operates in pixel
// coordinates. ignoring the matrix and clip.
//
// The specified ImageInfo and (x, y) offset specifies a rectangle: target.
//
//     target.SetXYWH(x, y, info.width(), info.height());
//
// Target is intersected with the bounds of the base-layer. If this intersection
// is not empty. then we have two sets of pixels (of equal size), the "src"
// specified by info+pixels+rowBytes and the "dst" by the canvas' backend.
// Replace the dst pixels with the corresponding src pixels, performing any
// colortype/alphatype transformations needed (in the case where the src and
// dst have different colirtypes or alphatypes).
//
// This call can fail, returing false, for several reasons:
// - If the src colortype/alpahtype cannot be converted to the canvas' types
// - If this canvas is not backed by pixels (e.g. picture or PDF)
func WritePixels(info *ImageInfo, pixels []byte, rowBytes int, x, y int) error {
	return nil
}

func (c *Canvas) Device() *BaseDevice {
	var rec = c.mcStack.Front().Value.(*tCanvasMCRec)
	return rec.Layer.Device
}

func (c *Canvas) BaseLayerSize() Size {
	var dev = c.Device()
	var sz Size
	if dev != nil {
		sz = MakeSize(dev.Width(), dev.Height())
	}
	return sz
}

type CanvasPointMode int

const (
	// DrawPoints draws each point separately
	KCanvasPointModePoints CanvasPointMode = iota
	// DrawPoints draws each pair of points as a line segment
	KCanvasPointModeLines
	// DrawPoints draws the array of points as a polygon
	KCanvasPointModePolygon
)

func (canvas *Canvas) DrawPoint(x, y Scalar, paint *Paint) {
	var pt Point
	pt.X, pt.Y = x, y
	canvas.DrawPoints(KCanvasPointModePoints, 1, []Point{pt}, paint)
}

func (canvas *Canvas) DrawColor(color Color, mode XfermodeMode) {
	var paint = NewPaint()
	paint.SetColor(color)
	if KXfermodeModeSrcOver == mode {
		paint.SetXfermodeMode(mode)
	}
	canvas.DrawPaint(paint)
}

func (canvas *Canvas) DrawPoints(mode CanvasPointMode, count int, pts []Point, paint *Paint) {
	canvas.OnDrawPoints(mode, count, pts, paint)
}

func (canvas *Canvas) DrawPaint(paint *Paint) {
	toimpl()
}

func (canvas *Canvas) OnDrawPoints(mode CanvasPointMode, count int, pts []Point, paint *Paint) {
	// if count <= 0 {
	// 	return
	// }
	// var r, storage Rect
	// var bounds *Rect = nil
	// if paint.CanComputeFastBounds() {
	// 	// special-case 2 points (common for drawing a single line)
	// 	if count == 2 {
	// 		r.SetLTRBPoint(pts[0], pts[1])
	// 	} else {
	// 		// TODO(impl)
	// 		// r.SetPoints(pts, count)
	// 	}
	// 	if canvas.QuickReject(paint.ComputeFastStrokeBounds(r, &storage)) {
	// 		return
	// 	}
	// 	bounds = &r
	// }
	// canvas.PredrawNotify()
	// var looper = newAutoDrawLooper(canvas, paint, false, bounds)
	// defer looper.Finalizer()
	// for looper.Next(KDrawFilterTypePoint) {
	// 	var iter = newDrawIter(canvas)
	// 	for iter.Next() {
	// 		iter.Device().DrawPoints(iter, mode, count, pts, looper.Paint())
	// 	}
	// }
}

func (canvas *Canvas) QuickReject(src Rect) bool {
	toimpl()
	return false
}

func (canvas *Canvas) PredrawNotify() {
	toimpl()
}

func (canvas *Canvas) SaveCount() int {
	return canvas.saveCount
}

func (canvas *Canvas) internalSaveLayer(rec *tCanvasSaveLayerRec, strategy CanvasSaveLayerStrategy) {
	toimpl()
}

func (canvas *Canvas) internalRestore() {
	toimpl()
}

type CanvasSaveFlags int

const (
	KSaveFlagHasAlphaLayer   = 0x01
	KSaveFlagFullColorLayer  = 0x02
	KSaveFlagClipToLayer     = 0x10
	KSaveFlagARGBNoClipLayer = 0x0F
	KSaveFlagARGBClipLayer   = 0x1F
)

// tDeviceCM is the record we keep for each BaseDevice that the user installs.
// The clip/matrix/proc are fields that reflect the top of the save/resotre
// stack. Whenever the canvas changes, it makes a dirty flag, and then before
// these are used (assuming we're not on a layer) we rebuild these cache values:
// they reflect the top of the save stack, but translated and clipped by the
// device's XY offset and bitmap-bounds.
type tDeviceCM struct {
	Next                 *tDeviceCM
	Device               *BaseDevice
	Clip                 *RasterClip
	Paint                *Paint
	Matrix               *Matrix
	MatrixStroage        *Matrix
	DeviceIsBitmapDevice bool
}

func newDeivceCM(dev *BaseDevice, paint *Paint, canvas *Canvas,
	conservativeRasterClip bool, deviceIsBitmapDevice bool) {
	var deviceCM = new(tDeviceCM)
	deviceCM.Next = nil
	deviceCM.Clip = NewRasterClip(conservativeRasterClip)
	deviceCM.DeviceIsBitmapDevice = deviceIsBitmapDevice
	if dev != nil {
		dev.OnAttachToCanvas(canvas)
	}
	runtime.SetFinalizer(deviceCM, func(d *tDeviceCM) {
		d.Device.OnDetachFromCanvas()
	})
	deviceCM.Device = dev
	deviceCM.Paint = paint
}

func (d *tDeviceCM) Reset(bounds Rect) {
	d.Clip.SetRect(bounds)
}

func (d *tDeviceCM) UpdateMC(totalMatrix *Matrix, totlaClip *RasterClip,
	clipStack *ClipStack, updateClip *RasterClip) {
	toimpl()
}

// tCanvasMCRec is the record we keep for each save/restore level in the stack.
// Since a level optionally copies the matrix and/or stack, we have pointers
// for these fields. If the value is copied for this level, the copy is stored
// in the ...Storage field, and the pointer points to that. If the value is not
// copied for this level, we ignore ...Storage, and just point at the
// corresponding value in the previous level in the stack.
type tCanvasMCRec struct {
	Filter *DrawFilter // the current filter (or nil)
	Layer  *tDeviceCM

	// If there are any layers in the stack, this points to the top-most
	// one that is at or below this level in the stack (so we know what
	// bitmap/device to draw into from this level. This value is NOT
	// reference counted, since the real owner is either our fLayer field,
	// or a previous one in a lower level.)
	TopLayer          *tDeviceCM
	RasterClip        *RasterClip
	Matrix            *Matrix
	DeferredSaveCount int
}

func newCanvasMCRec(conservativeRasterClip bool) *tCanvasMCRec {
	var rec = new(tCanvasMCRec)
	rec.RasterClip = NewRasterClip(conservativeRasterClip)
	rec.Filter = nil
	rec.Layer = nil
	rec.TopLayer = nil
	rec.Matrix.Reset()
	rec.DeferredSaveCount = 0
	// don't bother initializing Next
	// inc_rec()
	toimpl()
	return nil
}

// Subclass save/restore notifiers.
// Overriders should call the corresponding INHERITED method up the inheritance chain.
// getSaveLayerStrategy()'s return value may suppress full layer allocation.
type CanvasSaveLayerStrategy int

const (
	KCanvasSaveLayerStrategyFullLayer = iota
	KCanvasSaveLayerStrategyNoLayer
)

type tAutoDrawLooper struct {
	lazyPaintInit           *LazyPaint
	lazyPaintPerLooper      *LazyPaint
	canvas                  *Canvas
	origPaint               *Paint
	paint                   *Paint
	filter                  *DrawFilter
	saveCount               int
	tempLayerForImageFilter bool
	done                    bool
	isSimple                bool
	drawLooperContext       *tDrawLooperContext
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
		// We implement ImageFilters for a given draw by creating a layer, then applying the
		// imagefilter to the pixels of that layer (its backing surface/image), and then
		// we call restore() to xfer that layer to the main canvas.
		//
		// 1. SaveLayer (with a paint containing the current imagefilter and xfermode)
		// 2. Generate the src pixels:
		//     Remove the imagefilter and the xfermode from the paint that we (AutoDrawLooper)
		//     return (fPaint). We then draw the primitive (using srcover) into a cleared
		//     buffer/surface.
		// 3. Restore the layer created in #1
		//     The imagefilter is passed the buffer/surface from the layer (now filled with the
		//     src pixels of the primitive). It returns a new "filtered" buffer, which we
		//     draw onto the previous layer using the xfermode from the original paint.
		var tmp = NewPaint()
		tmp.SetImageFilter(looper.paint.ImageFilter())
		tmp.SetXfermode(looper.paint.Xfermode())
		// var storage Rect
		if rawBounds != nil {
			// Make rawBounds include all paint outsets except for those due to image filters.
			// rawBounds = applyPaintToBoundsSansImageFilter(looper.paint, *rawBounds, &storage)
		}
		canvas.internalSaveLayer(newCanvasSaveLayerRec(rawBounds, tmp, nil, KCanvasSaveLayerFlagIsOpaque), KCanvasSaveLayerStrategyFullLayer)
		looper.tempLayerForImageFilter = false
		// we remove the imagefilter/xfermode inside doNext()
	}
	if paint.Looper() != nil {
		// looper.drawLooperContext = paint.Looper().CreateContext(canvas)
		looper.isSimple = false
	} else {
		looper.drawLooperContext = nil
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
		return looper.DoNext(drawType)
	}
}

func (looper *tAutoDrawLooper) DoNext(drawType DrawFilterType) bool {
	toimpl()
	return false
}

func (looper *tAutoDrawLooper) Finalizer() {
	if looper.tempLayerForImageFilter {
		looper.canvas.internalRestore()
	}
}

type CanvasSaveLayerFlags int

const (
	KCanvasSaveLayerFlagIsOpaque = 1 << iota
	KCanvasSaveLayerFlagPreserveLCDText
	kCanvasSaveLayerFlagDontClipToLayer   // private
	KCanvasSaveLayerDontClipToLayerLegacy = kCanvasSaveLayerFlagDontClipToLayer
)

type tCanvasSaveLayerRec struct {
	bounds         *Rect
	paint          *Paint
	backdrop       *ImageFilter
	saveLayerFlags CanvasSaveLayerFlags
}

func newCanvasSaveLayerRec(bounds *Rect, paint *Paint, backdrop *ImageFilter,
	saveLayerFlags CanvasSaveLayerFlags) *tCanvasSaveLayerRec {
	return &tCanvasSaveLayerRec{
		bounds:         bounds,
		paint:          paint,
		backdrop:       backdrop,
		saveLayerFlags: saveLayerFlags,
	}
}

type tDrawIter struct {
	*Draw
	canvas    *Canvas
	currLayer *tDeviceCM
	paint     *Paint // May be nil
}

func newDrawIter(canvas *Canvas) *tDrawIter {
	return nil
}

func (iter *tDrawIter) Next() bool {
	toimpl()
	return false
}

type LazyPaint struct {
}

func imageToColorFilter(paint *Paint) *ColorFilter {
	toimpl()
	return nil
}

func setIfNeeded(lazyPaint *LazyPaint, paint *Paint) *Paint {
	toimpl()
	return nil
}

func applyPaintToBoundsSansImageFilter(paint *Paint, rowBounds Rect, storage Rect) Rect {
	toimpl()
	return RectZero
}
