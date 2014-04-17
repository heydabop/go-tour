package main

import(
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Pixel struct{
	R uint32
	G uint32
	B uint32
	A uint32
}

func (p Pixel) RGBA() (r, g, b, a uint32){
	return p.R, p.G, p.B, p.A
}

type Image struct{
	X int
	Y int
	Img [][]Pixel
}

func (i Image) ColorModel() color.Model{
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle{
	return image.Rectangle{image.Point{0, 0}, image.Point{i.X, i.Y}}
}

func (i Image) At(x, y int) color.Color{
	return i.Img [x][y]
}

func main(){
	x, y := 100, 100
	img := make([][]Pixel, x)
	for i := 0; i < x; i++ {
		img[i] = make([]Pixel, y)
	}
	m := Image{x, y, img}
	pic.ShowImage(m)
}
