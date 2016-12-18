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

func (clip *RasterClip) IsEmpty() bool {
	return clip.isEmpty
}

func (clip *RasterClip) IsBW() bool {
	return clip.isBW
}

func (clip *RasterClip) SetRect(rect Rect) bool {
	clip.isBW = true
	clip.aaclip.SetEmpty()
	clip.isRect = clip.bw.SetRect(rect)
	clip.isEmpty = !clip.isRect
	return clip.isRect
}

func (clip *RasterClip) BWRgn() *Region {
	toimpl()
	return nil
}

func (clip *RasterClip) Translate(x, y Scalar, otr *RasterClip) {
	toimpl()
}

func (clip *RasterClip) Op(rect Rect, op RegionOp) {
	toimpl()
}

func (clip *RasterClip) ForceGetBW() *Region {
	toimpl()
	return nil
}