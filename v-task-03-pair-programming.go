package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		a := make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Scan(&a[j])
		}
		for j := 0; j < n; j++ {
			if a[j] == 0 {
				continue
			}
			suitable := -1
			for k := j + 1; k < n; k++ {
				if suitable == -1 {
					suitable = k
				}
				if diff(a[k], a[j]) < diff(a[j], a[suitable]) {
					suitable = k
				}
			}
			a[suitable] = 0
			fmt.Println(j+1, suitable+1)
		}
	}
}

func diff(a, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}
