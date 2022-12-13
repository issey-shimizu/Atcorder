package main

import "fmt"

func main() {
	var N, M, X, T, D int
	fmt.Scan(&N, &M, &X, &T, &D)
	if M >= X {
		fmt.Println(T)
	} else {
		fmt.Println(T - D*(X-M))
	}
}
