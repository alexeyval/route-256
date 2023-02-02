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
	return ByRow{tables: table, len: n}
}

func (t Table) String() string {
	return strings.Trim(fmt.Sprint(t.row), "[]")
}

type ByRow struct {
	tables []Table
	c      int
	len    int
}

func (a ByRow) Swap(i, j int)      { a.tables[i], a.tables[j] = a.tables[j], a.tables[i] }
func (a ByRow) Less(i, j int) bool { return a.tables[i].row[a.c-1] < a.tables[j].row[a.c-1] }
func (a ByRow) Len() int           { return a.len }

func main() {
	var t int
	fmt.Scan(&t)

	for count := 0; count < t; count++ {
		var n, m int
		fmt.Scan(&n, &m)
		byRow := NewTable(n, m)

		var k int
		fmt.Scan(&k)
		for i := 0; i < k; i++ {
			fmt.Scan(&byRow.c)
			sort.Sort(byRow)
		}

		for _, row := range byRow.tables {
			fmt.Println(row)
		}
		fmt.Println()
	}
}
