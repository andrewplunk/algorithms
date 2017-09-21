/*Package unionfind efficiently support the following functions on network of nodes:
- union: connect two graph objects.
- find: determine if there is a path between two graph objects.

Connections:
- reflexive
- symmetric
- transitive
*/
package unionfind

// UF the union finder interface
type UF interface {
	Union(p, q int)
	Connected(p, q int) bool
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
func NewQuickFind(n int) UF {
	return &QuickFind{
		ids: initIDs(n),
		len: n,
	}
}

func (qf *QuickFind) String() string {
	return "QuickFind"
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
func NewQuickUnion(n int) UF {
	return &QuickUnion{
		ids: initIDs(n),
	}
}

func (qu *QuickUnion) String() string {
	return "QuickUnion"
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

// WeightedQuickUnion optimizes the union operation.
type WeightedQuickUnion struct {
	ids []int
	sz  []int
}

// NewWeightedQuickUnion using an array model p and q being connected if they have the same id.
func NewWeightedQuickUnion(n int) UF {
	ids := make([]int, n)
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i
		sz[i] = 1
	}

	return &WeightedQuickUnion{
		ids: ids,
		sz:  sz,
	}
}

func (w *WeightedQuickUnion) String() string {
	return "WeightedQuickUnion"
}

func (w *WeightedQuickUnion) root(p int) int {
	for p != w.ids[p] {
		// path compression
		w.ids[p] = w.ids[w.ids[p]]
		p = w.ids[p]
	}
	return p
}

// Union connect two nodes.
func (w *WeightedQuickUnion) Union(p, q int) {
	i := w.root(p)
	j := w.root(q)
	if i == j {
		return
	}

	if w.sz[i] < w.sz[j] {
		w.ids[i] = j
		w.sz[j] += w.sz[i]
		return
	}
	w.ids[j] = i
	w.sz[i] = j
}

// Connected returns true if two nodes connected, false otherwise.
func (w *WeightedQuickUnion) Connected(p, q int) bool {
	return w.root(p) == w.root(q)
}
