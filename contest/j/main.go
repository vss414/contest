package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	d := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &d[i])
	}
	sort.Slice(d, func(i, j int) bool { return len(d[i]) > len(d[j]) })

	var q int
	fmt.Fscan(in, &q)

	for i := 0; i < q; i++ {
		var t string
		fmt.Fscan(in, &t)
		maxS := 0
		maxI := 0
		for j := 0; j < n; j++ {
			if t == d[j] {
				continue
			}
			if maxS >= len(d[j]) {
				break
			}
			suffix := 0

			l := len(t) - 1
			for s := len(d[j]) - 1; s >= 0; s-- {
				if l >= 0 && d[j][s] == t[l] {
					suffix++
					l--
				} else {
					break
				}
			}

			if maxS < suffix {
				maxS = suffix
				maxI = j
			}
		}

		fmt.Fprintln(out, d[maxI])
	}
}
