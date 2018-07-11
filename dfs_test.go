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
		"slava": Arrows{
			{To: "misha", W: 6},
		},
		"misha": Arrows{
			{To: "slava", W: 4},
			{To: "roma", W: 1},
		},
		"roma": Arrows{
			{To: "misha", W: 8},
			{To: "slava", W: 12},
			{To: "sasha", W: 5},
		},
		"sasha": Arrows{
			{To: "roma", W: 11},
		},
	}
	expect := Graph{
		"slava": Arrows{
			{To: "misha", W: 2},
		},
		"misha": Arrows{},
		"roma": Arrows{
			{To: "misha", W: 7},
			{To: "slava", W: 12},
		},
		"sasha": Arrows{
			{To: "roma", W: 6},
		},
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

func TestWithoutCycles_removeCycles(t *testing.T) {
	tt := []struct {
		name   string
		g      Graph
		expect Graph
	}{
		{
			name: "decrease on tree",
			g: Graph{
				"nik": Arrows{
					{To: "slava", W: 10},
					{To: "sasha", W: 9},
				},
				"slava": Arrows{
					{To: "artem", W: 8},
					{To: "roma", W: 8},
				},
				"artem": Arrows{
					{To: "sasha", W: 6},
					{To: "nik", W: 5},
				},
				"roma": Arrows{
					{To: "nik", W: 4},
				},
			},
			expect: Graph{
				"nik": Arrows{
					{To: "slava", W: 1},
					{To: "sasha", W: 9},
				},
				"slava": Arrows{
					{To: "artem", W: 3},
					{To: "roma", W: 4},
				},
				"artem": Arrows{
					{To: "sasha", W: 6},
				},
				"roma": Arrows{},
			},
		},
		{
			name: "should revisit",
			g: Graph{
				"slava": Arrows{
					{To: "misha", W: 24},
					{To: "roma", W: 34},
				},
				"roma": Arrows{
					{To: "misha", W: 2},
					{To: "slava", W: 19},
				},
				"misha": Arrows{
					{To: "slava", W: 26},
				},
			},
			expect: Graph{
				"slava": Arrows{
					{To: "roma", W: 13},
				},
				"roma":  Arrows{},
				"misha": Arrows{},
			},
		},
		{
			name: "stack overflow",
			g: Graph{
				"misha": Arrows{
					{To: "nik", W: 10},
				},
				"nik": Arrows{
					{To: "slava", W: 40},
				},
				"slava": Arrows{
					{To: "roman", W: 40},
				},
				"roman": Arrows{
					{To: "nik", W: 40},
				},
			},
			expect: Graph{
				"nik":   Arrows{},
				"slava": Arrows{},
				"roman": Arrows{},
				"misha": Arrows{
					{To: "nik", W: 10},
				},
			},
		},
		{
			name: "multiple operations",
			g: Graph{
				"misha": Arrows{
					{To: "nik", W: 10},
				},
				"nik": Arrows{
					{To: "misha", W: 5},
				},
			},
			expect: Graph{
				"nik": Arrows{},
				"misha": Arrows{
					{To: "nik", W: 5},
				},
			},
		},
		{
			name: "big test",
			g: Graph{
				"nik": Arrows{
					{To: "slava", W: 40},
				},
				"slava": Arrows{
					{To: "artyom", W: 10},
					{To: "roman", W: 40},
				},
				"roman": Arrows{
					{To: "sasha", W: 18},
					{To: "nik", W: 30},
				},
				"misha": Arrows{
					{To: "nik", W: 15},
					{To: "artyom", W: 5},
				},
				"sasha": Arrows{
					{To: "slava", W: 18},
					{To: "artyom", W: 10},
				},
				"artyom": Arrows{
					{To: "misha", W: 30},
				},
			},
			expect: Graph{
				"nik": Arrows{
					{To: "slava", W: 8},
				},
				"misha": Arrows{
					{To: "nik", W: 5},
				},
				"roman": Arrows{
					{To: "nik", W: 8},
				},
				"slava": Arrows{},
				"artyom": Arrows{
					{To: "misha", W: 15},
				},
				"sasha": Arrows{
					{To: "artyom", W: 10},
				},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expect, WithoutCycles(tc.g))
		})
	}
}
