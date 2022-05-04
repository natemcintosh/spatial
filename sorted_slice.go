package spatial

import (
	"golang.org/x/exp/slices"
)

type LinearSlice struct {
	s []Point2d
}

func NewLinearSlice(points []Point2d) LinearSlice {
	return LinearSlice{points}
}

// Insert inserts p into s.
func (s *LinearSlice) Insert(p Point2d) {
	s.s = append(s.s, p)
}

// Remove removes the first instance of p from s. If it is not found, nothing happens
func (s *LinearSlice) Remove(p Point2d) {
	idx := slices.Index(s.s, p)
	if idx >= 0 {
		s.s = slices.Delete(s.s, idx, idx)
	}
}

// Contains returns true if p is in s.
func (s LinearSlice) Contains(p Point2d) bool {
	idx := slices.Index(s.s, p)
	if idx >= 0 {
		return true
	}
	return false
}

// Nearest returns the nearest point in s to p.
func (s LinearSlice) Nearest(p Point2d) Point2d {
	// Just search over all the points
	smallest_item := s.s[0]
	smallest_dist := Distance(s.s[0], p)
	for _, v := range s.s {
		if Distance(v, p) < smallest_dist {
			smallest_dist = Distance(v, p)
			smallest_item = v
		}
	}

	return smallest_item
}

// NearestN returns the nearest n points in s to p. If `n > len(s)`, returns an error
func (s LinearSlice) NearestN(p Point2d, n int) ([]Point2d, error) {
	if n > len(s.s) || n < 0 {
		return nil, ErrBadNumberOfPoints
	}

	// Make a copy
	result := make([]Point2d, len(s.s))
	copy(result, s.s)

	// Sort it
	slices.SortFunc(result, func(a, b Point2d) bool {
		return Distance(a, p) < Distance(b, p)
	})

	// Return only the first `n` items
	return result[:n], nil

}

// InRange returns all points in s within range of p.
func (s LinearSlice) InRange(p Point2d, r float64) []Point2d {
	idx, _ := slices.BinarySearchFunc(s.s, p, Compare)
	if idx < 0 {
		idx = -idx - 1
	}
	if idx >= len(s.s) {
		idx = len(s.s) - 1
	}
	var i, j int
	for i = idx; i >= 0; i-- {
		if Less(s.s[i], p) {
			break
		}
	}
	for j = idx + 1; j < len(s.s); j++ {
		if Less(p, s.s[j]) {
			break
		}
	}
	return s.s[i+1 : j]
}
