package ggk

/**
 *  Base-class for objects that draw into SkCanvas.
 *
 *  The object has a generation ID, which is guaranteed to be unique across all drawables. To
 *  allow for clients of the drawable that may want to cache the results, the drawable must
 *  change its generation ID whenever its internal state changes such that it will draw differently.
 */
type Drawable struct {

}
