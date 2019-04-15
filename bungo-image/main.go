package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
)

func main() {
	img, _ := png.Decode(os.Stdin)

	// confirm color mode
	// RGBAとは、ディスプレイ画像で色を表現するために用いられる、Red/Green/Blue + Alpha（透過度）を加えたもの。
	// RGBにアルファ値が乗算される形式: Pre-multiplied Alpha（乗算済みアルファ）
	// 乗算されていない形式: Straight Alpha (non-alpha-premultiplied)
	// image/colorパッケージでは、Straight Alphaをcolor.NRGBAModel構造体で表現している
	// Pre-multiplied Alphaを RGBA 構造体で表現している
	fmt.Println(img.ColorModel() == color.NRGBAModel)

	// get image realm (unit: px)
	fmt.Println(img.Bounds())

	// get the specified number of pixels
	fmt.Println(img.At(0, 0))
}
