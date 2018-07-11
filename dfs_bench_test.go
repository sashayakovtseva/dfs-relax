package relax

import "testing"

func BenchmarkWithoutCycles(b *testing.B) {
	g := Graph{
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
	}
	for i := 0; i < b.N; i++ {
		WithoutCycles(g)
	}
}

func BenchmarkWithoutCyclesNoCopy(b *testing.B) {
	g := Graph{
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
	}
	for i := 0; i < b.N; i++ {
		WithoutCyclesNoCopy(g)
	}
}
