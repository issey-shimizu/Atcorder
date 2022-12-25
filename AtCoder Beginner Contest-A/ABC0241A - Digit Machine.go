package main

import "fmt"

func main() {
	arr := [10]int{}
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr[arr[arr[0]]])
}

/*問題文のURL
https://atcoder.jp/contests/abc241/tasks/abc241_a
*/
