package rt

import (
	"image"
	"os"
	"image/png"
	"image/color"
)

type Result struct{
	W, H int
	Path string

	image *image.NRGBA64
}

func (result *Result) Init() {
	result.image = image.NewNRGBA64(image.Rect(0, 0, result.W, result.H))
}

func (result *Result) SetPixel(x, y int, col *color.NRGBA64) {
	result.image.Set(x, y, col)
}

func (result *Result) Save() {
	f, err := os.Create(result.Path)
	if err != nil {
		panic(err)
	}

	err = png.Encode(f, result.image)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}