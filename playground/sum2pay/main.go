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

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)

		var sum int64
		items := make(map[int64]int64)
		for j := 0; j < n; j++ {
			var p int64
			fmt.Fscan(in, &p)

			if items[p] == 2 {
				items[p] = 0
			} else {
				sum += p
				items[p] += 1
			}
		}

		fmt.Fprintln(out, sum)
	}
}
