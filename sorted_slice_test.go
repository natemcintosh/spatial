package spatial

import "testing"

func TestNearestds(t *testing.T) {
	s := NewSortedSlice([]Point2d{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
	})

	pbelow := Point2d{-1, -1}
	pabove := Point2d{7, 7}
	pmiddle := Point2d{3, 3}

	if s.Nearest(pbelow) != s.s[0] {
		t.Errorf("expected %v, got %v", Point2d{0, 0}, s.Nearest(pbelow))
	}
	if s.Nearest(pabove) != s.s[len(s.s)-1] {
		t.Errorf("expected %v, got %v", Point2d{6, 6}, s.Nearest(pabove))
	}
	if s.Nearest(pmiddle) != s.s[3] {
		t.Errorf("expected %v, got %v", Point2d{3, 3}, s.Nearest(pmiddle))
	}
}

func TestNearestN(t *testing.T) {
	s := NewSortedSlice([]Point2d{
		{-12.3, -12.3},
		{-0.99, -0.99},
		{-0.25, -0.25},
		{0.25, 0.25},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
	})
	testCases := []struct {
		desc string
		p    Point2d
		n    int
		want []Point2d
	}{
		// {
		// 	desc: "2 nearest (0, 0)",
		// 	p:    Point2d{0, 0},
		// 	n:    2,
		// 	want: []Point2d{
		// 		{-0.25, -0.25},
		// 		{0.25, 0.25},
		// 	},
		// },
		// {
		// 	desc: "2 nearest (-0.8, 0)",
		// 	p:    Point2d{-0.8, 0},
		// 	n:    2,
		// 	want: []Point2d{
		// 		{-0.99, -0.99},
		// 		{-0.25, -0.25},
		// 	},
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := s.NearestN(tC.p, tC.n)
			if len(got) != len(tC.want) {
				t.Errorf("got %v, want %v", got, tC.want)
			}
			for i, w := range tC.want {
				if w != got[i] {
					t.Errorf("got %v, want %v", got, tC.want)
				}
			}
		})
	}
}
