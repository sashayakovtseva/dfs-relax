package relax

type (
	// Vertex holds a signle vertex
	Vertex interface{}

	// Arrow represents a single arrow in a directed graph.
	Arrow struct {
		// To holds id of a vertext which arrow goes to.
		To Vertex
		// W is a weight of an arrow.
		W float32
	}

	// Arrows is a set of arrows from the particular vertex.
	Arrows []Arrow

	// Vertices is a set of graph vetices.
	Vertices []Vertex

	// Graph is graph.
	Graph map[Vertex]Arrows
)

// WithoutCycles returns graph with the same vertices as the original
// but without any cycles. Whenever cycle is found it's arrows are relaxed
// by the minimum value in the cycle found. WithoutCycles does not
// modify original graph.
func WithoutCycles(g Graph) Graph {
	c := make(Graph, len(g))
	for k, v := range g {
		c[k] = make(Arrows, len(v))
		copy(c[k], v)
	}
	dfs(c)
	return c
}

func dfs(g Graph) {
	enter := make(map[Vertex]int) // holds ts of vertex entering
	left := make(map[Vertex]int)  // holds ts of vertex leaving
	last := make(map[Vertex]int)  // index of the last non watched linked vertex
	ts := 0                       // timestamp

	var l lifo // lifo to hold vertexes in progress
	// while we have non visited vertexes
	for v := range g {
		if enter[v] != 0 {
			continue
		}

		ts++
		enter[v] = ts
		last[v] = 0
		l.Push(v)

	deep:
		// until lifo is not empty
		for !l.Empty() {
			// get the top vetex in the lifo
			cur := l.Top()

			// for each linked vertex
			for i := last[cur]; i < len(g[cur]); i++ {
				e := g[cur][i]
				last[cur] = i + 1

				// if linked vertex is not visited yet
				if enter[e.To] == 0 {
					// add next vertex to the lifo and continue
					ts++
					enter[e.To] = ts
					l.Push(e.To)
					continue deep
				}

				if left[e.To] == 0 {
					// if both vertices are in the lifo
					// we have found the back edge, aka cycle
					cycle := extractCycle(e.To, l) // find the cycle vertices
					relaxCycle(g, cycle)           // modify graph
					// clear data for vertices in the cycle as if we never actualy visited them
					for _, c := range cycle {
						last[c] = 0
						enter[c] = 0
						l.Pop()
					}
					if !l.Empty() {
						last[l.Top()] = 0
					}
					continue deep
				}
			}

			// we have visited all linked vertices by now
			// leave the current too
			ts++
			left[cur] = ts
			l.Pop()
		}
	}
}

func extractCycle(to Vertex, l lifo) Vertices {
	var cycle []Vertex
	for i := len(l) - 1; i >= 0; i-- {
		if l[i] == to {
			cycle = append(cycle, l[i:]...)
			break
		}
	}
	return cycle
}

func relaxCycle(g Graph, cycle Vertices) {
	var min float32

	// find minimum cost edge in the cycle
	for i := 0; i < len(cycle); i++ {
		cur := cycle[i]
		for _, e := range g[cur] {
			if e.To == cycle[(i+1)%len(cycle)] {
				if e.W < min || min == 0 {
					min = e.W
				}
				break
			}
		}
	}

	// relax all cycle edges according to found cost
	for i := 0; i < len(cycle); i++ {
		cur := cycle[i]
		for j := range g[cur] {
			if g[cur][j].To == cycle[(i+1)%len(cycle)] {
				g[cur][j].W -= min
				if g[cur][j].W == 0 {
					g[cur] = append(g[cur][:j], g[cur][j+1:]...)
				}
				break
			}
		}
	}
}
