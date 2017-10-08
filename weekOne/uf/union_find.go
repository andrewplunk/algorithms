/*Package uf efficiently support the following functions on network of nodes:
- union: connect two graph objects.
- find: determine if there is a path between two graph objects.

Connections:
- reflexive
- symmetric
- transitive
*/
package uf

// UF the union finder interface
type UF interface {
	Union(p, q int)
	Connected(p, q int) bool
	Find(int) *Node
}

func initIDs(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i
	}
	return ids
}

// QuickFind implement union find using an array where things are connected if they have the same id.
type QuickFind struct {
	ids []int
	len int
}

// NewQuickFind using an array model p and q being connected if they have the same id.
func NewQuickFind(n int) *QuickFind {
	return &QuickFind{
		ids: initIDs(n),
		len: n,
	}
}

// Union connect two nodes.
func (qf *QuickFind) Union(p, q int) {
	qid, pid := qf.ids[q], qf.ids[p]
	for i := 0; i < qf.len; i++ {
		if qf.ids[i] == pid {
			qf.ids[i] = qid
		}
	}
}

// Connected returns true if two nodes connected, false otherwise.
func (qf *QuickFind) Connected(p, q int) bool {
	return qf.ids[p] == qf.ids[q]
}

// QuickUnion optimizes the union operation.
type QuickUnion struct {
	ids []int
}

// NewQuickUnion using an array model p and q being connected if they have the same id.
func NewQuickUnion(n int) *QuickUnion {
	return &QuickUnion{
		ids: initIDs(n),
	}
}

func (qu *QuickUnion) root(p int) int {
	for p != qu.ids[p] {
		p = qu.ids[p]
	}
	return p
}

// Union connect two nodes.
func (qu *QuickUnion) Union(p, q int) {
	qu.ids[p] = qu.root(q)
}

// Connected returns true if two nodes connected, false otherwise.
func (qu *QuickUnion) Connected(p, q int) bool {
	return qu.root(p) == qu.root(q)
}

type Comparer func(*WeightedQuickUnion, *Node, *Node) bool

// Node a container to use rather than the typical union-find int id.
// This lets us store information at each site.
type Node struct {
	Root int
	// Value switch to interface{} if this needs to be made generic.
	Value bool
	Size  int
}

// WeightedQuickUnion optimizes the union operation.
type WeightedQuickUnion struct {
	ids      []*Node
	comparer Comparer
}

// NewWeightedQuickUnion using an array model p and q being connected if they have the same id.
func NewWeightedQuickUnion(n int, comparer Comparer) *WeightedQuickUnion {
	ids := make([]*Node, n)
	if comparer == nil {
		comparer = func(w *WeightedQuickUnion, i, j *Node) bool {
			return i.Size > j.Size
		}
	}

	for i := 0; i < n; i++ {
		ids[i] = &Node{
			Root:  i,
			Value: false, Size: 1,
		}
	}
	return &WeightedQuickUnion{ids: ids, comparer: comparer}
}

func (w *WeightedQuickUnion) root(p int) (n *Node) {
	//TODO make this cleaner
	n = w.ids[p]
	for p != w.ids[p].Root {
		// path compression
		w.ids[p] = w.ids[w.ids[p].Root]
		p = w.ids[p].Root
		n = w.ids[p]
	}
	return
}

// Union connect two nodes.
func (w *WeightedQuickUnion) Union(p, q int) {
	i := w.root(p)
	j := w.root(q)
	if i == j {
		return
	}

	if w.comparer(w, i, j) {
		jSize := w.ids[j.Root].Size
		w.ids[j.Root] = w.ids[i.Root]
		w.ids[i.Root].Size += jSize
		return
	}
	iSize := w.ids[i.Root].Size
	w.ids[i.Root] = w.ids[j.Root]
	w.ids[j.Root].Size += iSize
}

// Connected returns true if two nodes connected, false otherwise.
func (w *WeightedQuickUnion) Connected(p, q int) bool {
	pr := w.root(p)
	qr := w.root(q)
	return pr == qr
}

// Find ...
func (w *WeightedQuickUnion) Find(p int) *Node {
	return w.root(p)
}

// Get ...
func (w *WeightedQuickUnion) Get(p int) *Node {
	return w.ids[p]
}
