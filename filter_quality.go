package ggk

/**
 *  Controls how much filtering to be done when scaling/transforming complex colors
 *  e.g. images
 */
type FilterQuality int

const (
	KFilterQualityNone = FilterQuality(iota)//< fastest but lowest quality, typically nearest-neighbor
)
