package relax

import (
	"fmt"
)

type (
	edge struct {
		to int
		w  float32
	}

	graph map[int][]edge
)

func dfs(g graph) {
	enter := make(map[int]int)     // holds ts of vertex entering
	left := make(map[int]int)      // holds ts of vertex leaving
	lifo := make([]int, 0, len(g)) // lifo to hold vertexes in progress
	last := make(map[int]int)      // index of the last non watched linked vertex
	ts := 0                        // timestamp

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

				// if linked vertex was not visited yet
				if enter[e.to] == 0 {
					// add next vertex to the lifo and continue
					ts++
					enter[e.to] = ts
					lifo = append(lifo, e.to)
					last[cur] = i + 1

					fmt.Printf("found tree edge from %d to %d with cost %.2f\n", cur, e.to, e.w)
					fmt.Printf("entered %d at %d\n", e.to, enter[e.to])
					continue deep
				}

				// found vertex is not in the lifo!
				if left[e.to] == 0 {
					// if both vertexes are in the lifo
					// we have found the back edge, aka cycle
					fmt.Printf("found back edge from %d to %d with cost %.2f\n", cur, e.to, e.w)
					fmt.Printf("cycle is %v\n", lifo)
				} else if enter[cur] < enter[e.to] {
					// if the linked vertex was entered after the current vertex
					// we have found the forward edge
					fmt.Printf("found forward edge from %d to %d with cost %.2f\n", cur, e.to, e.w)
				} else {
					// if the linked vertex was entered before the current vertex
					// we have found the cross edge
					fmt.Printf("found cross edge from %d to %d with cost %.2f\n", cur, e.to, e.w)
				}
			}

			// we have visited all linked vertexes by now
			// leave the current too
			ts++
			left[cur] = ts
			lifo = lifo[:len(lifo)-1]
			fmt.Printf("left %d at %d\n", cur, left[cur])
		}
	}
}
