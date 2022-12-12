package main

import "fmt"

func main() {
	var L, R int
	fmt.Scan(&L, &R)
	fmt.Println("atcoder"[L-1 : R])
}
