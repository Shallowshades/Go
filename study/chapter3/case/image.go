package main

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

/*
图像

	image包定义了Image接口

	type Image interface {
	    ColorModel() color.Model
	    Bounds() Rectangle
	    At(x, y int) color.Color
	}
*/

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 200)
}
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x * y), uint8(x * y), 255, 255}
}

func main() {

	{
		m := image.NewRGBA(image.Rect(0, 0, 100, 100))
		fmt.Println(m.Bounds())
		fmt.Println(m.At(0, 0).RGBA())
	}

	{
		m := Image{}
		pic.ShowImage(m)
	}
}
