package boxspacing

import (
	"fmt"
	"math"

	"github.com/samber/lo"
)

var config = struct {
	indexStrategy string
	paddingSize   float64
	initialStep   float64
	ratioInterval float64
	thetas        []float64
}{
	indexStrategy: "heuristic",
	paddingSize:   0.1,
	initialStep:   0.1,
	ratioInterval: 0.2,
	thetas:        []float64{0, math.Pi / 4, -math.Pi / 4},
}

type BoxIndexer interface {
	Insert(rect Rectangle)
	Check(rect Rectangle) bool
}

func NewBoxIndexer(algo string, boxes []Rectangle) BoxIndexer {
	switch algo {
	case "heuristic":
		return &HeuristicBoxIndex{boxes}
	}
	return nil
}

type HeuristicBoxIndex struct {
	boxes []Rectangle
}

func (bi *HeuristicBoxIndex) Insert(rect Rectangle) {
	bi.boxes = append(bi.boxes, rect)
}

func (bi *HeuristicBoxIndex) Check(rect Rectangle) bool {
	for _, box := range bi.boxes {
		if IsOverlap(box, rect) {
			return false
		}
	}
	return true
}

func Process(boxes []Rectangle, takenAreas []Rectangle) []Rectangle {
	padded := lo.Map(boxes, func(rect Rectangle, _ int) Rectangle { return Pad(rect, config.paddingSize) })

	fmt.Print("Start to push rectangle with intersections to stack...\n")
	stack, settled := BuildRectangleStack(padded)
	indexer := NewBoxIndexer(config.indexStrategy, append(settled, takenAreas...))
	geoCenter := geometricCenter{}
	for _, rect := range settled {
		geoCenter.Add(rect)
	}

	fmt.Print("Start to process the rectangles in stack...\n")
	for stack.Len() != 0 {
		box := stack.Pop()
		gCenter := geoCenter.Point()
		bCenter := box.Center()
		dir := Vector{bCenter.X - gCenter.X, bCenter.Y - gCenter.Y}

		result := move(box, dir.Unit(), indexer)

		settled = append(settled, result)
		geoCenter.Add(result)
		indexer.Insert(result)
		fmt.Printf("A new box is successfully processed: %v\n", result)
	}

	unpadded := lo.Map(settled, func(rect Rectangle, _ int) Rectangle { return Unpad(rect, config.paddingSize) })
	return unpadded
}

func move(box Rectangle, vec Vector, indexer BoxIndexer) Rectangle {
	ratio := config.initialStep
	for {
		for _, theta := range config.thetas {
			newBox := step(box, vec, ratio, theta)
			if indexer.Check(newBox) {
				return newBox
			}
		}
		ratio += config.ratioInterval
		fmt.Printf("Using Ratio: %f\n", ratio)
	}
}

func step(box Rectangle, vector Vector, ratio float64, theta float64) Rectangle {
	newVec := vector.Mul(ratio)
	newVec = newVec.Rotate(theta)
	return box.Translate(newVec)
}
