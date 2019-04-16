# 文Go image package についての深掘り学習

## image packageについて
- https://golang.org/pkg/image/
- Package image implements a basic 2-D image library.
- Registration is typically automatic as a side effect of initializing that format's package so that, to decode a PNG image
    - `_` にてインポートしておく必要がある

### blank import / period import
https://qiita.com/shiena/items/c1ac3192af3b00f413ac
- blank importはインポートによる副作用（初期化）のためだけにパッケージをインポートするとき使う。
- 具体的にはInit関数を使うとき

## image.Image Interfaceについて

```go
// Image is a finite rectangular grid of color.Color values taken from a color
// model.
type Image interface {
	// ColorModel returns the Image's color model.
	ColorModel() color.Model
	// Bounds returns the domain for which At can return non-zero color.
	// The bounds do not necessarily contain the point (0, 0).
	Bounds() Rectangle
	// At returns the color of the pixel at (x, y).
	// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
	// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
	At(x, y int) color.Color
}
```

- color model と 書籍内での解説に用いられるカラーモードの違い
    - イコールなんだろうな、サクッと調べたけど color modeで引っかからなかった

## blank importの必要性

- image/png packageでDecodeで用いる画像フォーマットを登録している
- https://github.com/golang/go/blob/master/src/image/png/reader.go#L1030
- なにかimageパッケージで扱いたいものが現れたらここにフォーマット登録する

```go
func init() {
	image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)
}
```

## 線形補間法（lerp）
計算量が非常に少ないことが特徴のコンピュータグラフィックスを含む多くの分野で使われている補間