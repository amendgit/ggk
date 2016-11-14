package ggk

type Draw struct {
	dst        *Pixmap
	mat        *Matrix
	region     *Region
	rasterClip *RasterClip
	clipStack  *ClipStack
	device     *BaseDevice
	// procs      *DrawProcs
}

func (d *Draw) DrawPaint(paint *Paint) {
	if d.rasterClip.IsEmpty() {
		return
	}

	// var devRect Rect = MakeRect(0, 0, d.dst.Width(), d.dst.Height())

	if d.rasterClip.IsBW() {
		// TOIMPL
	}

	// normal case: use a blitter
	// ScanFillRect(devRect, p.rasterClip, blitter)
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
