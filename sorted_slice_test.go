package spatial

import "testing"

func TestInsert(t *testing.T) {
	s := NewLinearSlice([]Point2d{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
	})
	p := Point2d{2, 3}
	want := NewLinearSlice([]Point2d{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{2, 3},
	})
	s.Insert(p)
	if s.s[len(s.s)-1] != p {
		t.Errorf("got %v; want %v", s.s, want)
	}
}

func TestNearest(t *testing.T) {
	s := NewLinearSlice([]Point2d{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
	})
	testCases := []struct {
		desc string
		p    Point2d
		want Point2d
	}{
		{
			desc: "below",
			p:    Point2d{-1, -1},
			want: s.s[0],
		},
		{
			desc: "above",
			p:    Point2d{7, 7},
			want: s.s[len(s.s)-1],
		},
		{
			desc: "middle",
			p:    Point2d{3, 3},
			want: s.s[3],
		},
		{
			desc: "close under",
			p:    Point2d{2.9, 2.9},
			want: s.s[3],
		},
		{
			desc: "close over",
			p:    Point2d{3.1, 3.1},
			want: s.s[3],
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := s.Nearest(tC.p)
			if got != tC.want {
				t.Errorf("got %v, want %v", got, tC.want)
			}
		})
	}
}

func TestNearestN(t *testing.T) {
	s := NewLinearSlice([]Point2d{
		{-12.3, -12.3},
		{-0.99, -0.99},
		{-0.25, -0.25},
		{0.25, 0.25},
		{1., 1},
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
		{
			desc: "2 nearest (0, 0)",
			p:    Point2d{0, 0},
			n:    2,
			want: []Point2d{
				{-0.25, -0.25},
				{0.25, 0.25},
			},
		},
		{
			desc: "2 nearest (-0.8, 0)",
			p:    Point2d{-0.8, 0},
			n:    2,
			want: []Point2d{
				{-0.25, -0.25},
				{-0.99, -0.99},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, _ := s.NearestN(tC.p, tC.n)
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
