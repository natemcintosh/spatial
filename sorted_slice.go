package spatial

import (
	"golang.org/x/exp/slices"
)

type SortedSlice = []Point2d

func NewSortedSlice(points []Point2d) SortedSlice {
	sorted := make(SortedSlice, len(points))
	copy(sorted, points)
	slices.SortFunc(sorted, Less)
	return sorted
}
