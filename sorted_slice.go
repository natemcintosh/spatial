package spatial

import (
	"golang.org/x/exp/slices"
)

type SortedSlice struct {
	s []Point2d
}

func NewSortedSlice(points []Point2d) SortedSlice {
	sorted := make([]Point2d, len(points))
	copy(sorted, points)
	slices.SortFunc(sorted, Less)
	return SortedSlice{sorted}
}

// Insert inserts p into s.
func (s SortedSlice) Insert(p Point2d) {
	idx, _ := slices.BinarySearchFunc(s.s, p, Compare)
	s.s = slices.Insert(s.s, idx, p)
}

// Remove removes the first instance of p from s. If it is not found, nothing happens
func (s SortedSlice) Remove(p Point2d) {
	idx, exists := slices.BinarySearchFunc(s.s, p, Compare)
	if exists {
		s.s = slices.Delete(s.s, idx, idx)
	}
}

// Contains returns true if p is in s.
func (s SortedSlice) Contains(p Point2d) bool {
	_, exists := slices.BinarySearchFunc(s.s, p, Compare)
	return exists
}

// Nearest returns the nearest point in s to p.
func (s SortedSlice) Nearest(p Point2d) Point2d {
	idx, _ := slices.BinarySearchFunc(s.s, p, Compare)
	if idx < 0 {
		idx = -idx - 1
	}
	if idx >= len(s.s) {
		idx = len(s.s) - 1
	}
	return s.s[idx]
}

// NearestN returns the nearest n points in s to p.
func (s SortedSlice) NearestN(p Point2d, n int) []Point2d {
	idx, _ := slices.BinarySearchFunc(s.s, p, Compare)
	if idx < 0 {
		idx = -idx - 1
	}
	if idx >= len(s.s) {
		idx = len(s.s) - 1
	}
	if n > len(s.s) {
		n = len(s.s)
	}
	if idx+n > len(s.s) {
		idx = len(s.s) - n
	}
	return s.s[idx : idx+n]
}

// InRange returns all points in s within range of p.
func (s SortedSlice) InRange(p Point2d, r float64) []Point2d {
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
