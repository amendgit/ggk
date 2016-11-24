package ggk

type Draw struct {
	dst        *Pixmap
	matrix     *Matrix
	region     *Region
	rasterClip *RasterClip
	clipStack  *ClipStack
	device     *BaseDevice
	// procs      *DrawProcs
}

type tBitmapXferProc interface {
	Xfer(pixels []byte, data uint32)
}

type tBitmapXferClear int
func (*tBitmapXferClear)Xfer(pixels []byte, data uint32) {
	for i := 0; i < len(pixels); i++ {
		pixels[i] = 0
	}
}

type tBitmapXferDst int
func (*tBitmapXferDst)Xfer(pixels []byte, data uint32) {
	toimpl()
}

type tBitmapXferSrc32 int
func (*tBitmapXferSrc32) Xfer(pixels []byte, data uint32) {
	toimpl()
}

type tBitmapXferSrc16 int
func (*tBitmapXferSrc16) Xfer(pixels []byte, data uint32) {
	toimpl()
}

type tBitmapXferSrc8 int
func (*tBitmapXferSrc8) Xfer(pixels []byte, data uint32) {
	for i := 0; i < len(pixels); i++ {
		pixels[i] = byte(data)
	}
}

func chooseBitmapXferProc(dst *Pixmap, paint *Paint, data *uint32) tBitmapXferProc {
	toimpl()
	return nil
}

func callBitmapXferProc(dst *Pixmap, rect Rect, xferProc tBitmapXferProc, xferData uint32) {
	toimpl()
}

func (draw *Draw) DrawPaint(paint *Paint) {
	if draw.rasterClip.IsEmpty() {
		return
	}

	var devRect Rect
	devRect.SetXYWH(0, 0, draw.dst.Width(), draw.dst.Height())

	if draw.rasterClip.IsBW() {
		/*  If we don't have a shader (i.e. we're just a solid color) we may
		    be faster to operate directly on the device bitmap, rather than invoking
		    a blitter. Esp. true for xfermodes, which require a colorshader to be
		    present, which is just redundant work. Since we're drawing everywhere
		    in the clip, we don't have to worry about antialiasing.
		*/
		var xferData uint32 = 0
		var xferProc = chooseBitmapXferProc(draw.dst, paint, &xferData)
		if xferProc != nil {
			_, ok := xferProc.(*tBitmapXferDst)
			if ok { // nothing to draw.
				return
			}

			var iter = NewRegionIter(draw.rasterClip.BWRgn())
			for !iter.Done() {
				callBitmapXferProc(draw.dst, iter.Rect(), xferProc, xferData)
				iter.Next()
			}

			return
		}
	}

	// normal case: use a blitter
	var chooser = newAutoBlitterChooser(draw.dst, draw.matrix, paint, false)
	ScanFillRect(devRect, draw.rasterClip, chooser.Blitter())
}

func (d *Draw) DrawRect(rect Rect, paint *Paint) {
	toimpl()
	return
}

// each of these costs 8-bytes of stack space, so don't make it too large
// must be even for lines/polygon to work.
const kMaxDevPts = 32

func (draw *Draw) DrawPoints(mode CanvasPointMode, count int, pts []Point, paint *Paint, forceUseDevice bool) {
	switch mode {
	case KCanvasPointModePoints:
		var r Rect
		// temporarily mark the paint as filling.
		var newPaint = paint.Clone()
		newPaint.SetStyle(KPaintStyleFill)
		var width = Scalar(newPaint.StrokeWidth())
		var radius = ScalarHalf(width)

		for i := 0; i < count; i++ {
			r = Rect{
				Left:   pts[i].X - radius,
				Top:    pts[i].Y - radius,
				Width:  width,
				Height: width,
			}
			if draw.Device() != nil {
				draw.Device().DrawRect(draw, r, newPaint)
			} else {
				draw.DrawRect(r, newPaint)
			}
		}
	}
}

func (draw *Draw) Device() *BaseDevice {
	return draw.device
}

type tAutoBlitterChooser struct {
}

func newAutoBlitterChooser(dst *Pixmap, matrix *Matrix, paint *Paint, drawCoverage bool) *tAutoBlitterChooser {
	toimpl()
	return nil
}

func (chooser *tAutoBlitterChooser) Blitter() Blitter {
	toimpl()
	return nil
}