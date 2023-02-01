package main

import (
	"fmt"
	"sort"
	"strings"
)

type Table struct {
	row []int
}

func NewTable(n, m int) ByRow {
	table := make([]Table, n)
	for i := range table {
		row := make([]int, m)
		for j := range row {
			fmt.Scan(&row[j])
		}
		table[i].row = row
	}
	return table
}

func (t Table) String() string {
	return strings.Trim(fmt.Sprint(t.row), "[]")
}

type ByRow []Table

func (a ByRow) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRow) Less(i, j int) bool { return a[i].row[c-1] < a[j].row[c-1] }
func (a ByRow) Len() int           { return len(a) }

var c int

func main() {
	var t int
	fmt.Scan(&t)

	for count := 0; count < t; count++ {
		var n, m int
		fmt.Scan(&n, &m)
		table := NewTable(n, m)

		var k int
		fmt.Scan(&k)
		for i := 0; i < k; i++ {
			fmt.Scan(&c)
			sort.Sort(table)
		}

		for _, tt := range table {
			fmt.Println(tt)
		}
		fmt.Println()
	}
}
