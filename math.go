package boxspacing

import "math"

type Point struct {
	X, Y float64
}

func (p *Point) Add(vec Vector) Point {
	return Point{p.X + vec.X, p.Y + vec.Y}
}

func (p *Point) Sub(vec Vector) Point {
	return Point{p.X - vec.X, p.Y - vec.Y}
}

type Vector struct {
	X, Y float64
}

func (v *Vector) Unit() Vector {
	mod := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return v.Mul(1 / mod)
}

func (v *Vector) Mul(ratio float64) Vector {
	return Vector{v.X * ratio, v.Y * ratio}
}

func (v *Vector) Rotate(theta float64) Vector {
	cos := math.Cos(theta)
	sin := math.Sin(theta)

	return Vector{
		X: v.X*cos - v.Y*sin,
		Y: v.X*sin + v.Y*cos,
	}
}

type geometricCenter struct {
	weightedX float64
	weightedY float64
	weight    float64
}

func (g *geometricCenter) Add(rect Rectangle) {
	mass := rect.Area()
	p := rect.Center()
	g.weight += mass
	g.weightedX += mass * p.X
	g.weightedY += mass * p.Y
}

func (g *geometricCenter) Point() Point {
	return Point{
		g.weightedX / g.weight,
		g.weightedY / g.weight,
	}
}

func vector(x, y int) Vector {
	return Vector{float64(x), float64(y)}
}

func round[T float64 | float32](num T) int {
	return int(num + 0.5)
}
