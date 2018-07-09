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
	visit := make(map[int]int)
	lifo := make([]int, 0, len(g))
	last := make(map[int]int)
	left := make(map[int]int)
	t := 0

	for k := range g {
		if visit[k] != 0 {
			continue
		}

		t++
		lifo = append(lifo, k)
		visit[k] = t
		last[k] = 0
		fmt.Printf("visited %d at %d\n", k, visit[k])

	deep:
		for len(lifo) != 0 {
			top := lifo[len(lifo)-1]
			for _, e := range g[top][last[top]:] {
				if visit[e.to] == 0 {
					fmt.Printf("found tree edge from %d to %d with cost %.2f\n", top, e.to, e.w)
					last[top]++
					lifo = append(lifo, e.to)
					t++
					visit[e.to] = t
					fmt.Printf("visited %d at %d\n", e.to, visit[e.to])
					continue deep
				}
				if left[e.to] == 0 {
					fmt.Printf("found back edge from %d to %d with cost %.2f\n", top, e.to, e.w)
				} else if visit[top] < visit[e.to] {
					fmt.Printf("found forward edge from %d to %d with cost %.2f\n", top, e.to, e.w)
				} else {
					fmt.Printf("found cross edge from %d to %d with cost %.2f\n", top, e.to, e.w)
				}
			}
			left[top] = t
			t++
			lifo = lifo[:len(lifo)-1]
			fmt.Printf("completed visiting %d at %d\n", top, left[top])
		}
	}
}
