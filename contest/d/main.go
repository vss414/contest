package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
)

func hasUppercase(s string) bool {
	matched, _ := regexp.MatchString(`(?:[A-Z])`, s)
	return matched
}

func hasLowercase(s string) bool {
	matched, _ := regexp.MatchString(`(?:[a-z])`, s)
	return matched
}

func hasVowel(s string) bool {
	matched, _ := regexp.MatchString(`(?i:[euioay])`, s)
	return matched
}

func hasСonsonant(s string) bool {
	matched, _ := regexp.MatchString(`(?i:[bcdfghjklmnpqrstvwxz])`, s)
	return matched
}

func hasDigit(s string) bool {
	matched, _ := regexp.MatchString(`^($|.*[0-9])`, s)
	return matched
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var p string
		fmt.Fscan(in, &p)

		a := []string{
			"",
			"a",
			"b",
			"A",
			"B",
			"bA",
			"Ba",
		}

		for j := 0; j < len(a); j++ {
			tmp := p + a[j]
			if hasСonsonant(tmp) && hasVowel(tmp) && hasLowercase(tmp) && hasUppercase(tmp) {
				if !hasDigit(tmp) {
					tmp += strconv.Itoa(rand.Intn(9))
				}

				fmt.Fprintln(out, tmp)
				break
			}
		}
	}
}
