package ggk

import (
	"image"
)

var RectZero Rect

type Rect struct {
	Left   Scalar
	Top    Scalar
	Width  Scalar
	Height Scalar
}

func MakeRect(x, y, width, height Scalar) Rect {
	return Rect{x, y, width, height}
}

// Make rectangle from width and size, the left and top set to 0.
func MakeRectWH(width, height Scalar) Rect {
	return Rect{0, 0, width, height}
}

// Make rectangle from (left, top, right, bottom).
func MakeRectLTRB(left, top, right, bottom Scalar) Rect {
	return Rect{left, top, right - left, bottom - top}
}

func MakeRectEmpty() Rect {
	return Rect{0, 0, 0, 0}
}

// Return te left edge of the rect.
func (rect Rect) L() Scalar {
	return rect.Left
}

func (rect Rect) X() Scalar {
	return rect.Left
}

// Return the top edge of the rect.
func (rect Rect) T() Scalar {
	return rect.Top
}

func (rect Rect) Y() Scalar {
	return rect.Top
}

// Return the rectangle's width. This does not check for a valid rect
// (i.e. left <= right) so the result may be negative.
func (rect Rect) W() Scalar {
	return rect.Width
}

// Returns the rectangle's height. This does not check for a vaild rect
// (i.e. top <= bottom) so the result may be negative.
func (rect Rect) H() Scalar {
	return rect.Height
}

// Returns the rectangle's right edge.
func (rect Rect) R() Scalar {
	return rect.Left + rect.Width
}

// Returns the rectangle's bottom edge.
func (rect Rect) B() Scalar {
	return rect.Top + rect.Height
}

// Returns the rectangle's center x.
func (rect Rect) CenterX() Scalar {
	return rect.Left + rect.Width*0.5
}

// Returns the rectangle's center Y.
func (rect Rect) CenterY() Scalar {
	return rect.Top + rect.Height*0.5
}

// Return true if the rectangle's width or height are <= 0
func (rect Rect) IsEmpty() bool {
	return rect.Left <= 0 || rect.Height <= 0
}

func (rect Rect) SetEmpty() {
	rect.SetXYWH(0, 0, 0, 0)
}

// Return true if the two rectangles have same position and size.
func (a Rect) Equal(b Rect) bool {
	return a.Left == b.Left && a.Top == b.Top && a.Width == b.Width && a.Height == b.Height
}

// Set the rectangle's edges with (x, y, w, h)
func (rect *Rect) SetXYWH(x, y, width, height Scalar) {
	rect.Left, rect.Top, rect.Width, rect.Height = x, y, width, height
}

// Set the rectangle's edges with (left, top, right, bottom)
func (rect *Rect) SetLTRB(left, top, right, bottom Scalar) {
	rect.Left, rect.Top, rect.Width, rect.Height = left, top, right-left, bottom-top
}

func (rect *Rect) SetLTRBPoint(lt, rb Point) {
	rect.Left = ScalarMin(lt.X, rb.X)
	rect.Top = ScalarMin(lt.Y, rb.Y)
	rect.Width = ScalarAbs(rb.X - lt.X)
	rect.Height = ScalarAbs(rb.Y - lt.Y)
}

func (rect *Rect) IntersectXYWH(x, y, w, h Scalar) bool {
	toimpl()
	return false
}

func (rect *Rect) InsetLTRB(left, top, right, bottom Scalar) {
	rect.Left = rect.Left + left
	rect.Top = rect.Top + top
	rect.Width = rect.Width - left + right
	rect.Height = rect.Height - top + bottom
}

func (rect Rect) ToGoRect() image.Rectangle {
	return image.Rect(int(rect.Left), int(rect.Top), int(rect.Width), int(rect.Height))
}
