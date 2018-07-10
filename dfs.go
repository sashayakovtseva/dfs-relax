package relax

import (
	"fmt"
)

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

func dfs(g Graph) {
	enter := make(map[Vertex]int)     // holds ts of vertex entering
	left := make(map[Vertex]int)      // holds ts of vertex leaving
	lifo := make(Vertices, 0, len(g)) // lifo to hold vertexes in progress
	last := make(map[Vertex]int)      // index of the last non watched linked vertex
	ts := 0                           // timestamp

	// while we have non visited vertexes
	for v := range g {
		if enter[v] != 0 {
			continue
		}

		ts++
		enter[v] = ts
		last[v] = 0
		lifo = append(lifo, v)
		fmt.Printf("entered %d at %d\n", v, enter[v])

	deep:
		// until lifo is not empty
		for len(lifo) != 0 {
			// get the top vetex in the lifo
			cur := lifo[len(lifo)-1]

			// for each linked vertex
			for i := last[cur]; i < len(g[cur]); i++ {
				e := g[cur][i]
				last[cur] = i + 1

				// if linked vertex is not visited yet
				if enter[e.To] == 0 {
					// add next vertex to the lifo and continue
					ts++
					enter[e.To] = ts
					lifo = append(lifo, e.To)

					fmt.Printf("found tree edge from %d to %d with cost %.2f\n", cur, e.To, e.W)
					fmt.Printf("entered %d at %d\n", e.To, enter[e.To])
					continue deep
				}

				if left[e.To] == 0 {
					// if both vertices are in the lifo
					// we have found the back edge, aka cycle
					fmt.Printf("found back edge from %d to %d with cost %.2f\n", cur, e.To, e.W)
					cycle := extractCycle(e.To, lifo) // find the cycle vertices
					relaxCycle(g, cycle)              // modify graph
					// clear data for vertices in the cycle as if we never actualy visited them
					for _, c := range cycle {
						last[c] = 0
						enter[c] = 0
					}
					// pop lifo to restart dfs
					lifo = lifo[:len(lifo)-len(cycle)]
					if len(lifo) != 0 {
						last[lifo[len(lifo)-1]] = 0
					}

					fmt.Printf("restarting dfs on modified graph\n")
					continue deep
				} else if enter[cur] < enter[e.To] {
					// if the linked vertex was entered after the current vertex
					// we have found the forward edge
					fmt.Printf("found forward edge from %d to %d with cost %.2f\n", cur, e.To, e.W)
				} else {
					// if the linked vertex was entered before the current vertex
					// we have found the cross edge
					fmt.Printf("found cross edge from %d to %d with cost %.2f\n", cur, e.To, e.W)
				}
			}

			// we have visited all linked vertices by now
			// leave the current too
			ts++
			left[cur] = ts
			lifo = lifo[:len(lifo)-1]
			fmt.Printf("left %d at %d\n", cur, left[cur])
		}
	}
}

func extractCycle(to Vertex, lifo Vertices) Vertices {
	var cycle []Vertex
	for i := len(lifo) - 1; i >= 0; i-- {
		if lifo[i] == to {
			cycle = append(cycle, lifo[i:]...)
			break
		}
	}
	fmt.Printf("cycle is %v\n", cycle)
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

	fmt.Printf("cycle throughput is %.2f\n", min)
	fmt.Printf("%+v\n", g)

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

	fmt.Printf("%+v\n", g)
}
