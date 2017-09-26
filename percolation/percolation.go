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

	// Add an additional 2 slots at the end of the slice to store references to the top and bottom rows.
	sites := uf.NewWeightedQuickUnion(n + 2)

	topID := n - 2
	bottomID := n - 3
	// Connect the top and bottom rows.
	for i := 0; i < size; i++ {
		sites.Union(topID, i)
		sites.Union(bottomID, n-4+i)
	}

	cmp := func(i, j *uf.Node) bool {
		// topID is always the root.
		if i.Root == topID {
			return true
		}

		if j.Root == bottomID && i.Root != topID {
			return true
		}

		return false
	}

	sites.Get(topID).Comparer = cmp
	sites.Get(bottomID).Comparer = cmp
	return &Percolator{sites, 1, n, size, topID, bottomID}, nil
}

// Open requires 0 based indexing.
func (p *Percolator) Open(row, col int) error {
	if (row+1)*(col+1) > p.n {
		return fmt.Errorf("row:%d col:%d out of bounds for a matrix of size:%d", row, col, p.n)
	}

	base := p.len * col
	target := base + row + col
	targetNode := p.sites.Get(target)

	// Don't open already opened sites.
	if !targetNode.Value {
		p.openSites++
		targetNode.Value = true

		if row != 0 {
			leftNeighbor := base + (row - 1) + col
			p.connect(target, leftNeighbor)
		}

		if row < (p.len - 1) {
			rightNeighbor := base + (row + 1) + col
			p.connect(target, rightNeighbor)
		}

		if col != 0 {
			topNeighbor := base + row + (col - 1)
			p.connect(target, topNeighbor)
		}

		if col < (p.len - 1) {
			bottomNeighbor := base + row + (col + 1)
			p.connect(target, bottomNeighbor)
		}
	}

	return nil
}

func (p *Percolator) connect(target, neighbor int) {
	if p.sites.Get(neighbor).Value {
		p.sites.Union(target, neighbor)
	}
}
