package ggk

type Device interface {
	// ImageInfo returns ImageInfo for this device, If the canvas is not backed
	// by pixels (cpu or gpu), then the info's ColorType will be
	// KColorTypeUnknown.
	ImageInfo() *ImageInfo

	// AccessGPURenderTarget return the device's gpu render target or nil.
	// AccessGPURenderTarget() *GPURenderTarget

	// OnAttachToCanvas is invoked whenever a device is installed in a canvas
	// (i.e., SetDevice, SaveLayer (for the new device created by the save),
	// and Canvas' BaseDevice & Bitmap - taking ctors). It allows the device
	// to prepare for drawing (e.g., locking their pixels, etc.)
	// OnAttachToCanvas(*Canvas)

	// OnDetachFromCanvas()
	// SetMatrixClip(*Matrix, *Region, *ClipStack)
	DrawPaint(draw *Draw, paint *Paint)
	DrawPoints(draw *Draw, mode CanvasPointMode, count int, pts []Point, paint *Paint)
	DrawRect(draw *Draw, rect Rect, paint *Paint)
	// DrawOval(draw *Draw, oval Rect, paint *Paint)
	// DrawRRect(draw *Draw, RRect, *Paint)
	// DrawDRRect(*Draw, outer, inner RRect, *Paint)
	// DrawPath(draw *Draw, path *Path, mat *Matrix, paint *Paint)
	// DrawSprite(draw *Draw, bmp *Bitmap, x, y int, paint *Paint)
	// DrawBitmapRect(draw *Draw, bmp *Bitmap, srcOrNil *Rect, dst Rect, paint *Paint) (finalDst Rect)
	// DrawBitmapNine(draw *Draw, bmp *Bitmap, center Rect, dst Rect, paint *Paint)
	// DrawImage(draw *Draw, image *Image, x, y Scalar, paint *Paint)
	// DrawImageRect(draw *Draw, image *Image, src Rect, dst Rect, paint *Paint, SrcRectConstraint)
	// DrawText(draw *Draw, text string, x, y Scalar, paint *Paint)
	// DrawPosText(draw *Draw, text string, pos []Scalar, paint *Paint)
	// DrawVertices(Draw, VertexMode, vertexCount int, verts []Point, texs []Point, colors []Color, xmode *Xfermode, indices []uint16, indexCount int, Paint)
	// DrawTextBlob(Draw, TextBlob, x, y Scalar, Paint, DrawFilter)
	// DrawPatch(Draw, cubics [12]Point, colors []Color, texCoords [4]Point, xmode Xfermode, Paint)
	// DrawAtlas(Draw, atlas Image, []RSXform, []Rect, []Color, count int, XfermodeMode, Paint)
	// DrawDevice(draw *Draw, dev Device, x, y int, paint *Paint)
	// DrawTextOnPath(draw *Draw, texts []string, len int, path *Path, mat *Matrix, paint *Paint)

	// OnAccessBitmap() *Bitmap
	// CanHandleImageFilter(*ImageFilter) bool
	// FilterImage(filter *ImageFilter, bmp *Bitmap, ctxt *ImageFilterContext) (resultBmp *Bitmap, offset Point, ok bool)
	// OnPeekPixels(pixmap *Pixmap) bool
	// OnReadPixels(imageInfo ImageInfo, pixelBytes []byte, x, y int)
	// OnWritePixels(imageInfo ImageInfo, pixelBytes []byte, x, y int)
	// OnAccessPixels(pixmap *Pixmap) bool
	// OnCreateDevice(CreateInfo, Paint) Device
	// Flush()
	// GetImageFilterCache() *ImageFilterCache

	// used to change the backend's pixels (and possibly config/rowbytes)
	// but cannot change the width/height, so there should be no change to
	// any clip information.
	// TODO: move to SkBitmapDevice
	//replaceBitmapBackendForRasterSurface(bmp *Bitmap)

	forceConservativeRasterClip() bool

	// Causes any deferred drawing to the device to be completed.
	// flush()

}

type BaseDevice struct {
	Device Device
	origin Point
	imageInfo *ImageInfo
}

func NewBaseDevice() *BaseDevice {
	var baseDevice = &BaseDevice {
		imageInfo: NewImageInfoUnknown(0, 0),
	}
	return baseDevice
}

func (b *BaseDevice) Width() Scalar {
	return b.imageInfo.Width()
}

func (b *BaseDevice) Height() Scalar {
	return b.imageInfo.Height()
}

func (b *BaseDevice) ImageInfo() *ImageInfo {
	return b.imageInfo
}

/**
 *  Return the bounds of the device in the coordinate space of the root
 *  canvas. The root device will have its top-left at 0,0, but other devices
 *  such as those associated with saveLayer may have a non-zero origin.
 */
func (b *BaseDevice) GlobalBounds() Rect {
	var origin = b.Origin()
	var bounds Rect
	bounds.SetXYWH(origin.X, origin.Y, b.Width(), b.Height())
	return bounds
}

/**
 *  Return the device's origin: its offset in device coordinates from
 *  the default origin in its canvas' matrix/clip
 */
func (b *BaseDevice) Origin() Point {
	return b.origin
}

func (b *BaseDevice) OnAttachToCanvas(canvas *Canvas) {
	toimpl()
}

func (b *BaseDevice) OnDetachFromCanvas() {
	toimpl()
}

func (b *BaseDevice) ReadPixels(info *ImageInfo, pixels []byte, rowBytes int,
	x, y Scalar) error {
	toimpl()
	return nil
}

func (b *BaseDevice) AccessPixels(pixmap *Pixmap) bool {
	toimpl()
	return false
}

func (b *BaseDevice) DrawRect(draw *Draw, rect Rect, paint *Paint) {
	toimpl()
	return
}

func (b *BaseDevice) DrawPoints(draw *Draw, mode CanvasPointMode, count int, pts []Point, paint *Paint) {
	toimpl()
	return
}

func (b *BaseDevice) DrawPaint(draw *Draw, paint *Paint) {
	toimpl()
}

func (b *BaseDevice) forceConservativeRasterClip() bool {
	return false
}