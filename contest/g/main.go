package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	f := make(map[int][]int)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)

		f[a] = append(f[a], b)
		f[b] = append(f[b], a)
	}

	for j := 1; j <= n; j++ {
		if len(f[j]) == 0 {
			fmt.Fprintln(out, 0)
			continue
		}

		maxSet := make([]int, 0)
		hash := make(map[int]bool)
		var inds []int

		for s := 0; s < len(f[j]); s++ {
			hash[f[j][s]] = true
		}

		for s := 1; s <= n; s++ {
			if j == s {
				continue
			}
			if _, found := hash[s]; found {
				continue
			}

			set := make([]int, 0)
			for k := 0; k < len(f[s]); k++ {
				el := f[s][k]
				if _, found := hash[el]; found {
					set = append(set, el)
				}
			}

			if len(set) == 0 {
				continue
			}

			if len(maxSet) < len(set) {
				maxSet, set = set, maxSet
				inds = []int{s}
			} else if len(maxSet) == len(set) {
				inds = append(inds, s)
			}
		}

		l := len(inds)

		if l == 0 {
			fmt.Fprint(out, 0)
		}

		for s := 0; s < l; s++ {
			fmt.Fprint(out, inds[s])
			if s != l {
				fmt.Fprint(out, " ")
			}
		}
		fmt.Fprintln(out)
	}
}
