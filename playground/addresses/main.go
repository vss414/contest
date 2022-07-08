package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func getIndex(addresses []string, phone string) int {
	for i, v := range addresses {
		if v == phone {
			return i
		}
	}

	return -1
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

		a := make(map[string][]string)
		for j := 0; j < n; j++ {
			var name, phone string
			fmt.Fscan(in, &name, &phone)

			index := getIndex(a[name], phone)
			if index != -1 {
				a[name] = append(a[name][:index], a[name][index+1:]...)
			}

			var tmp []string
			a[name] = append(append(tmp, phone), a[name]...)

			if len(a[name]) == 6 {
				a[name] = a[name][:5]
			}
		}

		keys := make([]string, 0, len(a))
		for k := range a {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			fmt.Fprintln(out, k+":", len(a[k]), strings.Join(a[k], " "))
		}

		fmt.Fprintln(out)
	}
}
