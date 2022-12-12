package main

import "fmt"

func main() {
	var s string
	m := map[string]int{"Monday": 5, "Tuesday": 4, "Wednesday": 3, "Thursday": 2, "Friday": 1}
	fmt.Scan(&s)
	fmt.Println(m[s])
}
