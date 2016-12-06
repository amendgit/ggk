package ggk

var (
	gRegionRectRunHeadPtr *RegionRunHead = nil
)

// Region encapsulates the geometric region used to specify clippint areas for
// drawing.
type Region struct {
	bounds   Rect
	runHead *RegionRunHead
}

type RegionOp int

const (
	KRegionOpDifference = iota
	KRegionOpIntersect
	KRegionOpUnion
	KRegionOpXOR
	KRegionOpReverseDifference
	KRegionOpReplace
	KRegionOpLastEnum = KRegionOpReplace
)

func (r *Region) FromRegionOpRegion(rgna *Region, op RegionOp, rgnb *Region) bool {
	return false
}

func (r *Region) FromRegionOpRect(rgn *Region, op RegionOp, rect Rect) bool {
	return false
}

func (r *Region) FromRectOpRegion(rect Rect, op RegionOp, rgn *Region) bool {
	return false
}

func (rgn *Region) SetRect(r Rect) bool {
	return rgn.SetLTRB(r.L(), r.T(), r.R(), r.B())
}

func (rgn *Region) SetLTRB(l, t, r, b Scalar) bool {
	if l >= r || t >= b {
		return rgn.SetEmpty()
	}
	rgn.FreeRuns()
	rgn.bounds.SetLTRB(l, t, r, b)
	rgn.runHead = gRegionRectRunHeadPtr;
	return true
}

func (region *Region) SetEmpty() bool {
	toimpl()
	return false
}

func (region *Region) FreeRuns() {
	toimpl()
}

type RegionIterFunc func(rect Rect, skip *int, stop *bool)

func (rgn *Region) Iter(iterFunc RegionIterFunc) {
	if iterFunc == nil {
		return
	}

	var (
		skip int
		stop bool
		idx  int
		rect Rect
		ltrb []Scalar = rgn.runHead.CompactLTRBs()
	)

	var l, r, t, b = ltrb[3], ltrb[4], ltrb[0], ltrb[1]
	idx += 5
	for {
		rect = MakeRectLTRB(l, t, r, b)

		if skip != 0 {
			skip--
		}

		if skip == 0 {
			iterFunc(rect, &skip, &stop)
		}

		if stop {
			break
		}

		if ltrb[idx] < KRegionRunHeadLRTBSentinel {
			// valid X value
			l, r = ltrb[idx], ltrb[idx+1]
			idx += 2
		} else {
			// we're at the end of a line
			idx += 1
			if ltrb[idx] < KRegionRunHeadLRTBSentinel {
				// valid Y value
				var intervals = ltrb[idx+1]

				if intervals == 0 {
					// empty line
					t = ltrb[idx]
					idx += 3
				} else {
					t = b
				}

				b = ltrb[idx]
				l = ltrb[idx+2]
				r = ltrb[idx+3]
				idx += 4
			} else {
				break
			}
		}
	}
}

type RegionClipFunc func(rect Rect, skip *int, stop *bool)

func (r *Region) Clip(clip Rect, clipFunc RegionClipFunc) {
	return
}

type RegionSpanFunc func(left, right *int)

func (r *Region) Span(y, left, right int, spanFunc RegionSpanFunc) {
	return
}

const KRegionRunHeadLRTBSentinel = KScalarMax

type RegionRunHead struct {
	RefCount  int32
	GridCount int32

	yspanCount    int32
	intervalCount int32

	// e.g. T B N L R L R L R S T 0 S B N L R S S
	compactLTRBs []Scalar
}

func (h *RegionRunHead) YSpanCount() int32 {
	return h.yspanCount
}

func (h *RegionRunHead) IntervalCount() int32 {
	return h.intervalCount
}

func (h *RegionRunHead) CompactLTRBs() []Scalar {
	return h.compactLTRBs
}

type RegionIter struct {

}

func NewRegionIter(region *Region) *RegionIter {
	toimpl()
	return nil
}

func (iter *RegionIter) Next() bool {
	toimpl()
	return false
}

func (iter *RegionIter) Done() bool {
	toimpl()
	return false
}

func (iter *RegionIter) Rect() Rect {
	toimpl()
	return RectZero
}