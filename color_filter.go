package ggk

type ColorFilterFuncs interface {
	FilterSpan(src []PremulColor, count int, result []PremulColor)
}

type ColorFilter struct {
}

/** Construct a colorfilter whose effect is to first apply the inner filter and then apply
 *  the outer filter to the result of the inner's.
 *  The reference counts for outer and inner are incremented.
 *
 *  Due to internal limits, it is possible that this will return NULL, so the caller must
 *  always check.
 */
func NewColorFilterFromComposeFilter(outer, inner *ColorFilter) *ColorFilter {
	return &ColorFilter{}
}

func (filter *ColorFilter) FilterColor(color Color) Color {
	toimpl()
	return KColorBlack
}