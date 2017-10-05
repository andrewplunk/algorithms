package percolation

import (
	"fmt"
	"math"

	"github.com/andrewplunk/algorithms/uf"
)

// type percolator interface{
// 	Open(int, int)
// 	IsOpen(int, int) bool
// 	IsFull(int, int) bool
// 	OpenSites() int
// 	Percolates() bool
// }

// Percolator ...
type Percolator struct {
	sites     *uf.WeightedQuickUnion
	openSites int
	n         int
	len       int
	topID     int
	bottomID  int
}

// New ...
func New(n int) (*Percolator, error) {
	sqrt := math.Sqrt(float64(n))
	size := int(sqrt)
	if sqrt*sqrt != float64(size*size) {
		return nil, fmt.Errorf("This percolator is meant to operate on square matrixes and %d does not have a perfect square.", n)
	}

	topID := n
	bottomID := n + 1

	cmp := func(w *uf.WeightedQuickUnion, i, j *uf.Node) bool {
		// topID is always the root.
		if i.Root == topID {
			return true
		}

		if i.Root == bottomID && j.Root != topID {
			return true
		}

		return w.Get(i.Root).Size > w.Get(j.Root).Size
	}

	// Add an additional 2 slots at the end of the slice to store references to the top and bottom rows.
	sites := uf.NewWeightedQuickUnion(n+2, uf.Comparer(cmp))

	p := &Percolator{sites, 0, n, size, topID, bottomID}
	// Connect the top and bottom rows.
	for i := 0; i < size; i++ {
		sites.Union(topID, i)
		sites.Union(bottomID, p.calculateIndex(size-1, i))
	}

	return p, nil
}

// Percolates returns true if a site in the bottom row is connected to the top row.
func (p *Percolator) Percolates() bool {
	return p.sites.Connected(p.topID, p.bottomID)
}

// IsFull returns true if the site is open and has a path to the top row.
func (p *Percolator) IsFull(row, col int) bool {
	return p.sites.Connected(p.calculateIndex(row, col), p.topID)
}

// IsOpen returns true if the site is open.
func (p *Percolator) IsOpen(row, col int) bool {
	return p.sites.Get(p.calculateIndex(row, col)).Value
}

func (p *Percolator) calculateIndex(row, col int) int {
	return ((p.len - 1) * row) + row + col
}

// OpenSites returns the number of sites open
func (p *Percolator) OpenSites() int {
	return p.openSites
}

// Open requires 0 based indexing.
func (p *Percolator) Open(row, col int) error {
	if row > p.len || col > p.len {
		return fmt.Errorf("row:%d col:%d out of bounds for a matrix of size:%d", row, col, p.n)
	}

	target := p.calculateIndex(row, col)
	targetNode := p.sites.Get(target)

	// Don't open already opened sites.
	if !targetNode.Value {
		p.openSites++
		targetNode.Value = true

		if col != 0 {
			// left
			p.connect(target, p.calculateIndex(row, col-1))
		}

		if col < (p.len - 1) {
			// right
			p.connect(target, p.calculateIndex(row, col+1))
		}

		if row != 0 {
			// up
			p.connect(target, p.calculateIndex(row-1, col))
		}

		if row < (p.len - 1) {
			// down
			p.connect(target, p.calculateIndex(row+1, col))
		}
	}

	return nil
}

func (p *Percolator) connect(target, neighbor int) {
	if p.sites.Get(neighbor).Value {
		p.sites.Union(target, neighbor)
	}
}
