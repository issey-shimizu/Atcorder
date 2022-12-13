package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	m := make(map[string]int)

	m[s[0:1]]++
	m[s[1:2]]++
	m[s[2:3]]++

	for i, v := range m {
		if v == 1 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
