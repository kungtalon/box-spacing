package boxspacing

import (
	"math"
	"sort"
)

type Rectangle struct {
	Min, Max Point
}

func (r *Rectangle) Area() float64 {
	return (r.Max.X - r.Min.X) * (r.Max.Y - r.Min.Y)
}

func (r *Rectangle) Center() Point {
	return Point{(r.Min.X + r.Max.X) / 2.0, (r.Min.Y + r.Max.Y) / 2.0}
}

func (r *Rectangle) Translate(vec Vector) Rectangle {
	return Rectangle{Min: r.Min.Add(vec), Max: r.Max.Add(vec)}
}

// Pad enlargest the rectangle while keeping the center
func Pad(rect Rectangle, padSize float64) Rectangle {
	padding := Vector{padSize, padSize}
	newMin := rect.Min.Sub(padding)
	newMax := rect.Max.Add(padding)
	return Rectangle{Min: newMin, Max: newMax}
}

func Unpad(rect Rectangle, padSize float64) Rectangle {
	padding := Vector{padSize, padSize}
	newMin := rect.Min.Add(padding)
	newMax := rect.Max.Sub(padding)
	return Rectangle{Min: newMin, Max: newMax}
}

func CountIntersection(rects []Rectangle) []int {
	results := make([]int, len(rects))
	for i := range rects {
		for j := i + 1; j < len(rects); j++ {
			if IsOverlap(rects[i], rects[j]) {
				results[i]++
				results[j]++
			}
		}
	}
	return results
}

func IsOverlap(rectA Rectangle, rectB Rectangle) bool {
	if math.Max(rectA.Min.X, rectB.Min.X) < math.Min(rectA.Max.X, rectB.Max.X) &&
		math.Max(rectA.Min.Y, rectB.Min.Y) < math.Min(rectA.Max.Y, rectB.Max.Y) {
		return true
	}
	return false
}

type RectStack struct {
	Rectangles      []Rectangle
	numIntersection []int
}

func (r *RectStack) Push(rect Rectangle) {
	r.Rectangles = append(r.Rectangles, rect)
}

func (r *RectStack) Pop() Rectangle {
	l := len(r.Rectangles)
	res := r.Rectangles[l-1]
	r.Rectangles = r.Rectangles[:l-1]
	return res
}

func (r RectStack) Len() int {
	return len(r.Rectangles)
}

func (r RectStack) Swap(i, j int) {
	r.Rectangles[i], r.Rectangles[j] = r.Rectangles[j], r.Rectangles[i]
	r.numIntersection[i], r.numIntersection[j] = r.numIntersection[j], r.numIntersection[i]
}

func (r RectStack) Less(i, j int) bool {
	return r.numIntersection[i] > r.numIntersection[j]
}

func BuildRectangleStack(rects []Rectangle) (RectStack, []Rectangle) {
	numIntxn := CountIntersection(rects)
	stack := RectStack{Rectangles: rects, numIntersection: numIntxn}
	sort.Sort(stack)

	var settledRects []Rectangle
	for {
		idx := stack.Len() - 1
		if idx < 0 {
			break
		}
		if stack.numIntersection[idx] != 0 {
			break
		}
		settledRects = append(settledRects, stack.Pop())
	}
	return stack, settledRects
}
