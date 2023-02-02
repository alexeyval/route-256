package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	graph := map[int]map[int]bool{}

	for i := 0; i < M; i++ {
		var v1, v2 int
		fmt.Scan(&v1, &v2)
		for k := 0; k < 2; k++ {
			if graph[v1] == nil {
				graph[v1] = map[int]bool{}
			}
			graph[v1][v2] = false
			v1, v2 = v2, v1
		}
	}
	for k, v := range graph {
		fmt.Println(k, v)
	}
}
