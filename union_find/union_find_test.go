package unionfind

import (
	"fmt"
	"testing"
)

func graphStr(graph []int, width int) (str string) {
	for i := range graph {
		str += fmt.Sprintf("%d", graph[i])
		if (i+1)%width == 0 {
			str += "\n"
			continue
		}
		str += " "
	}
	return
}

type ufCons func(int) UF

type expectation func(*testing.T, UF)

func connected(p, q int, expected bool) func(*testing.T, UF) {
	return func(t *testing.T, impl UF) {
		if impl.Connected(p, q) != expected {
			switch x := impl.(type) {
			case *WeightedQuickUnion:
				t.Fatalf("Expected UF:%s Connected:%t p:%d q:%d\n ids:%s\n size:%s",
					impl, expected, p, q,
					graphStr(x.ids, 5),
					graphStr(x.sz, len(x.sz)),
				)
			default:
				t.Fatalf("Expected UF:%s Connected:%t p:%d q:%d", impl, expected, p, q)
			}
		}
	}
}

type test struct {
	graph        [][]int
	expectations []expectation
}

func TestUF(t *testing.T) {

	table := []test{
		test{
			graph: [][]int{
				[]int{0, 0, 2, 3, 4},
				[]int{5, 6, 7, 8, 4},
				[]int{10, 11, 12, 4, 14},
			},
			expectations: []expectation{
				connected(0, 1, true),
				connected(0, 2, false),
				connected(0, 5, false),
				connected(4, 9, true),
				connected(4, 13, true),
				connected(9, 13, true),
			},
		},
	}

	for _, constructor := range []ufCons{NewWeightedQuickUnion} {
		for _, test := range table {
			impl := constructor(len(test.graph) * len(test.graph[0]))

			// init UF graph
			for i := range test.graph {
				for j := range test.graph[i] {
					id := j + (len(test.graph[i]) * i)
					impl.Union(id, test.graph[i][j])
				}
			}

			for _, e := range test.expectations {
				e(t, impl)
			}
		}
	}
}
