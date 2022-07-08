package main

import (
	"bufio"
	"fmt"
	"os"
)

type worker struct {
	i, a   int
	inPair bool
}

func remove(slice []worker, s int) []worker {
	return append(slice[:s], slice[s+1:]...)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)

		a := make([]int, n)
		w := make([]worker, n)
		for j := 0; j < n; j++ {
			var p int
			fmt.Fscan(in, &p)
			a[j] = p
			w[j] = worker{
				i:      j + 1,
				a:      p,
				inPair: false,
			}
		}

		for x := 0; x < n/2; x++ {
			first := w[0].i
			firstA := w[0].a

			min := 101
			minInd := 0
			for y := 1; y < len(w); y++ {
				tmp := Abs(w[y].a - firstA)
				if tmp < min {
					min = tmp
					minInd = y
				}
			}

			second := w[minInd].i

			w = remove(w, minInd)
			w = remove(w, 0)

			fmt.Fprintf(out, "%d %d\n", first, second)
		}

		fmt.Fprintln(out)
	}
}
