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
	ErrBadNumberOfPoints    = errors.New("asked for too many points or < 0")
	ErrPointOutsideOfBounds = errors.New("point outside of bounds")
)

type Point2d struct {
	X, Y float64
}

func Point2dDistance(p, q Point2d) float64 {
	return math.Sqrt((p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y))
}

// Bound is a copy of
// https://github.com/paulmach/orb/blob/215f32c132d13f906979dbb36bd24c1e0511b6d2/bound.go#L12
type Bound struct {
	Min, Max Point2d
}

// Extend grows the bound to include the new point.
func (b Bound) Extend(point Point2d) Bound {
	// already included, no big deal
	if b.Contains(point) {
		return b
	}

	return Bound{
		Min: Point2d{
			math.Min(b.Min.X, point.X),
			math.Min(b.Min.Y, point.Y),
		},
		Max: Point2d{
			math.Max(b.Max.X, point.X),
			math.Max(b.Max.Y, point.Y),
		},
	}
}

// Union extends this bound to contain the union of this and the given bound.
func (b Bound) Union(other Bound) Bound {
	if other.IsEmpty() {
		return b
	}

	b = b.Extend(other.Min)
	b = b.Extend(other.Max)
	b = b.Extend(other.LeftTop())
	b = b.Extend(other.RightBottom())

	return b
}

// Contains determines if the point is within the bound.
// Points on the boundary are considered within.
func (b Bound) Contains(point Point2d) bool {
	if point.Y < b.Min.Y || b.Max.Y < point.Y {
		return false
	}

	if point.X < b.Min.X || b.Max.X < point.X {
		return false
	}

	return true
}

// Intersects determines if two bounds intersect.
// Returns true if they are touching.
func (b Bound) Intersects(bound Bound) bool {
	if (b.Max.X < bound.Min.X) ||
		(b.Min.X > bound.Max.X) ||
		(b.Max.Y < bound.Min.Y) ||
		(b.Min.Y > bound.Max.Y) {
		return false
	}

	return true
}

// Pad extends the bound in all directions by the given value.
func (b Bound) Pad(d float64) Bound {
	b.Min.X -= d
	b.Min.Y -= d

	b.Max.X += d
	b.Max.Y += d

	return b
}

// Center returns the center of the bounds by "averaging" the x and y coords.
func (b Bound) Center() Point2d {
	return Point2d{
		(b.Min.X + b.Max.X) / 2.0,
		(b.Min.Y + b.Max.Y) / 2.0,
	}
}

// Top returns the top of the bound.
func (b Bound) Top() float64 {
	return b.Max.Y
}

// Bottom returns the bottom of the bound.
func (b Bound) Bottom() float64 {
	return b.Min.Y
}

// Right returns the right of the bound.
func (b Bound) Right() float64 {
	return b.Max.X
}

// Left returns the left of the bound.
func (b Bound) Left() float64 {
	return b.Min.X
}

// LeftTop returns the upper left point of the bound.
func (b Bound) LeftTop() Point2d {
	return Point2d{b.Left(), b.Top()}
}

// RightBottom return the lower right point of the bound.
func (b Bound) RightBottom() Point2d {
	return Point2d{b.Right(), b.Bottom()}
}

// IsEmpty returns true if it contains zero area or if
// it's in some malformed negative state where the left point is larger than the right.
// This can be caused by padding too much negative.
func (b Bound) IsEmpty() bool {
	return b.Min.X > b.Max.X || b.Min.Y > b.Max.Y
}

// IsZero return true if the bound just includes just null island.
func (b Bound) IsZero() bool {
	return b.Max == Point2d{} && b.Min == Point2d{}
}

// Bound returns the the same bound.
func (b Bound) Bound() Bound {
	return b
}

// Equal returns if two bounds are equal.
func (b Bound) Equal(c Bound) bool {
	return b.Min == c.Min && b.Max == c.Max
}
