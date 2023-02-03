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

type Fail bool

func (fail Fail) String() string {
	switch fail {
	case true:
		return "NO"
	default:
		return "YES"
	}
}

type TimeSEs []timeSE

func (t TimeSEs) Less(i, j int) bool { return t[i].start.Before(t[j].start) }
func (t TimeSEs) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t TimeSEs) Len() int           { return len(t) }

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(reader, &t)

	for count := 0; count < t; count++ {
		fail := false
		times := TimeSEs{}
		var n int
		fmt.Fscan(reader, &n)
		for line := 0; line < n; line++ {
			var s string
			fmt.Fscan(reader, &s)

			var dates timeSE
			fail = fail || (!ParseDates(&dates, s) || dates.start.After(dates.end))
			times = append(times, dates)
		}

		sort.Sort(times)
		for i := 1; i < len(times); i++ {
			start := times[i].start
			endPrev := times[i-1].end
			fail = fail || start.Before(endPrev) || start.Equal(endPrev)
			if fail {
				break
			}
		}
		fmt.Println(Fail(fail))
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
