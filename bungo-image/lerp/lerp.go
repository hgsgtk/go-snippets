package lerp

type Point struct {
	X int
	Y int
}
type Points [4]Point

type PosDependFunc func(x, y int) float64

// Lerp 線形補間法の実装
// a, b 座標間の距離
func Lerp(f PosDependFunc, a, b float64, ps Points) float64 {
	return (1.0-b)*(1.0-b)*f(ps[0].X, ps[0].Y) +
		a*(1.0-b)*f(ps[1].X, ps[1].Y) +
		b*(1-a)*f(ps[0].X, ps[1].Y) +
		a*b*f(ps[1].X, ps[1].Y)
}
