package percolation

import "testing"

type expectation func(t *testing.T, p *Percolator)

func open(row, col int, expected bool) expectation {
	return func(t *testing.T, p *Percolator) {
		if p.IsOpen(row, col) != expected {
			t.Fatalf("Expected row:%d, col:%d to be open:%t!", row, col, expected)
		}
	}
}

func full(row, col int, expected bool) expectation {
	return func(t *testing.T, p *Percolator) {
		if p.IsFull(row, col) != expected {
			t.Fatalf("Expected row:%d, col:%d to be full:%t!", row, col, expected)
		}
	}
}

func percolates(expected bool) expectation {
	return func(t *testing.T, p *Percolator) {
		if p.Percolates() != expected {
			t.Fatalf("Expected to percolate:%t!", expected)
		}
	}
}

func openSites(expected int) expectation {
	return func(t *testing.T, p *Percolator) {
		actual := p.OpenSites()
		if actual != expected {
			t.Fatalf("Expected open sites:%d actual:%d!", expected, actual)
		}
	}
}

func (p *Percolator) build(t *testing.T, row, col int) *Percolator {
	if err := p.Open(row, col); err != nil {
		t.Fatal(err)
	}
	return p
}

func new(t *testing.T, n int) *Percolator {
	p, err := New(n)
	if err != nil {
		t.Fatal(err)
	}
	return p
}

func TestPercolation(t *testing.T) {
	tests := []struct {
		p *Percolator
		e []expectation
	}{
		{
			// test base case
			p: new(t, 100).build(t, 9, 5).build(t, 1, 2).build(t, 1, 5),
			e: []expectation{open(9, 5, true), open(1, 2, true), open(1, 5, true), open(1, 4, false)},
		},
		{
			// test full
			p: new(t, 16).build(t, 0, 1).build(t, 1, 1).build(t, 2, 1).build(t, 2, 2),
			e: []expectation{full(2, 2, true), full(2, 3, false), full(1, 1, true), full(2, 3, false)},
		},
		{
			// test percolates true
			p: new(t, 16).build(t, 0, 1).build(t, 1, 1).build(t, 2, 1).build(t, 2, 2).build(t, 2, 3).build(t, 3, 3),
			e: []expectation{percolates(true)},
		},
		{
			// test percolates false
			p: new(t, 16).build(t, 0, 1).build(t, 2, 1).build(t, 2, 2).build(t, 2, 3).build(t, 3, 3),
			e: []expectation{percolates(false)},
		},
		{
			// test open sites
			p: new(t, 16).build(t, 0, 1).build(t, 2, 1).build(t, 2, 2).build(t, 2, 3).build(t, 3, 3),
			e: []expectation{openSites(5)},
		},
	}

	for _, test := range tests {
		for _, e := range test.e {
			e(t, test.p)
		}
	}
}
