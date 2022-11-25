package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{width, height int}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0,0, i.width, i.height)
}
func (i Image) At(x, y int) color.Color {
	r:= uint8(x^y);
	g:= uint8(x*y);
	b:= uint8((x+y)/2);
	return color.RGBA{r,g,b, 255}
}


func main() {
	m := Image{ 256, 126}
	pic.ShowImage(m)
}
