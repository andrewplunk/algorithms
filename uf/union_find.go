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
}

// Finder searches for elements in the graph.
type Finder interface {
	Find(int, bool) int
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
	Root  int
	Value interface{}
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
	for i := 0; i < n; i++ {
		ids[i] = &Node{Root: i}
		sz[i] = 1
	}

	return &WeightedQuickUnion{
		ids: ids,
		sz:  sz,
	}
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
	i := w.root(p).Root
	j := w.root(q).Root
	if i == j {
		return
	}

	if w.sz[i] < w.sz[j] {
		w.ids[i] = w.ids[j]
		w.sz[j] += w.sz[i]
		return
	}
	w.ids[j] = w.ids[i]
	w.sz[i] = j
}

// Connected returns true if two nodes connected, false otherwise.
func (w *WeightedQuickUnion) Connected(p, q int) bool {
	return w.root(p) == w.root(q)
}

// Find ...
func (w *WeightedQuickUnion) Find(p int) *Node {
	return w.root(p)
}
