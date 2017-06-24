package main

import "fmt"

func main() {
	var n int32
	fmt.Scanf("%d", &n)
	s := fmt.Sprintf("%b", n)
	max := 0
	count := 0
	for _, r := range s {
		if fmt.Sprintf("%s", string(r)) == "1" {
			count++
		} else {
			count = 0
		}
		if count > max {
			max = count
		}
	}
	fmt.Println(max)
}
