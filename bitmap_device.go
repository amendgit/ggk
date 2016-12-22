package ggk

type BitmapDevice struct {
	*BaseDevice
}

// Construct a new device with the specified bitmap as its backend. It is valid
// for the bitmap to have no pixels associated with it. In that case, any
// drawing to this device will have no effect.
func NewBitmapDevice(bmp *Bitmap, props *SurfaceProps) *BitmapDevice {
	var device = &BitmapDevice{
		BaseDevice: NewBaseDevice(),
	}
	device.Device = device
	return device
}

func (dev *BitmapDevice) ImageInfo() *ImageInfo {
	toimpl()
	return nil
}

func (dev *BitmapDevice) DrawRect(draw *Draw, rect Rect, paint *Paint) {
	draw.DrawRect(rect, paint)
	return
}

func (dev *BitmapDevice) DrawPoints(draw *Draw, mode CanvasPointMode, count int, pts []Point, paint *Paint) {
	draw.DrawPoints(mode, count, pts, paint, false)
}
