package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func topologicalSortUtil(m map[string][]string, v string, visited *map[string]bool, stack *[]string) {
	(*visited)[v] = true

	for i := 0; i < len(m[v]); i++ {
		if !(*visited)[m[v][i]] {
			topologicalSortUtil(m, m[v][i], visited, stack)
		}
	}

	*stack = append(*stack, v)
}

func topologicalSort(m map[string][]string, visited map[string]bool) []string {
	var stack []string

	for i, v := range visited {
		if !v {
			topologicalSortUtil(m, i, &visited, &stack)
		}
	}

	return stack
}

func getDependencies(old map[string]bool, d map[string]bool, m map[string][]string, module string) map[string]bool {
	for _, v := range m[module] {
		_, ok1 := d[v]
		_, ok2 := old[v]
		if !ok1 && !ok2 {
			d[v] = true
		}
		d = getDependencies(old, d, m, v)
	}

	return d
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		m := make(map[string][]string)
		visited := make(map[string]bool)

		for j := 0; j < n; j++ {
			scanner.Scan()
			p := strings.Split(scanner.Text(), " ")
			name := strings.Trim(p[0], ":")

			m[name] = p[1:]
			visited[name] = false
		}

		sorted := topologicalSort(m, visited)

		scanner.Scan()
		q, _ := strconv.Atoi(scanner.Text())

		old := make(map[string]bool)
		d := make(map[string]bool)
		for j := 0; j < q; j++ {
			scanner.Scan()
			module := scanner.Text()
			if _, ok := old[module]; !ok {
				d[module] = true
			}
			d = getDependencies(old, d, m, module)

			fmt.Print(len(d), " ")
			ind := 0
			for _, v := range sorted {
				if _, ok := d[v]; ok {
					fmt.Print(v, " ")
					ind++
					if ind >= len(d) {
						break
					}
				}
			}
			for k, v := range d {
				old[k] = v
			}
			d = map[string]bool{}
			fmt.Println()
		}
		fmt.Println()
	}
}
