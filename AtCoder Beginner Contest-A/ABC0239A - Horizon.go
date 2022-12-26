package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)
	fmt.Println(math.Sqrt(x * (x + 12800000)))
}

/*問題文のURL
https://atcoder.jp/contests/abc239/tasks/abc239_a
*/
