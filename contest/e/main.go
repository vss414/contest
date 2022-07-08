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

		a := make(map[int]int, n)
		answer := "YES"
		for j := 0; j < n; j++ {
			var p int
			fmt.Fscan(in, &p)

			ind, ok := a[p]
			if ok && ind+1 != j {
				answer = "NO"
			}
			a[p] = j
		}
		fmt.Fprintln(out, answer)
	}
}
