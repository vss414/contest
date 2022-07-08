package main

import (
	"bufio"
	"fmt"
	"os"
)

type matrix struct {
	matrix          [][]int
	sortIndex, n, m int
}

func (m matrix) print(out *bufio.Writer) {
	for i := 0; i < m.n; i++ {
		for j := 0; j < m.m; j++ {
			fmt.Fprint(out, m.matrix[i][j])
			if j < m.m-1 {
				fmt.Fprint(out, " ")
			}
		}
		fmt.Fprintln(out)
	}
	fmt.Fprintln(out)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for l := 0; l < t; l++ {
		var n, m int
		fmt.Fscan(in, &n, &m)

		matrix := matrix{
			matrix:    make([][]int, n),
			sortIndex: -1,
			n:         n,
			m:         m,
		}
		for i := range matrix.matrix {
			matrix.matrix[i] = make([]int, m)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Fscan(in, &matrix.matrix[i][j])
			}
		}

		var k int
		fmt.Fscan(in, &k)
		for i := 0; i < k; i++ {
			c := matrix.sortIndex
			fmt.Fscan(in, &matrix.sortIndex)

			matrix.sortIndex -= 1

			if c != matrix.sortIndex {
				var s1 = 1
				for s1 < matrix.n {
					var j = s1
					for j >= 1 && matrix.matrix[j][matrix.sortIndex] < matrix.matrix[j-1][matrix.sortIndex] {
						matrix.matrix[j], matrix.matrix[j-1] = matrix.matrix[j-1], matrix.matrix[j]

						j--
					}

					s1++
				}
			}
		}
		matrix.print(out)
	}
}
