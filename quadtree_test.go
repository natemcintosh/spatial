package spatial

import (
	"math"
	"math/rand"
	"testing"

	"github.com/natemcintosh/set"
)

func identity[T any](t T) T {
	return t
}

func TestNewQuadtree(t *testing.T) {
	b := Bound{Point2d{-10, -10}, Point2d{10, 10}}
	tree := NewQuadtree(b, identity[Point2d], Point2dDistance)

	pts := []Point2d{
		{-5, -5},
		{0, 0},
		{5, 5},
	}
	pts_set := set.NewSet(pts)

	for _, p := range pts {
		tree.Add(p)
	}

	got := set.NewSet(tree.Slice())

	if !got.Equals(pts_set) {
		t.Errorf("got %v, want %v", got, pts)
	}
}

func FuzzAdd(f *testing.F) {
	f.Add(int64(1))

	// Instead, try having it set a random seed, and from there populate a much larger
	// number of points
	f.Fuzz(func(t *testing.T, seed int64) {
		// Set the seed. Need fuzz tests to always give the same output for a given input
		rand.Seed(seed)

		// How many points will we add?
		var minx, maxx, miny, maxy float64
		n_points := rand.Int31n(200)
		pts := make([]Point2d, n_points)
		if n_points > 0 {
			x := rand.Float64() * 1000
			y := rand.Float64() * 1000
			minx = x
			maxx = x
			miny = y
			maxy = y
			pts[0] = Point2d{x, y}

			for i := 1; i < int(n_points); i++ {
				x = rand.Float64() * 1000
				if rand.Float32() > 0.5 {
					x *= -1
				}
				minx = math.Min(minx, x)
				maxx = math.Max(maxx, x)

				y = rand.Float64() * 1000
				if rand.Float32() > 0.5 {
					y *= -1
				}
				miny = math.Min(miny, y)
				maxy = math.Max(maxy, y)

				pts[i] = Point2d{x, y}
			}
		}

		// Create tree and add items
		bbox := Bound{Point2d{minx, miny}, Point2d{maxx, maxy}}
		tree := NewQuadtree(bbox, identity[Point2d], Point2dDistance)
		for _, pt := range pts {
			tree.Add(pt)
		}

		// Get items out of tree
		s := tree.Slice()
		got := set.NewSet(s)

		// Create a set from the points
		want := set.NewSet(pts)

		// Compare
		if !got.Equals(want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
