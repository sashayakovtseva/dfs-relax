package relax

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDFS(t *testing.T) {
	g := Graph{
		1: Arrows{
			{To: 2, W: 3},
		},
		2: Arrows{
			{To: 4, W: 1},
		},
		3: Arrows{
			{To: 1, W: 5},
			{To: 2, W: 2},
		},
		4: Arrows{
			{To: 1, W: 4},
		},
	}
	expect := Graph{
		1: Arrows{
			{To: 2, W: 2},
		},
		2: Arrows{},
		3: Arrows{
			{To: 1, W: 5},
			{To: 2, W: 2},
		},
		4: Arrows{
			{To: 1, W: 3},
		},
	}
	dfs(g)
	require.Equal(t, expect, g)
}

func TestDFS2(t *testing.T) {
	g := Graph{
		1: Arrows{
			{To: 2, W: 3},
		},
		2: Arrows{
			{To: 4, W: 2},
		},
		3: Arrows{
			{To: 2, W: 7},
		},
		4: Arrows{
			{To: 3, W: 1},
			{To: 1, W: 5},
		},
		5: Arrows{
			{To: 4, W: 9},
			{To: 1, W: 6},
		},
	}
	expect := Graph{
		1: Arrows{
			{To: 2, W: 2},
		},
		2: Arrows{},
		3: Arrows{
			{To: 2, W: 6},
		},
		4: Arrows{
			{To: 1, W: 4},
		},
		5: Arrows{
			{To: 4, W: 9},
			{To: 1, W: 6},
		},
	}
	dfs(g)
	require.Equal(t, expect, g)
}
