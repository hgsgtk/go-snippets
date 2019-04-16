package main

import (
	"github.com/hgsgtk/go-snippets/bungo-image/lerp"
	"image"
	"image/color"
)

// Resize 与えられた画像を線形補間法を使用して画像を拡大・縮小する
func Resize(img image.Image, xRatio, yRatio float64) image.Image {
	width := int(float64(img.Bounds().Size().X) * xRatio)
	height := int(float64(img.Bounds().Size().Y) * yRatio)

	newRect := image.Rect(0, 0, width, height)
	dst := image.NewRGBA64(newRect)
	bounds := dst.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			dst.Set(x, y, LerpEffect(img, xRatio, yRatio, x, y))
		}
	}
	return dst
}

// LerpEffect Lerpを行った結果の画素数を返す関数
// srcは元画像 x,yはdst
func LerpEffect(src image.Image, xRatio, yRatio float64, x, y int) color.RGBA64 {
	// 元画像の近傍４画素の座標と、alpha, beta を取得
	x1, x2, alpha := getLerpParam(x, xRatio)
	y1, y2, beta := getLerpParam(y, yRatio)

	ps := lerp.Points{
		{X: x1, Y: y1},
		{X: x2, Y: y1},
		{X: x1, Y: y2},
		{X: x2, Y: y2},
	}

	r := lerp.Lerp(initGetOneColorFunc(src, "R"), alpha, beta, ps)
	g := lerp.Lerp(initGetOneColorFunc(src, "G"), alpha, beta, ps)
	b := lerp.Lerp(initGetOneColorFunc(src, "B"), alpha, beta, ps)
	a := lerp.Lerp(initGetOneColorFunc(src, "A"), alpha, beta, ps)

	return color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)}
}

// getLerpParam １軸に対して、Lerpで使うパラメータと近傍点をdstの座標と拡大縮小比率を取得する
func getLerpParam(dstPos int, ratio float64) (int, int, float64) {
	// 拡大前の座標 = 拡大後の座標 / リサイズ比率
	v1float := float64(dstPos) / ratio

	// 拡大前の座標から最も近い2つの整数値
	v1 := int(v1float)
	v2 := v1 + 1

	// 拡大前の座標の浮動小数点数 - 拡大前の座標の整数値
	v3 := v1float - float64(v1)
	return v1, v2, v3
}

// initGetOneColorFunc src の RGBAからいずれか１つを抽出する関数fを返す関数
func initGetOneColorFunc(src image.Image, colorName string) lerp.PosDependFunc {
	return func(x, y int) float64 {
		var c uint32
		switch colorName {
		case "R":
			c, _, _, _ = src.At(x, y).RGBA()
		case "G":
			_, c, _, _ = src.At(x, y).RGBA()
		case "B":
			_, _, c, _ = src.At(x, y).RGBA()
		case "A":
			_, _, _, c = src.At(x, y).RGBA()
		}
		return float64(c)
	}
}

// 2. gray gopher
//func main() {
//	img, _ := png.Decode(os.Stdin)
//	bounds := img.Bounds()
//
//	// generate image that color model is Gray16
//	// it's a black image.
//	dst := image.NewGray16(bounds)
//
//	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
//		for x := bounds.Min.X; x < bounds.Max.X; x++ {
//			// convert original image pixel color to Gray16
//			c := color.Gray16Model.Convert(img.At(x, y))
//			gray, _ := c.(color.Gray16)
//			dst.Set(x, y, gray)
//		}
//	}
//	png.Encode(os.Stdout, dst)
//}

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
