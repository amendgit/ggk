package ggk

type MaskFormat int

const (
	KMaskFormatBW = MaskFormat(iota)
	KMaskFormatA8
	KMaskFormat3D
	KMaskFormatARGB32
	KMaskFormatLCD16
)

type MaskFilter struct {

}

func (filter *MaskFilter) Format() MaskFormat {
	return KMaskFormatARGB32
}
