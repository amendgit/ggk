package ggk

type XfermodeMode int

const (
	KXfermodeModeSrcOver XfermodeMode = iota //!< [Sa + Da * (1 - Sa), Sc + Dc * (1 - Sa)]
	KXfermodeModeDst
	KXfermodeModeSrc
	KXfermodeModeClear

)

// Xfermode
//
// Xfermode is the base class for objects that are called to implement custom
// "transfer-modes" in the drawing pipeline. The static function Create(Modes)
// can be called to return an instance of any of the predefined subclasses as
// specified in the Modes enum. When an Xfermode is assigned to an SkPaint,
// then objects drawn with that paint have the xfermode applied.
//
// All subclasses are required to be reentrant-safe : it must be legal to share
// the same instance between several threads.
type Xfermode struct {
	// empty
}

func NewXfermode() *Xfermode {
	toimpl()
	return &Xfermode{}
}

func NewXfermodeWithMode(mode XfermodeMode) *Xfermode {
	toimpl()
	return nil
}

func XfermodeIsMode(xfer *Xfermode, mode XfermodeMode) bool {
	toimpl()
	return false
}

func XfermodeAsMode(xfer *Xfermode) (XfermodeMode, bool) {
	toimpl()
	return KXfermodeModeSrcOver, false
}

func (xfermode *Xfermode) AppendStages(pipeline *RasterPipeline) bool {
	toimpl()
	return false
}