package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type answer string

const (
	yes answer = "YES"
	no  answer = "NO"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)

		a := yes
		var starts []time.Time
		var ends []time.Time
		for j := 0; j < n; j++ {
			var p string
			fmt.Fscan(in, &p)

			if a == no {
				continue
			}

			s := strings.Split(p, "-")

			start, err := time.Parse("15:04:05", s[0])
			if err != nil {
				a = no
				continue
			}
			end, err := time.Parse("15:04:05", s[1])
			if err != nil {
				a = no
				continue
			}

			if end.Sub(start).Seconds() < 0 {
				a = no
				continue
			}

			for q := 0; q < len(starts); q++ {
				if (start.Sub(starts[q]).Seconds() >= 0 && start.Sub(ends[q]).Seconds() <= 0) ||
					(end.Sub(starts[q]).Seconds() >= 0 && end.Sub(ends[q]).Seconds() <= 0) ||
					(starts[q].Sub(start).Seconds() >= 0 && ends[q].Sub(end).Seconds() <= 0) {
					a = no
					break
				}
			}

			if a == yes {
				starts = append(starts, start)
				ends = append(ends, end)
			}
		}
		fmt.Fprintln(out, a)
	}
}
