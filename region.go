package ggk

var (
	gRegionRectRunHeadPtr  *RegionRunHead = nil
	gRegionEmptyRunHeadPtr *RegionRunHead = nil
)

// Region encapsulates the geometric region used to specify clippint areas for
// drawing.
type Region struct {
	bounds  Rect
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

func NewRegion() *Region {
	var rgn = &Region{
		bounds:  RectZero,
		runHead: nil,
	}
	return rgn
}

func (rgn *Region) Bounds() Rect {
	return rgn.bounds
}

func (rgn *Region) FromRegionOpRegion(rgna *Region, op RegionOp, otr *Region) bool {
	toimpl()
	return false
}

func (rgn *Region) FromRegionOpRect(r *Region, op RegionOp, rect Rect) bool {
	toimpl()
	return false
}

func (rgn *Region) FromRectOpRegion(rect Rect, op RegionOp, otr *Region) bool {
	toimpl()
	return false
}

/** Return true if this region is empty */
func (rgn *Region) IsEmpty() bool {
	toimpl()
	return false
}

/** Return true if this region is a single, non-empty rectangle */
func (rgn *Region) IsRect() bool {
	toimpl()
	return false
}

/** Return true if this region consists of more than 1 rectangular area */
func (rgn *Region) IsComplex() bool {
	return !rgn.IsEmpty() && !rgn.IsRect()
}

/**
 *  If rect is non-empty, set this region to that rectangle and return true,
 *  otherwise set this region to empty and return false.
 */
func (rgn *Region) SetRect(r Rect) bool {
	return rgn.SetLTRB(r.L(), r.T(), r.R(), r.B())
}

/**
 *  If left < right and top < bottom, set this region to that rectangle and
 *  return true, otherwise set this region to empty and return false.
 */
func (rgn *Region) SetLTRB(l, t, r, b Scalar) bool {
	if l >= r || t >= b {
		return rgn.SetEmpty()
	}
	rgn.FreeRuns()
	rgn.bounds.SetLTRB(l, t, r, b)
	rgn.runHead = gRegionRectRunHeadPtr
	return true
}

/**
 *  Set the region to be empty, and return false, since the resulting
 *  region is empty
 */
func (rgn *Region) SetEmpty() bool {
	rgn.FreeRuns()
	rgn.bounds.SetEmpty()
	rgn.runHead = gRegionEmptyRunHeadPtr
	return false
}

func (rgn *Region) FreeRuns() {
	if !rgn.IsComplex() {
		rgn.runHead = nil
	}
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

func (rgn *Region) Clip(clip Rect, clipFunc RegionClipFunc) {
	toimpl()
	return
}

type RegionSpanFunc func(left, right *int)

func (rgn *Region) Span(y, left, right int, spanFunc RegionSpanFunc) {
	toimpl()
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

func (runHead *RegionRunHead) YSpanCount() int32 {
	return runHead.yspanCount
}

func (runHead *RegionRunHead) IntervalCount() int32 {
	return runHead.intervalCount
}

func (runHead *RegionRunHead) CompactLTRBs() []Scalar {
	return runHead.compactLTRBs
}

type RunType int32

const (
	KRunTypeSentinel RunType = 0x7FFFFFFF
)

/**
 *  Returns the sequence of rectangles, sorted in Y and X, that make up
 *  this region.
 */
type RegionIterator struct {
	rgn  *Region
	runs *RunType
	rect Rect
	done bool
}

func NewRegionIterator(rgn *Region) *RegionIterator {
	toimpl()
	return nil
}

func (iter *RegionIterator) Next() bool {
	toimpl()
	return false
}

func (iter *RegionIterator) Done() bool {
	toimpl()
	return false
}

func (iter *RegionIterator) Rect() Rect {
	toimpl()
	return RectZero
}
