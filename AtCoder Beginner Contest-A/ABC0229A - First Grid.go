package main

import "fmt"

func main() {
	var s, t, ans string
	fmt.Scan(&s, &t)

	ans = "Yes"
	if s == "#." && t == ".#" || s == ".#" && t == "#." {
		ans = "No"
	}
	fmt.Println(ans)
}

/*問題文のURL
https://atcoder.jp/contests/abc229/tasks/abc229_a
*/
