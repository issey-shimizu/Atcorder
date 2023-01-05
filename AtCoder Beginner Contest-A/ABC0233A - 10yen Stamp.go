package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	num := 0

	for x < y {
		num++
		x += 10
	}
	fmt.Println(num)
}

/*問題文のURL
https://atcoder.jp/contests/abc233/tasks/abc233_a
*/
