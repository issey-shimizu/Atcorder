package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n > 41 {
		n++
	}
	fmt.Printf("AGC%03d\n", n)
}

/*問題文のURL
https://atcoder.jp/contests/abc230/tasks/abc230_a
*/
