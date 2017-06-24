package main

import "fmt"

func main() {
	var n int32
	fmt.Scanf("%d", &n)
	fmt.Println(factorial(n))
}
func factorial(n int32) int32 {

	switch n {
	case 1:
		return 1
	default:
		return n * factorial(n - 1);
	}
}
