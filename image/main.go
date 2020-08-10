package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	srcF, err := os.Open("google-2002-04.jpg")
	if err != nil {
		panic(err)
	}
	srcImg, _, err := image.Decode(srcF)
	if err != nil {
		panic(err)
	}

	sr := srcImg.Bounds()
	sl := sr.Dx()
	if sr.Dy() < sl {
		sl = sr.Dy()
	}
	img := image.NewRGBA(image.Rect(0, 0, sl, sl))
	draw.Draw(img, img.Bounds(), srcImg, image.Point{}, draw.Src)
	dstImage := imaging.Resize(img, 70, 70, imaging.Lanczos)
	//dstImage :=imaging.Thumbnail(srcImg, 70, 70, imaging.Lanczos)

	file, err := os.Create("img.jpg")
	if err != nil {
		panic(err)
	}
	err = jpeg.Encode(file, dstImage, &jpeg.Options{Quality: 100})
	if err != nil {
		panic(err)
	}
}
