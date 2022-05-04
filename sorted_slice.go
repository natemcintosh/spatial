package spatial

import (
	"golang.org/x/exp/slices"
)

// LinearSlice is meant for testing the other types. It is optimized for testing ease
type LinearSlice[T comparable] struct {
	s   []T
	dst func(T, T) float64
}

func NewLinearSlice[T comparable](points []T, distance_func func(T, T) float64) LinearSlice[T] {
	return LinearSlice[T]{points, distance_func}
}

// Insert inserts p into s.
func (s *LinearSlice[T]) Insert(p T) {
	s.s = append(s.s, p)
}

// Remove removes the first instance of p from s. If it is not found, nothing happens
func (s *LinearSlice[T]) Remove(p T) {
	idx := slices.Index(s.s, p)
	if idx >= 0 {
		s.s = slices.Delete(s.s, idx, idx)
	}
}

// Contains returns true if p is in s.
func (s LinearSlice[T]) Contains(p T) bool {
	idx := slices.Index(s.s, p)
	if idx >= 0 {
		return true
	}
	return false
}

// Nearest returns the nearest point in s to p.
func (s LinearSlice[T]) Nearest(p T) T {
	// Just search over all the points
	smallest_item := s.s[0]
	smallest_dist := s.dst(s.s[0], p)
	for _, v := range s.s {
		if s.dst(v, p) < smallest_dist {
			smallest_dist = s.dst(v, p)
			smallest_item = v
		}
	}

	return smallest_item
}

// NearestN returns the nearest n points in s to p. If `n > len(s)`, returns an error
func (s LinearSlice[T]) NearestN(p T, n int) ([]T, error) {
	if n > len(s.s) || n < 0 {
		return nil, ErrBadNumberOfPoints
	}

	// Make a copy
	result := make([]T, len(s.s))
	copy(result, s.s)

	// Sort it
	slices.SortFunc(result, func(a, b T) bool {
		return s.dst(a, p) < s.dst(b, p)
	})

	// Return only the first `n` items
	return result[:n], nil

}

// InRange returns all points in s within range of p.
func (s LinearSlice[T]) InRange(p T, r float64) []T {
	result := make([]T, 0)

	for _, v := range s.s {
		if s.dst(v, p) <= r {
			result = append(result, v)
		}
	}

	return result
}

// InRangeCount counts all the points in `s` that are closer to `p` than `r`
func (s LinearSlice[T]) InRangeCount(p T, r float64) int {
	count := 0
	for _, v := range s.s {
		if s.dst(v, p) <= r {
			count += 1
		}
	}

	return count
}
