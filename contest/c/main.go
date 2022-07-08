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

	var n, q int
	fmt.Fscan(in, &n, &q)

	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[i+1] = 0
	}
	mGlobalCounter := 0
	mCounter := 0
	for i := 0; i < q; i++ {
		var t, id int
		fmt.Fscan(in, &t, &id)

		if t == 1 {
			mCounter++
			if id == 0 {
				mGlobalCounter = mCounter
			} else {
				m[id] = mCounter
			}
		}
		if t == 2 {
			if m[id] < mGlobalCounter {
				fmt.Fprintln(out, mGlobalCounter)
				continue
			}

			fmt.Fprintln(out, m[id])
		}
	}
}
