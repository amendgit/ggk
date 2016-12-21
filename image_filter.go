package ggk

type ImageFilter struct {
}

func NewImageFilter() *ImageFilter {
	return &ImageFilter{}
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
