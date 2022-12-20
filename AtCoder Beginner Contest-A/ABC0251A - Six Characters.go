package main

¬	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	for i := 0; i < 6/len(s); i++ {
		fmt.Print(s)
	}
}

/*問題のURL
https://atcoder.jp/contests/abc251/tasks/abc251_a
*/
