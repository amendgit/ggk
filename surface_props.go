package ggk

// SurfaceProps describe of how the LCD strips are arranged for each pixel. If
// this is unknown, or the pixels are meant to be "protable" and/or transformed
// before showing (e.g. rotated, scaled) then use KPixelsGeometryUnknown.
type PixelGeometry int

const (
	KPixelGeometryUnknown = iota
	KPixelGeometryRGBH
	KPixelGeometryBGRH
	KPixelGeometryRGBV
	KPixelGeometryBGRV
)

type SurfacePropsFlags int

const (
	KSurfacePropsFlagNone              = 0
	KSurfacePropsFlagDisallowAntiAlias = 1 << iota
	KSurfacePropsFlagDisallowDither
	KSurfacePropsFlagUseDistanceFieldFonts
)

type SurfacePropsInitType int

const (
	KSurfacePropsInitTypeNone = iota
	KSurfacePropsInitTypeLegacyFontHost
)

type SurfacePropsContentChangeMode int

const (
	KSurfaceContentChangeModeRetain = SurfacePropsContentChangeMode(iota)
	KSurfaceContentChangeModeDiscard
)

type SurfaceProps struct {
	flags         SurfacePropsFlags
	pixelGeometry PixelGeometry
}

func NewSurfaceProps(flags SurfacePropsFlags, initType SurfacePropsInitType) *SurfaceProps {
	var props = &SurfaceProps{}
	props.flags = flags
	if initType == KSurfacePropsInitTypeLegacyFontHost {
		props.pixelGeometry = computeDefaultGeometry()
	}
	return props
}

func (props *SurfaceProps) OutstandingImageSnapshot() *BaseSurface {
	toimpl()
	return nil
}

func (props *SurfaceProps) AboutToDraw(mode SurfacePropsContentChangeMode) {
	toimpl()
}

func computeDefaultGeometry() PixelGeometry {
	var order = FontLCDConfigSubpixelOrientation()
	if order == KLCDOrderNone {
		return KPixelGeometryUnknown
	} else {
		// Bit0 is RGB(0), BGR(1)
		// Bit1 is H(0), V(1)
		var gGeo = []PixelGeometry{
			KPixelGeometryRGBH,
			KPixelGeometryBGRH,
			KPixelGeometryRGBV,
			KPixelGeometryBGRV,
		}
		var index = 0
		if order == KLCDOrderBGR {
			index = index | 1
		}
		if FontLCDConfigSubpixelOrientation() == KLCDOrientationVertical {
			index = index | 2
		}
		return gGeo[index]
	}
	return KPixelGeometryBGRH
}