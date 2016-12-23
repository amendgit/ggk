package ggk

type BitmapDevice struct {
	*BaseDevice

	bitmap *Bitmap
}

// Construct a new device with the specified bitmap as its backend. It is valid
// for the bitmap to have no pixels associated with it. In that case, any
// drawing to this device will have no effect.
func NewBitmapDevice(bitmap *Bitmap, props *SurfaceProps) *BitmapDevice {
	var device = &BitmapDevice{
		BaseDevice: NewBaseDevice(),
		bitmap: bitmap,
	}
	device.Device = device
	return device
}

func (dev *BitmapDevice) ImageInfo() *ImageInfo {
	return dev.bitmap.Info()
}

func (bmpdev *BitmapDevice) OnAccessPixels(pixmap *Pixmap) bool {
	if bmpdev.Device.OnPeekPixels(pixmap) {
		bmpdev.bitmap.NotifyPixelsChanged()
		return true
	}
	return false
}

func (bmpdev *BitmapDevice) OnPeekPixels(pixmap *Pixmap) bool {
	var info = bmpdev.bitmap.Info()
	if bmpdev.bitmap.PixelBytes() != nil && info.ColorType() != KColorTypeUnknown {
		var colorTable *ColorTable = nil
		pixmap.Reset(info, bmpdev.bitmap.PixelBytes(), bmpdev.bitmap.RowBytes(), colorTable)
		return true
	}
	return false
}

func (dev *BitmapDevice) DrawRect(draw *Draw, rect Rect, paint *Paint) {
	draw.DrawRect(rect, paint)
	return
}

func (dev *BitmapDevice) DrawPoints(draw *Draw, mode CanvasPointMode, count int, pts []Point, paint *Paint) {
	draw.DrawPoints(mode, count, pts, paint, false)
}

func (bmpdev *BitmapDevice) DrawPaint(draw *Draw, paint *Paint) {
	toimpl()
}