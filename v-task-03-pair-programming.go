package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &a[j])
		}
		for j := 0; j < n; j++ {
			if a[j] == 0 {
				continue
			}
			suitable := -1
			for k := j + 1; k < n; k++ {
				if a[k] == 0 {
					continue
				}
				if suitable == -1 {
					suitable = k
				}
				if diff(a[k], a[j]) < diff(a[suitable], a[j]) {
					suitable = k
				}
			}
			a[j] = 0
			a[suitable] = 0
			fmt.Println(j+1, suitable+1)
		}
		fmt.Println()
	}
}

func diff(a, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}
