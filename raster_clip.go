package ggk

type RasterClip struct {
	bw                     *Region
	isBW                   bool
	forceConservativeRects bool
	aaclip                 *AAClip

	isEmpty bool
	isRect  bool
}

func NewRasterClipClone(otr *RasterClip) *RasterClip {
	toimpl()
	return &RasterClip{}
}

func NewRasterClip(forceConservativeRects bool) *RasterClip {
	var clip = &RasterClip{
		forceConservativeRects: forceConservativeRects,
		bw: NewRegion(),
		isBW:    true,
		isEmpty: true,
		isRect:  false,
		aaclip: NewAAClip(),
	}
	return clip
}

func (r *RasterClip) IsEmpty() bool {
	return r.isEmpty
}

func (r *RasterClip) IsBW() bool {
	return r.isBW
}

func (clip *RasterClip) SetRect(rect Rect) bool {
	clip.isBW = true
	clip.aaclip.SetEmpty()
	clip.isRect = clip.bw.SetRect(rect)
	clip.isEmpty = !clip.isRect
	return clip.isRect
}

func (r *RasterClip) BWRgn() *Region {
	toimpl()
	return nil
}
