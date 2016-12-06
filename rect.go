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

// Return te left edge of the rect.
func (r Rect) L() Scalar {
	return r.Left
}

func (r Rect) X() Scalar {
	return r.Left
}

// Return the top edge of the rect.
func (r Rect) T() Scalar {
	return r.Top
}

func (r Rect) Y() Scalar {
	return r.Top
}

// Return the rectangle's width. This does not check for a valid rect
// (i.e. left <= right) so the result may be negative.
func (r Rect) W() Scalar {
	return r.Width
}

// Returns the rectangle's height. This does not check for a vaild rect
// (i.e. top <= bottom) so the result may be negative.
func (r Rect) H() Scalar {
	return r.Height
}

// Returns the rectangle's right edge.
func (r Rect) R() Scalar {
	return r.Left + r.Width
}

// Returns the rectangle's bottom edge.
func (r Rect) B() Scalar {
	return r.Top + r.Height
}

// Returns the rectangle's center x.
func (r Rect) CenterX() Scalar {
	return r.Left + r.Width*0.5
}

// Returns the rectangle's center Y.
func (r Rect) CenterY() Scalar {
	return r.Top + r.Height*0.5
}

// Return true if the rectangle's width or height are <= 0
func (r Rect) IsEmpty() bool {
	return r.Left <= 0 || r.Height <= 0
}

func (r Rect) SetEmpty() {
	toimpl()
}

// Return true if the two rectangles have same position and size.
func (a Rect) Equal(b Rect) bool {
	return a.Left == b.Left && a.Top == b.Top && a.Width == b.Width && a.Height == b.Height
}

// Set the rectangle's edges with (x, y, w, h)
func (r *Rect) SetXYWH(x, y, width, height Scalar) {
	r.Left, r.Top, r.Width, r.Height = x, y, width, height
}

// Set the rectangle's edges with (left, top, right, bottom)
func (r *Rect) SetLTRB(left, top, right, bottom Scalar) {
	r.Left, r.Top, r.Width, r.Height = left, top, right-left, bottom-top
}

func (r *Rect) SetLTRBPoint(lt, rb Point) {
	r.Left = ScalarMin(lt.X, rb.X)
	r.Top = ScalarMin(lt.Y, rb.Y)
	r.Width = ScalarAbs(rb.X - lt.X)
	r.Height = ScalarAbs(rb.Y - lt.Y)
}

func (r *Rect) IntersectXYWH(x, y, w, h Scalar) bool {
	toimpl()
	return false
}

func (r Rect) ToGoRect() image.Rectangle {
	return image.Rect(int(r.Left), int(r.Top), int(r.Width), int(r.Height))
}
