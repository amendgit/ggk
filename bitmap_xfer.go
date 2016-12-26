package ggk

type tBitmapXferProc interface {
	Xfer(pixels []byte, data uint32)
}

type tBitmapXferProcClear int
func (*tBitmapXferProcClear)Xfer(pixels []byte, data uint32) {
	for i := 0; i < len(pixels); i++ {
		pixels[i] = 0
	}
}

type tBitmapXferProcDst int
func (*tBitmapXferProcDst)Xfer(pixels []byte, data uint32) {
	toimpl()
}

type tBitmapXferProcSrcD32 int
func (*tBitmapXferProcSrcD32) Xfer(pixels []byte, data uint32) {
	toimpl()
}

type tBitmapXferProcSrcD16 int
func (*tBitmapXferProcSrcD16) Xfer(pixels []byte, data uint32) {
	toimpl()
}

type tBitmapXferProcSrcD8 int
func (*tBitmapXferProcSrcD8) Xfer(pixels []byte, data uint32) {
	for i := 0; i < len(pixels); i++ {
		pixels[i] = byte(data)
	}
}

type tBitmapXferProcSrcDA8 int
func (*tBitmapXferProcSrcDA8) Xfer(pixels []byte, data uint32) {
	toimpl()
}
