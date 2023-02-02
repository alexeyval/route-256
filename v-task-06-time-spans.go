package main

import (
	"fmt"
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
	var t int
	fmt.Scan(&t)

	result := make([]string, t)
	for count := 0; count < t; count++ {
		result[count] = "YES"
		times := TimeSEs{}
		var n int
		fmt.Scan(&n)
		for line := 0; line < n; line++ {
			var s string
			fmt.Scan(&s)

			var dates [2]time.Time
			if !ParseDates(&dates, s) || dates[0].After(dates[1]) {
				result[count] = "NO"
			}
			times = append(times, timeSE{start: dates[0], end: dates[1]})
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
	}

	fmt.Println(strings.Join(result, "\n"))
}

func ParseDates(dates *[2]time.Time, lineDates string) bool {
	date := strings.Split(lineDates, "-")
	for i := range date {
		var err error
		dates[i], err = time.Parse("15:04:05", date[i])
		return err == nil
	}
	return true
}
