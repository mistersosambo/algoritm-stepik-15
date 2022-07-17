package main

import "fmt"

func main() {
	var a, b int
	max := -10000
	fmt.Scan(&a)
	fmt.Scan(&b)
	for a <= b {
		if a%7 == 0 && a >= max {
			max = a
		}
		a++
	}
	if max != -10000 {
		fmt.Println(max)
	} else {
		fmt.Println("NO")
	}
}
