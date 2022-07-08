package main

import (
	"fmt"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Printf("failed to read n")
	}

	for i := 0; i < n; i++ {
		var a, b int
		_, err := fmt.Scanf("%d %d", &a, &b)
		if err != nil {
			fmt.Printf("failed to read a and b")
			return
		}
		fmt.Println(a + b)
	}
}
