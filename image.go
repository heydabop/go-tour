package main

import(
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
	"time"
	"math/rand"
)

type Pixel struct{
	R uint8
	G uint8
	B uint8
	A uint8
}

func (p Pixel) RGBA() color.Color{
	return color.RGBA{p.R, p.G, p.B, p.A}
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
	return i.Img[x][y].RGBA()
}

func main(){
	rand.Seed(time.Now().Unix())

	x, y := 100, 100
	img := make([][]Pixel, x)
	for i := 0; i < x; i++ {
		img[i] = make([]Pixel, y)
		for j := 0; j < y; j++ {
			img[i][j] = Pixel{uint8(rand.Intn(255)),
				uint8(rand.Intn(255)),
				uint8(rand.Intn(255)),
				255}
		}
	}
	m := Image{x, y, img}
	pic.ShowImage(m)
}
