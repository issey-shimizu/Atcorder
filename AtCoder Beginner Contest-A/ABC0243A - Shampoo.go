package main

import (
	"fmt"
)

func main() {
	var v, a, b, c int
	fmt.Scan(&a, &b, &c, &d)
	v = v % (a + b + c)

	if v < a {
		fmt.Println("F")
	} else if v < a+b {
		fmt.Println("M")
	} else {
		fmt.Println("T")
	}
}

/*問題文のURL
https://atcoder.jp/contests/abc243/tasks/abc243_a
*/
