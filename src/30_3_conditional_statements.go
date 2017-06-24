package main

import (
	"fmt"
)

func main() {
	var age int32
	fmt.Scanf("%d", &age)
	if (age % 2 == 1) {
		fmt.Println("Weird")
	} else {
		if age >= 2 && age <= 5 {
			fmt.Println("Not Weird")
		} else if age >= 6 && age <= 20 {
			fmt.Println("Weird")
		} else if age > 20 {
			fmt.Println("Not Weird")
		}
	}
}
