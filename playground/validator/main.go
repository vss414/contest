package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type answer string

const (
	yes answer = "yes"
	no  answer = "no"
)

func validate(login string) answer {
	if len(login) < 2 || len(login) > 24 {
		return no
	}

	if matched, _ := regexp.MatchString(`^[\w\-]+$`, login); !matched {
		return no
	}

	if login[0] == '-' {
		return no
	}

	return yes
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

		logins := make(map[string]bool)
		for j := 0; j < n; j++ {
			var login string
			fmt.Fscan(in, &login)
			login = strings.ToLower(login)

			if _, ok := logins[login]; ok {
				fmt.Fprintln(out, no)
				continue
			}

			res := validate(login)
			if res == yes {
				logins[login] = true
			}

			fmt.Fprintln(out, res)
		}
		fmt.Fprintln(out)
	}
}
