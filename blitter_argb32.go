package ggk

type ARGB32Blitter struct {
}

func NewARGB32Blitter(device *Pixmap, paint *Paint) Blitter {
	toimpl()
	return nil
}

type ARGB32ShaderBlitter struct {
}

func NewARGB32ShaderBlitter(device *Pixmap, paint *Paint, shaderContext *ShaderContext) Blitter {
	toimpl()
	return nil
}

func (blitter *ARGB32ShaderBlitter)BlitRect(x, y, width, height int) {
	// assert(x >= 0 && y >= 0 && x + width <= blitter.device.Width() && y + height <= blitter.device.Height())
	// var (
	// 	device []uint32 = blitter.device.writable_addr32
	// 	deviceRB int    = blitter.device.RowBytes()
	// 	shaderContext *ShaderContext = blitter.shaderContext
	// 	span *PMColor = blitter.buffer;
	// )

	// if blitter.constInY {
	// 	if blitter.shadeDirectlyIntoDevice {
	// 		// shade the first row directly into the device.
	// 		shaderContext.shadeSpan(x, y, device, width)
	// 		span = device
	// 		for height--; height > 0; height-- {
	// 			device = device[deviceRB:]
	// 			memcpy(device, span, width << 2)
	// 		}
	// 	} else {
	// 		shaderContext.shadeSpan(x, y, span, width)
	// 		var xfer = blitter.xfermode
	// 		if xfer != nil {
				
	// 		}
	// 	}
	// }
}

type ARGB32BlackBlitter struct {
}

func NewARGB32BlackBlitter(device *Pixmap, paint *Paint) Blitter {
	toimpl()
	return nil
}

type ARGB32OpaqueBlitter struct {
}

func NewARGB32OpaqueBlitter(device *Pixmap, paint *Paint) Blitter {
	toimpl()
	return nil
}