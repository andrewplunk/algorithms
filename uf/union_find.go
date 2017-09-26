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

// Node a container to use rather than the typical union-find int id.
// This lets us store information at each site.
type Node struct {
	Root int
	// Value switch to interface{} if this needs to be made generic.
	Value    bool
	Comparer func(*Node, *Node) bool
}

// CompareTo ...
func (n *Node) CompareTo(other *Node) bool {
	return n.Comparer(n, other)
}

// WeightedQuickUnion optimizes the union operation.
type WeightedQuickUnion struct {
	ids []*Node
	sz  []int
}

// NewWeightedQuickUnion using an array model p and q being connected if they have the same id.
func NewWeightedQuickUnion(n int) *WeightedQuickUnion {
	ids := make([]*Node, n)
	sz := make([]int, n)

	wqu := &WeightedQuickUnion{
		ids: ids,
		sz:  sz,
	}

	for i := 0; i < n; i++ {
		ids[i] = &Node{Root: i, Value: false, Comparer: wqu.defaultComparer}
		sz[i] = 1
	}
	return wqu
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

func (w *WeightedQuickUnion) defaultComparer(i, j *Node) bool {
	return w.sz[i.Root] > w.sz[j.Root]
}

// Union connect two nodes.
func (w *WeightedQuickUnion) Union(p, q int) {
	i := w.root(p)
	j := w.root(q)
	if i == j {
		return
	}

	if i.CompareTo(j) {
		w.ids[j.Root] = w.ids[i.Root]
		w.sz[i.Root] = j.Root
		return
	}
	w.ids[i.Root] = w.ids[j.Root]
	w.sz[j.Root] += w.sz[i.Root]
}

// Connected returns true if two nodes connected, false otherwise.
func (w *WeightedQuickUnion) Connected(p, q int) bool {
	return w.root(p) == w.root(q)
}

// Find ...
func (w *WeightedQuickUnion) Find(p int) *Node {
	return w.root(p)
}

// Get ...
func (w *WeightedQuickUnion) Get(p int) *Node {
	return w.ids[p]
}
