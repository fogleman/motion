package motion

import "math"

type Point struct {
	X, Y float64
}

func (a Point) add(b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}

func (a Point) sub(b Point) Point {
	return Point{a.X - b.X, a.Y - b.Y}
}

func (a Point) mulScalar(b float64) Point {
	return Point{a.X * b, a.Y * b}
}

func (a Point) dot(b Point) float64 {
	return a.X*b.X + a.Y*b.Y
}

func (a Point) length() float64 {
	return math.Hypot(a.X, a.Y)
}

func (a Point) distance(b Point) float64 {
	return math.Hypot(a.X-b.X, a.Y-b.Y)
}

func (a Point) normalize() Point {
	d := a.length()
	if d == 0 {
		return Point{}
	}
	return a.mulScalar(1 / d)
}

func (a Point) lerps(b Point, s float64) Point {
	v := b.sub(a).normalize()
	return a.add(v.mulScalar(s))
}
