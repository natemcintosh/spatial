package spatial

import "testing"

func TestLess(t *testing.T) {
	testCases := []struct {
		desc string
		p    Point2d
		q    Point2d
		want bool
	}{
		{
			desc: "all less",
			p:    Point2d{X: 0, Y: 0},
			q:    Point2d{X: 1, Y: 1},
			want: true,
		},
		{
			desc: "all greater",
			p:    Point2d{X: 1, Y: 1},
			q:    Point2d{X: 0, Y: 0},
			want: false,
		},
		{
			desc: "equal",
			p:    Point2d{X: 1, Y: 1},
			q:    Point2d{X: 1, Y: 1},
			want: false,
		},
		{
			desc: "x less",
			p:    Point2d{X: 0, Y: 5},
			q:    Point2d{X: 1, Y: 0},
			want: true,
		},
		{
			desc: "x greater",
			p:    Point2d{X: 1, Y: 0},
			q:    Point2d{X: 0, Y: 5},
			want: false,
		},
		{
			desc: "y less",
			p:    Point2d{X: 5, Y: -1},
			q:    Point2d{X: 5, Y: 0},
			want: true,
		},
		{
			desc: "y greater",
			p:    Point2d{X: 5, Y: 0},
			q:    Point2d{X: 5, Y: -1},
			want: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Less(tC.p, tC.q)
			if got != tC.want {
				t.Errorf("Less(%v, %v) = %v, want %v", tC.p, tC.q, got, tC.want)
			}
		})
	}
}
