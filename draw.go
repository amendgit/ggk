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

func chooseBitmapXferProc(dst *Pixmap, paint *Paint, data *uint32) tBitmapXferProc {
	// todo: we can apply colorfilter up front if no shader, so we wouldn't
	// need to abort this fastpath
	if paint.Shader() != nil || paint.ColorFilter() != nil {
		return nil
	}

	var mode XfermodeMode
	var ok bool
	if mode, ok = XfermodeAsMode(paint.Xfermode()); !ok {
		return nil
	}

	var color = paint.Color()

	// collaps modes based on color...
	if mode == KXfermodeModeSrcOver {
		var alpha = color.Alpha()
		if alpha == 0 {
			mode = KXfermodeModeDst
		} else {
			mode = KXfermodeModeSrc
		}
	}

	switch mode {
	case KXfermodeModeClear:
		return new(tBitmapXferProcClear)
	case KXfermodeModeDst:
		return new(tBitmapXferProcDst)
	case KXfermodeModeSrc:
		// Should I worry about dithering for the lower depths.
		var pmc, _ = PremultiplyColor(color)
		switch dst.ColorType() {
		case KColorTypeN32:
			if (data != nil) {
				*data = uint32(pmc)
			}
			return new(tBitmapXferProcSrcD8)
		case KColorTypeRGB565:
			if (data != nil) {
				*data = Pixel32ToPixel16(uint32(pmc))
			}
			return new(tBitmapXferProcSrcD16)
		case KColorTypeAlpha8:
			if (data != nil) {
				*data = GetPackedA32(uint32(pmc))
			}
			return new(tBitmapXferProcSrcDA8)
		}
	}

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
		/* If we don't have a shader (i.e. we're just a solid color) we may
	    be faster to operate directly on the device bitmap, rather than invoking
	    a blitter. Esp. true for xfermodes, which require a colorshader to be
	    present, which is just redundant work. Since we're drawing everywhere
	    in the clip, we don't have to worry about antialiasing. */
		var xferData uint32 = 0
		var xferProc = chooseBitmapXferProc(draw.dst, paint, &xferData)
		if xferProc != nil {
			if _, ok := xferProc.(*tBitmapXferProcDst); ok { // < nothing to draw.
				return
			}

			var it = NewRegionIterator(draw.rasterClip.BWRgn())
			for !it.Done() {
				callBitmapXferProc(draw.dst, it.Rect(), xferProc, xferData)
				it.Next()
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
	blitter Blitter
}

func newAutoBlitterChooser(dst *Pixmap, matrix *Matrix, paint *Paint, drawCoverage bool) *tAutoBlitterChooser {
	var chooser = &tAutoBlitterChooser{
		blitter : BlitterChoose(dst, matrix, paint, drawCoverage),
	}
	return chooser
}

func (chooser *tAutoBlitterChooser) Blitter() Blitter {
	toimpl()
	return nil
}