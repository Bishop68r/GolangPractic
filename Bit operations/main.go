package main

import "fmt"

func main() {
	var a int = 0 & 1
	var b int = 0 | 1

	fmt.Println(a, b)
}
