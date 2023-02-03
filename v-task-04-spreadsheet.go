package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Table [][]int

func NewTable(n, m int) Table {
	table := make([][]int, n)
	for i := range table {
		row := make([]int, m)
		for j := range row {
			fmt.Fscan(reader, &row[j])
		}
		table[i] = row
	}
	return table
}

func (t Table) String() string {
	sTable := ""
	for _, row := range t {
		sTable += strings.Trim(fmt.Sprint(row), "[]") + "\n"
	}
	return sTable
}

type ByColumn struct {
	table Table
	c     int
	len   int
}

func (a ByColumn) Swap(i, j int)      { a.table[i], a.table[j] = a.table[j], a.table[i] }
func (a ByColumn) Less(i, j int) bool { return a.table[i][a.c-1] < a.table[j][a.c-1] }
func (a ByColumn) Len() int           { return a.len }

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(reader, &t)

	for count := 0; count < t; count++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		table := NewTable(n, m)

		var k int
		fmt.Fscan(reader, &k)
		for i := 0; i < k; i++ {
			byColumn := ByColumn{table: table, len: n}
			fmt.Fscan(reader, &byColumn.c)
			sort.Stable(byColumn)
		}

		fmt.Println(table)
	}
}
