package spatial

type node[T any] struct {
	// The data associated with this node
	item T

	is_initialized bool

	// Pointers to each corner
	children [4]*node[T]
}

// Quadtree is based heavily off of
// https://github.com/paulmach/orb/blob/215f32c132d13f906979dbb36bd24c1e0511b6d2/quadtree/quadtree.go
type Quadtree[T any] struct {
	root  node[T]
	bound Bound

	// A function to get the 2D position out this node
	pnt func(T) Point2d

	// A function to get the distance between two nodes
	dst func(T, T) float64
}

func NewQuadtree[T any](
	bounds Bound,
	get_point func(T) Point2d,
	calc_distance func(T, T) float64,
) *Quadtree[T] {
	root_node := node[T]{is_initialized: false}
	return &Quadtree[T]{root_node, bounds, get_point, calc_distance}
}

// Add a point to the Quadtree. It must be within the bounds of the tree.
func (q *Quadtree[T]) Add(p T) error {
	point := q.pnt(p)

	if !q.bound.Contains(point) {
		return ErrPointOutsideOfBounds
	}

	if !q.root.is_initialized { // the start of the tree
		q.root = node[T]{p, true, [4]*node[T]{nil, nil, nil, nil}}
		return nil
	}

	q.add(
		&q.root,
		p,
		q.pnt(p),
		q.bound.Left(),
		q.bound.Right(),
		q.bound.Bottom(),
		q.bound.Top(),
	)

	return nil
}

// add is the recursive search to find a place to add the point
func (q *Quadtree[T]) add(
	n *node[T],
	p T,
	point Point2d,
	left float64,
	right float64,
	bottom float64,
	top float64,
) {
	i := 0

	// Figure out which child of this internal node the point is in
	child_y := (bottom + top) / 2.0
	if point.Y <= child_y {
		top = child_y
		i = 2
	} else {
		bottom = child_y
	}

	child_x := (left + right) / 2.0
	if point.X >= child_x {
		left = child_x
		i += 1
	} else {
		right = child_x
	}

	// It is no longer a leaf node
	n.is_initialized = true

	if n.children[i] == nil {
		n.children[i] = &node[T]{p, true, [4]*node[T]{nil, nil, nil, nil}}
		return
	}

	// Proceed down to the child to see if it's a leaf yet and we can add the pointer there
	q.add(n.children[i], p, point, left, right, bottom, top)
}

// Slice is designed to be as obviously correct as possible to start
// We'll see about greater efficiency later
func (q *Quadtree[T]) Slice() []T {
	// Create the slice
	items := make([]T, 0)

	if !q.root.is_initialized {
		return items
	}

	items = append(items, q.root.item)

	for _, child_node := range q.root.children {
		if child_node != nil {
			items = append(items, child_node.slice()...)
		}
	}

	return items
}

func (n *node[T]) slice() []T {
	// We have already established that this item is not nil
	s := []T{n.item}

	// Check each child
	for _, child_node := range n.children {
		if child_node != nil {
			s = append(s, child_node.slice()...)
		}
	}

	return s
}
