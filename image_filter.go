package ggk

type ImageFilter struct {
}

func (imageFilter *ImageFilter) CanComputeFastBounds() bool {
	return false
}

type ImageFilterContext struct {
}
