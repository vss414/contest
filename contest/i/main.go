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

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]uint64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	e := make(map[int]uint64)
	var sum uint64 = 0
	for i := 0; i < m+1; i++ {
		var t, l uint64
		fmt.Fscan(in, &t, &l)

		p := -1
		for j := range a {
			if _, ok := e[j]; !ok || e[j] <= t {
				e[j] = t + l
				p = j
				break
			}
		}

		if p == -1 {
			continue
		}

		sum += l * a[p]
	}
	fmt.Fprintln(out, sum)
}
