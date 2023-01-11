package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	m := make(map[string]int)
	ans := "Yes"
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		if strings.ContainsAny(s[:1], "HDCS") && strings.ContainsAny(s[1:], "A23456789TJQK") && m[s] == 0 {
			m[s]++
		} else {
			ans = "No"
		}
	}
	fmt.Println(ans)
}

/*問題文のURL
https://atcoder.jp/contests/abc277/tasks/abc277_b
*/
