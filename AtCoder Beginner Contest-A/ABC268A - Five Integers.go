package main

import "fmt"

func main() {
	s := make(map[int]int)
	for i := 0; i < 5; i++ {
		var x int
		fmt.Scan(&x)
		s[x]++
	}
	fmt.Println(len(s))
}
