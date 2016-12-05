package ggk

type tBitmapXferProc interface {
	Xfer(pixels []byte, data uint32)
}

type tBitmapXferClear int

func (*tBitmapXferClear)Xfer(pixels []byte, data uint32) {
	for i := 0; i < len(pixels); i++ {
		pixels[i] = 0
	}
}

type tBitmapXferDst int

func (*tBitmapXferDst)Xfer(pixels []byte, data uint32) {
	toimpl()
}

type tBitmapXferSrc32 int

func (*tBitmapXferSrc32) Xfer(pixels []byte, data uint32) {
	toimpl()
}

type tBitmapXferSrc16 int

func (*tBitmapXferSrc16) Xfer(pixels []byte, data uint32) {
	toimpl()
}

type tBitmapXferSrc8 int

func (*tBitmapXferSrc8) Xfer(pixels []byte, data uint32) {
	for i := 0; i < len(pixels); i++ {
		pixels[i] = byte(data)
	}
}
