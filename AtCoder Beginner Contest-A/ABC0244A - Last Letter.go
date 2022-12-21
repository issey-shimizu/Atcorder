package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	fmt.Println(string(s[n-1]))
}

/*問題文のURL
https://atcoder.jp/contests/abc244/tasks/abc244_a
*/
