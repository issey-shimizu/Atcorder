package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	Takahashi := a*60 + b
	Aoki := c*60 + d

	if Takahashi <= Aoki {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}
