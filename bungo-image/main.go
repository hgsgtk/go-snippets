package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	img, _ := png.Decode(os.Stdin)
	bounds := img.Bounds()

	// generate image that color model is Gray16
	// it's a black image.
	dst := image.NewGray16(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// convert original image pixel color to Gray16
			c := color.Gray16Model.Convert(img.At(x, y))
			gray, _ := c.(color.Gray16)
			dst.Set(x, y, gray)
		}
	}
	png.Encode(os.Stdout, dst)
}

// 1. basic of image package
//func main() {
//	img, _ := png.Decode(os.Stdin)
//
//	// confirm color mode
//	// RGBAとは、ディスプレイ画像で色を表現するために用いられる、Red/Green/Blue + Alpha（透過度）を加えたもの。
//	// RGBにアルファ値が乗算される形式: Pre-multiplied Alpha（乗算済みアルファ）
//	// 乗算されていない形式: Straight Alpha (non-alpha-premultiplied)
//	// image/colorパッケージでは、Straight Alphaをcolor.NRGBAModel構造体で表現している
//	// Pre-multiplied Alphaを RGBA 構造体で表現している
//	fmt.Println(img.ColorModel() == color.NRGBAModel)
//
//	// get image realm (unit: px)
//	fmt.Println(img.Bounds())
//
//	// get the specified number of pixels
//	fmt.Println(img.At(0, 0))
//}
