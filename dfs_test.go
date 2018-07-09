package relax

import "testing"

func TestDFS(t *testing.T) {
	g := graph{
		1: []edge{
			{to: 2, w: 3},
		},
		2: []edge{
			{to: 4, w: 1},
		},
		3: []edge{
			{to: 1, w: 5},
			{to: 2, w: 2},
		},
		4: []edge{
			{to: 1, w: 4},
		},
	}
	dfs(g)
}
