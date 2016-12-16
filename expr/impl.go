package main

import (
	"fmt"
)

type CanvasImpl interface {
	OnDrawPoints()
	OnDrawRect()
}

type Canvas struct {
	Impl CanvasImpl
}

func NewCanvas() *Canvas {
	var canvas = &Canvas{}
	
	canvas.Impl = canvas
	
	return canvas
}

func (canvas *Canvas) OnDrawPoints() {
	fmt.Printf("Canvas OnDrawPoints\n")
}

func (canvas *Canvas) OnDrawRect() {
	fmt.Printf("Canvas OnDrawRect\n")
}

type LuaCanvas struct {
	*Canvas 
	name string
}

func NewLuaCanvas() *LuaCanvas {
	var canvas = &LuaCanvas {
		Canvas: NewCanvas(),
		name: "LuaCanvas",
	}
	
	canvas.Impl = canvas

	return canvas
}

func (canvas *LuaCanvas) OnDrawRect() {
	fmt.Printf("LuaCanvas %v OnDrawRect name %v\n", canvas, canvas.name)
}

type PDFCanvas struct {
	*Canvas 
	name string
}

func NewPDFCanvas() *PDFCanvas {
	var canvas = &PDFCanvas {
		Canvas: NewCanvas(),
		name: "PDFCanvas",
	}
	
	canvas.Impl = canvas
	
	return canvas
}

func (canvas *PDFCanvas) OnDrawRect() {
	fmt.Printf("PDFCanvas %v OnDrawRect name %v\n", canvas, canvas.name)
}

func main() {
	var canvas *Canvas

	fmt.Printf("PDFCanvas \n")
	var pdfCanvas = NewPDFCanvas()
    
    pdfCanvas.OnDrawPoints()
    pdfCanvas.OnDrawRect()

    pdfCanvas.Impl.OnDrawPoints()
    pdfCanvas.Impl.OnDrawRect()

    fmt.Printf("\n")
    fmt.Printf("Cast PDFCanvas %v to Canvas %v\n", pdfCanvas, canvas)

    canvas = pdfCanvas.Canvas
    canvas.OnDrawPoints()
    canvas.OnDrawRect()
    canvas.Impl.OnDrawPoints()
    canvas.Impl.OnDrawRect()

    fmt.Printf("\n")
    fmt.Printf("LuaCanvas\n")

	var luaCanvas = NewLuaCanvas()
	luaCanvas.OnDrawPoints()
	luaCanvas.OnDrawRect()
	luaCanvas.Impl.OnDrawPoints()
	luaCanvas.Impl.OnDrawRect()

	fmt.Printf("\n")
    fmt.Printf("Cast LuaCanvas %v to Canvas %v\n", luaCanvas, canvas)
	
	canvas = luaCanvas.Canvas
    canvas.OnDrawPoints()
    canvas.OnDrawRect()
    canvas.Impl.OnDrawPoints()
    canvas.Impl.OnDrawRect()
}