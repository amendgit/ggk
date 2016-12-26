package ggk

// Pixmap pairs ImageInfo with actual pixels and rowbytes. This class does not
// try to manage the lifetime of the pixel memory (nor the colortable if
// provided).
type Pixmap struct {
	imageInfo  *ImageInfo
	colorTable *ColorTable
	pixels     []byte
	rowBytes   int
}

func NewPixmap() *Pixmap {
	var pixmap = &Pixmap{}
	toimpl()
	return pixmap
}

func (pixmap *Pixmap) Width() Scalar {
	return pixmap.imageInfo.Width()
}

func (pixmap *Pixmap) Height() Scalar {
	return pixmap.imageInfo.Height()
}

func (pixmap *Pixmap) Reset(imageInfo *ImageInfo, pixelbytes []byte, rowBytes int, colorTable *ColorTable) {
	pixmap.imageInfo = imageInfo
	pixmap.pixels = pixelbytes
	pixmap.rowBytes = rowBytes
	pixmap.colorTable = colorTable
}

func (pixmap *Pixmap) ColorType() ColorType {
	toimpl()
	return KColorTypeN32
}

type AutoPixmapUnlock struct {
}
