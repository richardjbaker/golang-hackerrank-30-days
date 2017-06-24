package main

import (
	"fmt"
)

func main() {
	var value int32
	fmt.Scanf("%d", &value)
	for i := int32(0); i < 10; i++ {
		fmt.Printf("%d x %d = %d\n", value, i + 1, value * (i + 1))
	}
}
