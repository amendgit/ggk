package ggk

func ScanFillRect(rect Rect, rasterClip *RasterClip, blitter Blitter) {
	if rasterClip.IsEmpty() || rect.IsEmpty() {
		return
	}

	if rasterClip.IsBW() {
		ScanFillRectRegion(rect, rasterClip.BWRgn(), blitter)
		return
	}

	var wrapper = NewAAClipBlitterWrapper(rasterClip, blitter)
	ScanFillRectRegion(rect, wrapper.Rgn(), wrapper.Blitter())
}

func ScanFillRectRegion(rect Rect, rgn *Region, blitter Blitter) {
	toimpl()
}