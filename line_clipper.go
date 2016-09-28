package ggk

const (
	kLineClipperMaxPoints = 4
	kLineClipperMaxClippedLineSegments = kLineClipperMaxPoints - 1
)

type tLineClipper struct {
	// empty.
}

// Clip the line pts[0]...pt[1] against clip, ignoring segments that lie 
// completely above or below the clip. For portions to the left or right, turn 
// those into vertical line segments that are aligned to the edge of the clip.
//
// Return the number of line segments that result, and store the end-points of 
// those segments sequentially in lines as follows:
//     1st segment: lines[0]..lines[1]
//     2nd segment: lines[1]..lines[2]
//     3rd segment: lines[2]..lines[3]
func (clipper *tLineClipper) ClipLine(pts [2]Point, clip Rect, lines [kLineClipperMaxPoints]Point, canCullToTheRight bool) int {
	toimpl()
	return 0
}

// Intersect the line segment against the rect. If there is a non-empty
// resulting segment, return true and set dst[] to that segment. If not,
// return false and ignore dst[].
//
// ClipLine is specialized for scan-conversion, as it adds vertical
// segments on the sides to show where the line extended beyond the
// left or right sides. IntersectLine does not.
func (clipper *tLineClipper) IntersectLine(src [2]Point, clip Rect, dst [2]Point) bool {
	toimpl()
	return false
}