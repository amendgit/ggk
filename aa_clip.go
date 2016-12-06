package ggk

// AAClip is anti-alising clip.
type AAClip struct {
	bounds  Rect
	runHead *AAClipRunHead
}

func NewAAClip() *AAClip {
	toimpl()
	var clip = &AAClip{
		bounds: RectZero,
		runHead: nil,
	}
	return clip
}

func (clip *AAClip) Assign(otr *AAClip) *AAClip {
	toimpl()
	return nil
}

func (clip *AAClip) Equal(otr *AAClip) bool {
	toimpl()
	return false
}

func (clip *AAClip) Swap(otr *AAClip) bool {
	toimpl()
	return false
}

func (clip *AAClip) IsEmpty() bool {
	toimpl()
	return false
}

func (clip *AAClip) Bounds() Rect {
	toimpl()
	return clip.bounds
}

// Returns true iff the clip is not empty, and is just a hard-edged rect (no partial alpha).
// If true, getBounds() can be used in place of this clip.
func (clip *AAClip) IsRect() bool {
	toimpl()
	return false
}

func (clip *AAClip) SetEmpty() bool {
	clip.FreeRuns()
	clip.bounds.SetEmpty()
	clip.runHead = nil
	return false
}

func (clip *AAClip) SetRect(rect Rect, doAA bool) bool {
	toimpl()
	return false
}

func (clip *AAClip) SetPath(path *Path, region *Region, doAA bool) bool {
	toimpl()
	return false
}

func (clip *AAClip) SetRegion(region *Region) bool {
	toimpl()
	return false
}

func (clip *AAClip) Op(a *AAClip, b *AAClip, op RegionOp) bool {
	toimpl()
	return false
}

// func (clip *AAClip) Op(rect Rect, op RegionOp, doAA bool) bool {
// 	toimpl()
// 	return false
// }

// func (clip *AAClip) Op(a *AAClip, op RegionOp) bool {
// 	toimpl()
// 	return false
// }

func (clip *AAClip) TranslateTo(dx, dy int, dst *AAClip) bool {
	toimpl()
	return false
}

func (clip *AAClip) Translate(dx, dy int) bool {
	toimpl()
	return false
}

// Allocates a mask the size of the aaclip, and expands its data into
// the mask, using kA8_Format
func (clip *AAClip) CopyToMask(mask *Mask) {
	toimpl()
}

// called internally
func (clip *AAClip) QuickContains(left, top, right, bottom int) bool {
	toimpl()
	return false
}

func (clip *AAClip) QuickContainsRect(rect Rect) bool {
	toimpl()
	return false
}

func (clip *AAClip) FindRow(y int, lastYForRow *int) *uint8 {
	toimpl()
	return nil
}

func (clip *AAClip) FindX(data []uint8, x int, initialCount *int) *uint8 {
	toimpl()
	return nil
}

func (clip *AAClip) Validate() {
	toimpl()
}

func (clip *AAClip) Debug(compressY bool) {
	toimpl()
}

func (clip *AAClip) FreeRuns() {
	clip.runHead = nil
}

const kAAClipBlitterGrayMaskScratchSize = 32 * 32

type AAClipBlitter struct {
	Blitter

	blitter      *Blitter
	aaclip       *AAClip
	aaclipBounds Rect
	// point into scanlineScratch
	runs *int16
	aa   *Alpha
	grayMaskScratch []uint8
}

func (blitter *AAClipBlitter) BlitH(x, y, width int) {
	toimpl()
	return
}

func (blitter *AAClipBlitter) BlitAntiH(x, y int, alphas []Alpha, runs []int16) {
	toimpl()
	return 
}

func (blitter *AAClipBlitter) BlitV(x, y ,height int, alpha Alpha) {
	toimpl()
	return
}

func (blitter *AAClipBlitter) BlitRect(x, y, width, height int) {
	toimpl()
	return
}

func (blitter *AAClipBlitter) BlitMask(mask *Mask, clip Rect) {
	toimpl()
	return
}

func (blitter *AAClipBlitter) JustAnOpaqueColor(value *uint32) *Pixmap {
	toimpl()
	return nil
}

type AAClipRunHead struct {

}