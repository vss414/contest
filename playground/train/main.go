package main

import (
	"bufio"
	"fmt"
	"os"
)

type answer string

const (
	yes answer = "SUCCESS"
	no  answer = "FAIL"
)

func check(in *bufio.Reader, out *bufio.Writer, train *[]bool, n int) {
	var query, p int
	fmt.Fscan(in, &query)

	switch query {
	case 1:
		fmt.Fscan(in, &p)
		if v := (*train)[p]; !v {
			(*train)[p] = true
			fmt.Fprintln(out, yes)
			return
		}
	case 2:
		fmt.Fscan(in, &p)
		if v := (*train)[p]; v {
			(*train)[p] = false
			fmt.Fprintln(out, yes)
			return
		}
	case 3:
		for i := 1; i < 2*n+1; i += 2 {
			if !(*train)[i] && !(*train)[i+1] {
				(*train)[i] = true
				(*train)[i+1] = true
				fmt.Fprintf(out, "%s %d-%d\n", yes, i, i+1)
				return
			}
		}
	}

	fmt.Fprintln(out, no)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n, q int
		fmt.Fscan(in, &n, &q)

		train := make([]bool, 2*n+1)
		for j := 0; j < q; j++ {
			check(in, out, &train, n)
		}
		fmt.Fprintln(out)
	}
}
