package ggk

type XfermodeInterpretation int

const (
	KXfermodeInterpretationNormal      = XfermodeInterpretation(iota) //< draw normally
	KXfermodeInterpretationSrcOver                                    //< draw as if in srcover mode
	KXfermodeInterpretationSkipDrawing                                //< draw nothing
)

func InterpretXfermode(paint *Paint, dstIsOpaque bool) XfermodeInterpretation {
	toimpl()
	return KXfermodeInterpretationNormal
}
