package ggk

type ColorFilterFuncs interface {
	FilterSpan(src []PremulColor, count int, result []PremulColor)
}

type ColorFilter struct {

}