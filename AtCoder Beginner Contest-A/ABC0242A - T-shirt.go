package main

import (
	"fmt"
)

func main() {
	var a, b, c, x, p float64
	fmt.Scan(&a, &b, &c, &x)

	if x <= a {
		p = 1

	} else if x <= b {
		p = c / (b - a)
	} else {
		p = 0
	}

	fmt.Println(p)
}

/*問題文のURL
https://atcoder.jp/contests/abc242/tasks/abc242_a
*/
