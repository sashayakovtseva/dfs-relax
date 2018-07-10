package relax

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithoutCycles_simple(t *testing.T) {
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
	require.Equal(t, expect, WithoutCycles(g))
}

func TestWithoutCycles_twoCycles(t *testing.T) {
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
	require.Equal(t, expect, WithoutCycles(g))
}

func TestWithoutCycles_majorCycle(t *testing.T) {
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
			{To: 1, W: 5},
			{To: 3, W: 1},
		},
		5: Arrows{
			{To: 4, W: 9},
			{To: 1, W: 6},
		},
	}
	expect := Graph{
		1: Arrows{
			{To: 2, W: 1},
		},
		2: Arrows{},
		3: Arrows{
			{To: 2, W: 7},
		},
		4: Arrows{
			{To: 1, W: 3},
			{To: 3, W: 1},
		},
		5: Arrows{
			{To: 4, W: 9},
			{To: 1, W: 6},
		},
	}
	require.Equal(t, expect, WithoutCycles(g))
}

func TestWithoutCycles_mutual(t *testing.T) {
	g := Graph{
		1: Arrows{
			{To: 2, W: 6},
		},
		2: Arrows{
			{To: 1, W: 8},
			{To: 3, W: 9},
			{To: 4, W: 7},
		},
		3: Arrows{
			{To: 1, W: 5},
		},
		4: Arrows{},
	}
	expect := Graph{
		1: Arrows{},
		2: Arrows{
			{To: 1, W: 2},
			{To: 3, W: 9},
			{To: 4, W: 7},
		},
		3: Arrows{
			{To: 1, W: 5},
		},
		4: Arrows{},
	}
	require.Equal(t, expect, WithoutCycles(g))
}

func TestWithoutCycles_mutualWithCycle(t *testing.T) {
	g := Graph{
		1: Arrows{
			{To: 2, W: 8},
		},
		2: Arrows{
			{To: 1, W: 6},
			{To: 3, W: 9},
			{To: 4, W: 7},
		},
		3: Arrows{
			{To: 1, W: 5},
		},
		4: Arrows{},
	}
	expect := Graph{
		1: Arrows{},
		2: Arrows{
			{To: 3, W: 7},
			{To: 4, W: 7},
		},
		3: Arrows{
			{To: 1, W: 3},
		},
		4: Arrows{},
	}
	require.Equal(t, expect, WithoutCycles(g))
}
