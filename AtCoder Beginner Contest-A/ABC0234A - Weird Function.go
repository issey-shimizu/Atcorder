package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	fmt.Println(f(f(f(t)+t) + f(f(t))))
}

func f(x int) int {
	return x*x + 2*x + 3
}

/*問題文のURL
https://atcoder.jp/contests/abc234/tasks/abc234_a
*/
