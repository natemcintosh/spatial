package spatial

import (
	"errors"
	"math"
)

type Spatial interface {
	Insert(p Point2d)
	Remove(p Point2d)
	Contains(p Point2d) bool
	Nearest(p Point2d) Point2d
	NearestN(p Point2d, n int) []Point2d
	InRange(p Point2d, r float64) []Point2d
	InRangeCount(p Point2d, r float64) int
}

var (
	ErrBadNumberOfPoints = errors.New("asked for too many points or < 0")
)

type Point2d struct {
	X, Y float64
}

func Point2dDistance(p, q Point2d) float64 {
	return math.Sqrt((p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y))
}
