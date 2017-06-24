package main

import "fmt"

func main() {
	var cases int32
	fmt.Scanf("%d", &cases)
	for i := int32(0); i < cases; i++ {
		var s string
		fmt.Scanf("%s\n", &s)
		printEven(s)
		fmt.Print(" ")
		printOdd(s)
		fmt.Print("\n")
	}
}

func printEven(s string) {
	for i, j := range s {
		if i % 2 == 0 {
			fmt.Print(string(j))
		}
	}
}

func printOdd(s string) {
	for i, j := range s {
		if i % 2 == 1 {
			fmt.Print(string(j))
		}
	}
}
