package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	hh := k / 60
	mm := k % 60
	fmt.Printf("%d:%02d", 21+hh, mm)
}
