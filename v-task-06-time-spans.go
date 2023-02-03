package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type timeSE struct {
	start time.Time
	end   time.Time
}

type TimeSEs []timeSE

func (s TimeSEs) Less(i, j int) bool { return s[i].start.Before(s[j].start) }
func (s TimeSEs) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s TimeSEs) Len() int           { return len(s) }

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(reader, &t)

	result := make([]string, t)
	for count := 0; count < t; count++ {
		result[count] = "YES"
		times := TimeSEs{}
		var n int
		fmt.Fscan(reader, &n)
		for line := 0; line < n; line++ {
			var s string
			fmt.Fscan(reader, &s)

			var dates timeSE
			if !ParseDates(&dates, s) || dates.start.After(dates.end) {
				result[count] = "NO"
			}
			times = append(times, dates)
		}

		sort.Sort(times)
		for i := 1; i < len(times); i++ {
			start := times[i].start
			endPrev := times[i-1].end
			if start.Before(endPrev) || start.Equal(endPrev) {
				result[count] = "NO"
				break
			}
		}
		fmt.Println(result[count])
	}
}

func ParseDates(t *timeSE, lineDates string) bool {
	dates := strings.Split(lineDates, "-")
	var d [2]time.Time
	for i := range dates {
		var err error
		d[i], err = time.Parse("15:04:05", dates[i])
		if err != nil {
			return false
		}
	}
	t.start = d[0]
	t.end = d[1]
	return true
}
