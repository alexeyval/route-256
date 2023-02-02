package main

import (
	"fmt"
	"sort"
	"strings"
)

type Row struct {
	row []int
}

func NewTable(n, m int) []Row {
	table := make([]Row, n)
	for i := range table {
		row := make([]int, m)
		for j := range row {
			fmt.Scan(&row[j])
		}
		table[i].row = row
	}
	return table
}

func (t Row) String() string {
	return strings.Trim(fmt.Sprint(t.row), "[]")
}

type ByColumn struct {
	table []Row
	c     int
	len   int
}

func (a ByColumn) Swap(i, j int)      { a.table[i], a.table[j] = a.table[j], a.table[i] }
func (a ByColumn) Less(i, j int) bool { return a.table[i].row[a.c-1] < a.table[j].row[a.c-1] }
func (a ByColumn) Len() int           { return a.len }

func main() {
	var t int
	fmt.Scan(&t)

	for count := 0; count < t; count++ {
		var n, m int
		fmt.Scan(&n, &m)
		table := NewTable(n, m)
		byColumn := ByColumn{table: table, len: n}

		var k int
		fmt.Scan(&k)
		for i := 0; i < k; i++ {
			fmt.Scan(&byColumn.c)
			sort.Sort(byColumn)
		}

		for _, row := range table {
			fmt.Println(row)
		}
		fmt.Println()
	}
}
