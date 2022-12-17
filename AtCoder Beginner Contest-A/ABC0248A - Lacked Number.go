package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ans := 45

	for i, _ := range s {
		ans = ans - int(s[i]-'0')
	}
	fmt.Println(ans)

}

/*　問題文のURL
https://atcoder.jp/contests/abc248/submissions/37314716
*/
