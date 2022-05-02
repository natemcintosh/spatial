package spatial

type Spatial interface {
	Insert(p Point2d)
	Remove(p Point2d)
	Contains(p Point2d) bool
	Nearest(p Point2d) Point2d
	NearestN(p Point2d, n int) []Point2d
	InRange(p Point2d, r float64) []Point2d
}

type Point2d struct {
	X, Y float64
}

// Compare sorts by X, then Y. -1 if p < q, 0 if equal, and +1 if p > q
func Compare(p, q Point2d) int {
	if p.X < q.X {
		return -1
	} else if p.X > q.X {
		return 1
	} else if p.Y < q.Y {
		return -1
	} else if p.Y > q.Y {
		return 1
	}
	return 0
}

// Less returns true if p < q
func Less(p, q Point2d) bool {
	if p.X < q.X {
		return true
	} else if p.X > q.X {
		return false
	} else if p.Y < q.Y {
		return true
	}
	return false
}
