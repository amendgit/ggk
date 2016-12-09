package ggk

type ImageFilter struct {
}

func (imageFilter *ImageFilter) CanComputeFastBounds() bool {
	toimpl()
	return false
}

func (ImageFilter *ImageFilter) AsAColorFilter() (*ColorFilter, bool) {
	toimpl()
	return new(ColorFilter), false
}

type ImageFilterContext struct {
}
