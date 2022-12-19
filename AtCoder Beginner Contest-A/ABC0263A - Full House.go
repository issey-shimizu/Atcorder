package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	s := []int{}
	for i := 0; i < 5; i++ {
		fmt.Scan(&n)
		s = append(s, n)
	}
	sort.Ints(s)
	if s[0] == s[1] && s[3] == s[4] {
		if s[2] == s[0] || s[2] == s[4] {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

/* 問題文のURL
https://atcoder.jp/contests/abc263/tasks/abc263_a
*/
